package verify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"solverhost/solvers"
)

/*
Выполнить проверку выполнения заданий посредством внешнего сервиса.
Имеются две точки входа:
- TASKS_URI - получить тестовые входные данные
- VERIFY_URI - отправить результат на проверку

Необходимо проверить все задачи, выполнив запрос на получение входных данных
и отправиви результаты на сервис проверки.

Следует учесть специфику работы. Связь и работа стороннего сервиса может быть нарушена.
Пользователь может прервать выполнения программы до завершения.

TODO: Handle erorrs:
// errors:
//  - http errors
//  - timeout errors
//  - decode errors
*/

func Do() {
	var tasks = []solvers.Task{
		solvers.Rotation,
		solvers.Unpaired,
		solvers.Sequence,
		solvers.Missed,
	}

	for _, t := range tasks {
		// get payload (payload samples)
		payload, err := getDataSets(t)
		if err != nil {
			panic(err)
			//continue
		}
		// prepare result collection (use map and postprocessing if length unknown)
		results := make([]interface{}, len(payload))
		// get task instance
		handler := handlers[t]
		// process data
		for pn, p := range payload {
			results[pn] = handler(p)
		}

		// send results to verification service
		vreq := Request{
			UserName: cfg.username,
			Task:     solvers.Name(t),
			Results: Results{
				Payload: payload,
				Results: results,
			},
		}

		vr, err := verifyResults(&vreq)
		if err != nil {
			log.Println(err)
		}
		reportResults(t, vr)
	}
}

func getDataSets(taskID solvers.Task) (data []json.RawMessage, err error) {
	u := cfg.dataURL + solvers.Name(taskID)
	fmt.Println("url:", u)

	resp, err := http.Get(u)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()

	defer io.Copy(io.Discard, resp.Body)
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(&data)
	if err != nil {
		return nil, err
	}

	// rp, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// err = json.Unmarshal(rp, &data)
	// if err != nil {
	// 	return nil, err
	// }

	return
}

// Send results for verification
func verifyResults(vreq *Request) (*Response, error) {
	u := cfg.verifyURL
	fmt.Println(u)

	// Prepare Request payload
	// pr, pw := io.Pipe()
	// go func() {
	// 	json.NewEncoder(pw).Encode(vreq)
	// 	pw.Close()
	// }()
	b, err := json.Marshal(vreq)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	pr := bytes.NewReader(b)

	// Do request
	resp, err := http.Post(u, "application/json", pr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Response part
	r := resp.Body
	//r := io.TeeReader(resp.Body, os.Stdout)

	defer io.Copy(io.Discard, r)
	var sr Response
	dec := json.NewDecoder(r)
	err = dec.Decode(&sr)

	if err != nil {
		// handle error
		return nil, err
	}
	//fmt.Println()
	return &sr, nil
}

// Verification report
func reportResults(t solvers.Task, vr *Response) {
	fmt.Printf("Task: %s\n", solvers.Name(t))
	fmt.Printf("Pass: %3d%%\n", vr.Percent)
	if fl := len(vr.Fails); fl > 0 {
		fmt.Printf("Failed: %d\n", fl)
		for _, ft := range vr.Fails {
			fmt.Printf("  Failed DataSet: %d\n", ft.DataSet)
			fmt.Printf("  Expected: %s\n", ft.OriginalResult)
			fmt.Printf("  Received: %s\n", ft.ExternalResult)
			fmt.Println()
		}
	} else {
		fmt.Println("All tests passed")
		fmt.Println()
	}
	fmt.Println()
}
