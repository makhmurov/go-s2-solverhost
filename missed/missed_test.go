package missed

import (
	"fmt"
	"testing"
)

func ExampleSolution() {
	td := []int{2, 3, 1, 5}
	a := Solution(td)
	fmt.Println(a)
	//Output: 4
}

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]int{2, 3, 1, 5}}, 4},
		{"empty", args{[]int{}}, 1},
		{"last", args{[]int{1}}, 2},
		{"first", args{[]int{2}}, 1},
		{"middle", args{[]int{1, 3}}, 2},
		// Unconditional
		//{"empty", args{[]int{0}}, 3},
		//{"not unique", args{[]int{1, 2, 3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
