package sequence

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func ExampleSolution() {
	fmt.Println(Solution([]int{5, 1, 4, 2, 3}))
	// Output: 1
}

func ExampleSolution_inconsistent() {
	fmt.Println(Solution([]int{4, 1, 3}))
	// Output: 0
}

type args struct {
	arr []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"empty", args{[]int{}}, 0},
	{"example", args{[]int{5, 1, 4, 2, 3}}, 1},
	{"consistent", args{[]int{3, 5, 2, 4, 1}}, 1},
	{"inconsistent", args{[]int{4, 1, 3}}, 0},
	{"inconsistent", args{[]int{4, 2, 3}}, 0},
	{"big100", args{genTest(100)}, 1},
	{"big1000", args{genTest(1000)}, 1},
	{"big10000", args{genTest(10000)}, 1},
	{"big1000000", args{genTest(1000000)}, 1},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.arr); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

type impl func([]int) int

func (a *args) UnmarshalJSON(buf []byte) error {
	var tmp = []interface{}{&a.arr}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
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

func genTest(len int) []int {
	a := make([]int, len+1)
	for i := range a {
		a[i] = i
	}
	a = a[1:]
	rand.Shuffle(len, func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

func benchFunction(f impl, b *testing.B) {
	b.StopTimer()

	btd, err := loadTests("testdata/data.json")
	if err != nil {
		panic(err)
	}
	bigct := genTest(1000000)
	btd = append(btd, tests[1].args, tests[2].args, tests[3].args, args{bigct})

	b.ResetTimer()
	b.StartTimer()
	for n, td := range btd {
		tname := fmt.Sprintf("test %d", n)
		b.Run(tname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(td.arr)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	benchFunction(Solution, b)
}

func BenchmarkSolutionB(b *testing.B) {
	benchFunction(SolutionB, b)
}

func BenchmarkSolutionC(b *testing.B) {
	benchFunction(SolutionC, b)
}
