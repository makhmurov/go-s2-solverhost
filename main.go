package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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


*/

func getData(taskID solvers.Task) (data []json.RawMessage, err error) {
	u := cfg.dataURL + solvers.Name(taskID)
	fmt.Println("url:", u)

	resp, err := http.Get(u)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	//
	if err != nil {
		// handle error
		return nil, err
	}
	//_ = body

	//fmt.Println(resp.Status)
	//io.Copy(os.Stdout, resp.Body)
	//fmt.Print(string(body))
	//fmt.Println()

	// Unmarshal body
	// NOTE: в будущем мы будем отправлять тесты нв вычисление по мере чтения
	// обмен должен быть осуществлён посредством канала
	// Что мы будем делать с очередью тестовых данных?
	err = json.Unmarshal(body, &data)

	//jd := json.NewDecoder(resp.Body)
	//err = jd.Decode(data)
	if err != nil {
		return nil, err
	}
	//fmt.Print(data)

	return
}

func postSolution(vreq *verify.Request) (*verify.Response, error) {
	u := cfg.verifyURL

	fmt.Println(u)

	/*
		[ - Tests
			[ - Arguments
				[] - Array argument
			]
		]
	*/

	// vreq := verify.Request{
	// 	UserName: cfg.username,
	// 	Task:     solvers.Name(solvers.Rotation),
	// 	Results: verify.Results{
	// 		Payload: []interface{}{args{[]int{1, 2, 4, 5}, 2}},
	// 		Results: []interface{}{[]int{4, 5, 1, 2}},
	// 	},
	// }

	// TODO: use JSON encoder
	//enc := json.NewEncoder()
	body, err := json.Marshal(vreq)
	if err != nil {
		fmt.Printf("marhsal err: %v\n", err)
		return nil, err
	}
	//fmt.Println(string(body))
	//buf := bytes.NewBuffer(body)
	buf := bytes.NewReader(body)

	_ = buf

	if false {
		return nil, err
	}

	// TODO: hide communication implementation, but leave transport
	resp, err := http.Post(u, "application/json", buf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	//io.Copy(os.Stdout, resp.Body)
	//fmt.Println()
	r := io.TeeReader(resp.Body, os.Stdout)

	var sr verify.Response
	dec := json.NewDecoder(r)
	err = dec.Decode(&sr)
	//r, err = io.ReadAll(resp.Body)
	//
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
*/

// handler function interface
type handlerFunc func(json.RawMessage) interface{}

type config struct {
	username  string
	dataURL   string
	verifyURL string
}

func (c *config) Init() {
	c.username = os.Getenv("USER_NAME")
	c.dataURL = os.Getenv("TASKS_URI")
	c.verifyURL = os.Getenv("VERIFY_URI")
}

var cfg config

func main() {
	// TODO: check that env are set
	cfg.Init()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = ctx

	turl := os.Getenv("TASKS_URI")
	vurl := os.Getenv("VERIFY_URI")
	_ = vurl

	u, err := url.Parse(turl)
	if err != nil {
		//http.Get(u.String())
		_ = u
	}

	//getData()
	//postSolution()

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

	var tasks = []solvers.Task{
		solvers.Rotation,
		solvers.Unpaired,
		solvers.Sequence,
		solvers.Missed,
	}

	var sol = map[solvers.Task]handlerFunc{
		solvers.Rotation: func(data json.RawMessage) interface{} {
			// TODO: unmarshal args
			var a struct {
				A []int
				K int
			}
			var tmp = []interface{}{&a.A, &a.K}
			err := json.Unmarshal(data, &tmp)
			if err != nil {
				panic(err)
			}
			//A := args[0].([]int)
			//K := args[1].(int)
			return rotation.Solution(a.A, a.K)
		},
		solvers.Unpaired: func(data json.RawMessage) interface{} {
			// TODO: unmarshal args
			var a struct {
				A []int
			}
			var tmp = []interface{}{&a.A}
			err := json.Unmarshal(data, &tmp)
			if err != nil {
				panic(err)
			}
			//A := args[0].([]int)
			//R := make([]int, 1)
			//R[0] = unpaired.Solution(a.A)
			return unpaired.Solution(a.A)
		},
		solvers.Sequence: func(data json.RawMessage) interface{} {
			// TODO: unmarshal args
			var a struct {
				A []int
			}
			var tmp = []interface{}{&a.A}
			err := json.Unmarshal(data, &tmp)
			if err != nil {
				panic(err)
			}
			//A := args[0].([]int)
			//R := make([]int, 1)
			//R[0] = sequence.Solution(a.A)
			return sequence.Solution(a.A)
		},
		solvers.Missed: func(data json.RawMessage) interface{} {
			// TODO: unmarshal args
			var a struct {
				A []int
			}
			var tmp = []interface{}{&a.A}
			err := json.Unmarshal(data, &tmp)
			if err != nil {
				panic(err)
			}
			//A := args[0].([]int)
			//R := make([]int, 1)
			//R[0] = missed.Solution(a.A)
			return missed.Solution(a.A)
		},
	}

	for _, t := range tasks {
		// get data (data samples)
		data, err := getData(t) // aka service.Get
		if err != nil {
			panic(err)
			//continue
		}
		// prepare result collection
		// (use map and postprocessing if length unknown)
		trd := make([]interface{}, len(data)) // TODO: declare whanted Results type
		// get task instance
		handle := sol[solvers.Task(t)]
		// process data
		for dsid, d := range data {
			// d - list of arguments (array of arguments)
			// NOTE: Solution - defined handlerFunc!!!
			//r := ti.Solution(d) // unified processing interface
			r := handle(d) // unified processing interface
			// NOTE: inside
			// 	unwrap (preprocess) using custom wrapper
			//  wrap (postprocess) result using custom wrapper
			// 	get pp function? but interface?
			//  unwrap may be defined inside concrete task wrapper
			//  are there common task wrapper possible?

			// r - wrapped results? (WARNING: единственной ответственности)
			// save r at data sample index
			trd[dsid] = r
		}

		fmt.Println(trd)

		// send result to verification services
		// NewRequest, use trd as Results, use data as Payload
		// username?, taskname, payload, results
		vreq := verify.Request{
			UserName: cfg.username,
			Task:     solvers.Name(t),
			Results: verify.Results{
				Payload: &data,
				Results: trd,
			},
		}
		_ = vreq
		fmt.Println(vreq)
		fmt.Println(vreq.Results.Results)
		vr, err := postSolution(&vreq)
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
