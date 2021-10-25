package unpaired

import (
	"fmt"
	"testing"
)

func ExampleSolution() {
	s := Solution([]int{9, 3, 9, 3, 9, 7, 9})
	fmt.Println(s)
	// Output: 7
}

/*
func gen(size int) []int {
	const max = 1000000000
	//var want int
	var arr = make([]int, size)
	hs := size / 2
	//p := arr[:hs]
	//	for i, _ := range p {
	//		arr[i] = rand.Intn(max)
	//	}
	t := rand.Perm(hs + 1)
	copy(arr[:hs], t)
	copy(arr[hs:], arr[:hs])
	//want = rand.Intn(max)
	arr[size-1] = t[hs-1]
	rand.Shuffle(size, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}
*/

func TestSolution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]int{9, 3, 9, 3, 9, 7, 9}}, 7},
		{"example", args{[]int{9, 3, 1, 3, 9, 7, 9, 5, 1, 7, 5}}, 9},
		{"repetitive", args{[]int{9, 3, 6, 3, 9, 7, 9, 6, 7}}, 9},
		{"repetitive", args{[]int{1, 1, 1, 2, 1, 1, 1}}, 2},
		{"repetitive", args{[]int{1, 1, 1, 1, 2, 1, 2, 1, 2}}, 2},
		{"one element", args{[]int{100}}, 100},
		// unconditional
		{"even, one pair", args{[]int{102, 102}}, 0},
		{"even heterogeneous", args{[]int{9, 3, 1, 3, 9, 7, 9, 5, 1, 7}}, 0},
		//{"odd heterogeneous", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0},
		{"empty", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.arr); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
