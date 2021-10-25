package task

type Task int

const (
	rotation Task = iota
	unpaired
	sequence
	missed
)

var Names = map[Task]string{
	rotation: "Циклическая ротация",
	unpaired: "Чудные вхождения в массив",
	sequence: "Проверка последовательности",
	missed:   "Поиск отсутствующего элемента",
}

func Name(id Task) string {
	return Names[id]
}
