package rotation

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type impl func([]int, int) []int

type args struct {
	A []int
	K int
}

func (a *args) UnmarshalJSON(buf []byte) error {
	var tmp = []interface{}{&a.A, &a.K}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{name: "example 1", args: args{[]int{3, 8, 9, 7, 6}, 3}, want: []int{9, 7, 6, 3, 8}},
	{name: "example 2", args: args{[]int{1, 2, 3, 4}, 4}, want: []int{1, 2, 3, 4}},

	{name: "positive shift", args: args{[]int{1, 2, 3, 4, 5}, 3}, want: []int{3, 4, 5, 1, 2}},
	{name: "overflow shift", args: args{[]int{1, 2, 3, 4, 5}, 6}, want: []int{5, 1, 2, 3, 4}},
	{name: "negative shift", args: args{[]int{1, 2, 3, 4, 5}, -1}, want: []int{2, 3, 4, 5, 1}},
	{name: "zero shift", args: args{[]int{1, 2, 3, 4, 5}, 0}, want: []int{1, 2, 3, 4, 5}},
	{name: "bound shift", args: args{[]int{1, 2, 3, 4, 5}, 5}, want: []int{1, 2, 3, 4, 5}},
	{name: "even shift", args: args{[]int{1, 2, 3, 4}, 2}, want: []int{3, 4, 1, 2}},

	{name: "one el", args: args{[]int{1}, 2}, want: []int{1}},
	{name: "empty", args: args{[]int{}, -1}, want: nil},
}

func ExampleSolution() {
	a, k := []int{3, 8, 9, 7, 6}, 3
	r := Solution(a, k)
	fmt.Printf("%#v\n", r)
	// Output: []int{9, 7, 6, 3, 8}
}

func ExampleSolution_e2() {
	a, k := []int{1, 2, 3, 4}, 4
	r := Solution(a, k)
	fmt.Printf("%#v\n", r)
	// Output: []int{1, 2, 3, 4}
}

func testFunction(fname string, f impl) func(t *testing.T) {
	return func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", fname, tt.args.A, tt.args.K, got, tt.want)
				}
			})
		}
	}
}

func TestSolution(t *testing.T) {
	testFunction("Solution", Solution)(t)
}

func loadTests(path string) (tests []args, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&tests)
	if err != nil {
		return nil, err
	}
	return tests, nil
}

func benchFunction(f impl) func(b *testing.B) {
	var tt []args = make([]args, 0, 20)
	for _, t := range tests {
		tt = append(tt, t.args)
	}
	bt, err := loadTests("testdata/data.json")
	if err != nil {
		panic(err)
	}
	tt = append(tt, bt...)
	return func(b *testing.B) {
		b.ResetTimer()
		for i, t := range tt {
			tname := fmt.Sprintf("test %d", i) //strconv.Itoa(i)
			b.Run(tname, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(t.A, t.K)
				}
			})
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	benchFunction(Solution)(b)
}
