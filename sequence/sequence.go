package sequence

/*
Не пустой массив А состоит из N целых чисел.
Проверка всех элементов - это последовательность, которая включает все элементы
от 1 до N и только единожды.

Например массив А:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2

это последовательность, но такой массив:

    A[0] = 4
    A[1] = 1
    A[2] = 3
это не последовательность, т.к. не хватает 2.

Цель - проверить, что массив А это последовательность.

Напишите функцию:
func Solution(A []int) int {
}

Функция возвращает 1 если массив А это последовательность и 0 если нет.

N это целое число в диапазоне [1..100,000];
каждый элемент А находится в диапазоне [1..1,000,000,000].
*/

func Solution(arr []int) int {
	return SolutionC(arr)
}

// Выделение памяти сразу для всех значений
func SolutionB(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	//sort.Ints(arr)
	max := len(arr)
	um := make(map[int]struct{}, max)
	for _, v := range arr {
		if v > max {
			return 0
		}
		if v < 1 {
			return 0
		}
		if _, ok := um[v]; ok {
			return 0
		}
		um[v] = struct{}{}
	}
	return 1
}

// Выделение по мере использования
func SolutionC(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	//sort.Ints(arr)
	max := len(arr)
	um := make(map[int]struct{})
	for _, v := range arr {
		if v > max {
			return 0
		}
		if v < 1 {
			return 0
		}
		if _, ok := um[v]; ok {
			return 0
		}
		um[v] = struct{}{}
	}
	return 1
}
