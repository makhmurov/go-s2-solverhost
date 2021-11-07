package main

import (
	"encoding/json"
	"solverhost/solvers"
	"solverhost/solvers/missed"
	"solverhost/solvers/rotation"
	"solverhost/solvers/sequence"
	"solverhost/solvers/unpaired"
)

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

type handlerFunc func(json.RawMessage) interface{}

var handlers = map[solvers.Task]handlerFunc{
	solvers.Rotation: rotationHandler,
	solvers.Unpaired: unpairedHandler,
	solvers.Sequence: sequenceHandler,
	solvers.Missed:   missedHandler,
}

func rotationHandler(p json.RawMessage) interface{} {
	var A []int
	var K int
	var tmp = []interface{}{&A, &K}
	err := json.Unmarshal(p, &tmp)
	if err != nil {
		panic(err)
	}
	return rotation.Solution(A, K)
}

func unpairedHandler(p json.RawMessage) interface{} {
	var A []int
	var tmp = []interface{}{&A}
	err := json.Unmarshal(p, &tmp)
	if err != nil {
		panic(err)
	}
	return unpaired.Solution(A)
}

func sequenceHandler(p json.RawMessage) interface{} {
	var A []int
	var tmp = []interface{}{&A}
	err := json.Unmarshal(p, &tmp)
	if err != nil {
		panic(err)
	}
	return sequence.Solution(A)
}

func missedHandler(p json.RawMessage) interface{} {
	var A []int
	var tmp = []interface{}{&A}
	err := json.Unmarshal(p, &tmp)
	if err != nil {
		panic(err)
	}
	return missed.Solution(A)
}
