package unpaired

/*
Дан не пустой массив A состоит из N целых чисел.
Массив содержит четное число элементов и каждый элемент может быть парой для другого элемента с таким же значением,
кроме одного элемента, который не содержит пары. Необходимо найти этот элемент.

К примеру:

	A[0] = 9  A[1] = 3  A[2] = 9
	A[3] = 3  A[4] = 9  A[5] = 7
	A[6] = 9

Элементы по индексам 0 и 2 имеют значение 9,
а элемент по индексу 5 имеет значение 7 и не имеет пары.

Необходимо написать функцию, которая вернет элемент без пары

func Solution(A []int) int

N четное число в диапазоне [1..1,000,000];
каждый элемент массива A целое число в диапазоне [1..1,000,000,000];
*/

func Solution(arr []int) int {
	if len(arr)%2 == 0 {
		return 0
	}
	var xs = 0
	for _, v := range arr {
		xs ^= v
	}
	return xs
}
