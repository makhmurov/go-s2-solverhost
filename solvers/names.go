package solvers

type Task int

const (
	Rotation Task = iota
	Unpaired
	Sequence
	Missed
)

var Names = map[Task]string{
	Rotation: "Циклическая ротация",
	Unpaired: "Чудные вхождения в массив",
	Sequence: "Проверка последовательности",
	Missed:   "Поиск отсутствующего элемента",
}

func Name(id Task) string {
	return Names[id]
}
