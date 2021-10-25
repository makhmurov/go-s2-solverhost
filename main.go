package main

import (
	"fmt"
	"math"
	"math/rand"

	"solverhost/sequence"
)

const EulerGamma = 0.57721566490153286060651209008240243104215933593992359880576723 // https://oeis.org/A001620

func test(arr []int) float64 {
	var acc float64 = 0
	for _, v := range arr {
		acc += 1.0 / float64(v)
	}
	fmt.Println(acc)
	return acc
}

func test2(arr []int) float64 {
	m := map[int]float64{}
	for _, r := range arr {
		m[r]++
	}
	var hm float64
	for _, c := range m {
		hm += c * math.Log2(c)
	}
	var l = float64(len(arr))
	var e = math.Log2(l) - hm/l
	fmt.Println(e)
	return e
}

func main() {
	var arr []int
	//arr = []int{7, 4, 9, 1, 3, 5, 2, 6}
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	max := len(arr)
	fmt.Println(max)
	test(arr)
	fmt.Println("Entropy")
	test2(arr)

	fmt.Println("\nTest2")
	max = 20
	arr = make([]int, max+1)
	for i := range arr {
		arr[i] = i
	}
	arr = arr[1:]
	test(arr)
	fmt.Println("Entropy")
	test2(arr)

	fmt.Println("\nShuffle")
	rand.Shuffle(max, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	test(arr)
	fmt.Println("Entropy")
	test2(arr)

	fmt.Println("\nChange one element")
	arr[max/2] = max/2 - 2
	test(arr)
	fmt.Println("Entropy")
	test2(arr)

	fmt.Println("Euler asymtomatic")
	fmt.Println(math.Log(float64(max)) + EulerGamma)
	fmt.Println(math.Lgamma(float64(max)))

	arr = []int{1, 2, 4, 4, 5, 6, 6, 8}
	test(arr)
	//fmt.Println((arr[0] + arr[max-1]) * max / 2.)

	var res int
	fmt.Println("\nSolutionE")
	arr = []int{1, 2, 3, 4}
	res = sequence.SolutionE(arr)
	fmt.Println(res)
	arr = []int{1, 2, 2, 12}
	res = sequence.SolutionE(arr)
	fmt.Println(res)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	res = sequence.SolutionF(arr)
	fmt.Println(res)
	arr = []int{1, 2, 2, 12, 5, 6, 7, 8, 9, 10, 11, 12}
	res = sequence.SolutionF(arr)
	fmt.Println(res)
}
