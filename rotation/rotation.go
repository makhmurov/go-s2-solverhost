package rotation

/*
Массив A состоит из N целых чисел.
Ротация массива - это сдвиг каждого элемента вправо, все элементы с конца двигаются в начало. Например, ротация массива
A = [3, 8, 9, 7, 6] это [6, 3, 8, 9, 7] (все элементы сдвинуты вправо на 1 элемент и 6 сдвигается на первое место).
Цель - это сдвинуть массив A К раз.

Необходимо написать функцию:

func Solution(A []int, K int) []int

К примеру для параметров

    A = [3, 8, 9, 7, 6]
    K = 3

функция должна вернуть [9, 7, 6, 3, 8]. Необходимо сделать 3 ротации:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]

Другой пример:

    A = [1, 2, 3, 4]
    K = 4

результат [1, 2, 3, 4]

Условия:

N и K целые числа в диапазоне [0..100];
каждый элемент массива A целые числа в диапазоне [−1,000..1,000].
*/

func Solution(arr []int, shift int) []int {
	return solutionB(arr, shift)
}

// In loop copy
func solutionA(arr []int, shift int) []int {
	high := len(arr)
	if high == 0 {
		return nil
	}

	res := make([]int, high)
	shift = (high + shift%high) % high
	for i := 0; i < high; i++ {
		dsti := (shift + i) % high
		res[dsti] = arr[i]
	}
	return res
}

// Slice copy
func solutionB(arr []int, shift int) []int {
	len := len(arr)
	if len == 0 {
		return nil
	}
	res := make([]int, len)
	shift = (len + shift%len) % len
	split := len - shift
	copy(res[shift:], arr[:split])
	copy(res[:shift], arr[split:])
	return res
}

// WARNING: Will modify source array
// Append
func solutionC(arr []int, shift int) []int {
	len := len(arr)
	if len == 0 {
		return nil
	}
	shift = len - (len+shift%len)%len
	res := append(arr, arr[:shift]...)
	return res[shift:]
}
