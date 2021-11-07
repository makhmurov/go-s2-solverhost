package verify

type Task int

const (
	Rotation Task = iota
	Unpaired
	Sequence
	Missed
)

var names = map[Task]string{
	Rotation: "Циклическая ротация",
	Unpaired: "Чудные вхождения в массив",
	Sequence: "Проверка последовательности",
	Missed:   "Поиск отсутствующего элемента",
}

func (t Task) Name() string {
	return names[t]
}

func Name(id Task) string {
	return names[id]
}
