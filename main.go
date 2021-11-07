package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"solverhost/solvers"
	"solverhost/solvers/missed"
	"solverhost/solvers/rotation"
	"solverhost/solvers/sequence"
	"solverhost/solvers/unpaired"
	"solverhost/verify"
	"syscall"
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

Достаточно ли одной итерации?

> construct task
> data adapter
>

// Structured objects

*/

func getPayload(taskID solvers.Task) (data []json.RawMessage, err error) {
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

func verifyResults(vreq *verify.Request) (*verify.Response, error) {
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
	// TODO: hide communication implementation, but leave transport
	resp, err := http.Post(u, "application/json", pr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Response part
	//r := resp.Body
	r := io.TeeReader(resp.Body, os.Stdout)

	// TODO: LimitReader?
	defer io.Copy(io.Discard, r)
	var sr verify.Response
	dec := json.NewDecoder(r)
	err = dec.Decode(&sr)

	if err != nil {
		// handle error
		return nil, err
	}
	fmt.Println()
	return &sr, nil
}

/*
Интерфейсы функций:
Rotation: "Циклическая ротация",
	func Solution(A []int, K int) []int

Unpaired: "Чудные вхождения в массив",
	func Solution(A []int) int

Sequence: "Проверка последовательности"
	func Solution(A []int) int

Missed:   "Поиск отсутствующего элемента"
	func Solution(A []int) int

Для каждого отдельного набора данных:
Сервис предоставляет аргументы в виде массива агрументов.
Если аргумент один - один элемент массива.
Сервис проверки ожидает значения непосредственно возвращаемое функцией.
*/

type config struct {
	username  string
	dataURL   string
	verifyURL string
}

var cfg config = config{
	username:  os.Getenv("USER_NAME"),
	dataURL:   os.Getenv("TASKS_URI"),
	verifyURL: os.Getenv("VERIFY_URI"),
}

func run() {
	var tasks = []solvers.Task{
		solvers.Rotation,
		//solvers.Unpaired,
		solvers.Sequence,
		//solvers.Missed,
	}

	var handlers = map[solvers.Task]func(json.RawMessage) interface{}{
		solvers.Rotation: func(p json.RawMessage) interface{} {
			var A []int
			var K int
			var tmp = []interface{}{&A, &K}
			err := json.Unmarshal(p, &tmp)
			if err != nil {
				panic(err)
			}
			return rotation.Solution(A, K)
		},
		solvers.Unpaired: func(p json.RawMessage) interface{} {
			var A []int
			var tmp = []interface{}{&A}
			err := json.Unmarshal(p, &tmp)
			if err != nil {
				panic(err)
			}
			return unpaired.Solution(A)
		},
		solvers.Sequence: func(p json.RawMessage) interface{} {
			var A []int
			var tmp = []interface{}{&A}
			err := json.Unmarshal(p, &tmp)
			if err != nil {
				panic(err)
			}
			return sequence.Solution(A)
		},
		solvers.Missed: func(p json.RawMessage) interface{} {
			var A []int
			var tmp = []interface{}{&A}
			err := json.Unmarshal(p, &tmp)
			if err != nil {
				panic(err)
			}
			return missed.Solution(A)
		},
	}

	for _, t := range tasks {
		// get payload (payload samples)
		payload, err := getPayload(t)
		if err != nil {
			panic(err)
			//continue
		}
		// prepare result collection
		// (use map and postprocessing if length unknown)
		// TODO: declare whanted Results type
		results := make([]interface{}, len(payload))
		// get task instance
		handler := handlers[t]
		// process data
		for pn, p := range payload {
			results[pn] = handler(p)
		}

		fmt.Println(results)

		// send results to verification service
		vreq := verify.Request{
			UserName: cfg.username,
			Task:     solvers.Name(t),
			Results: verify.Results{
				Payload: payload,
				Results: results,
			},
		}

		vr, err := verifyResults(&vreq)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Response: %#v\n", vr)
		// errors:
		//  - http errors
		//  - timeout errors
		//  - decode errors
		// print results
	}
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = ctx

	/*
		Отдельный pipline для каждой задачи? Или общий для всёх?
		Недостатки общего: тормозит очередь выполнения.
		Особенно, если задача не предназначена для свободного обработчика.

		Следовательно, ...

		Могут иметься общие этапы для различных задач.


		for avaliable task,

		pipeline:
			1. get task data
			for every element in data
				2. wrap/unwrap task data
				3. process data (call concrete solver) - solver
				*. wrap result data
				*. collect result
			4. create verify request
			5. send verify request
			6. get verify response
			7. collect verify results

			done. print verify results
			// end or wait command
	*/

	run()

	/*
		for {
			select {
			case <-ctx.Done():
				log.Println("all done, shutdown")
				return
			case <-interrupt:
				log.Println("interrupt, shutdown")
				cancel()
				return
			}
		}
	*/
}
