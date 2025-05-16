package stack

type Stack struct {
	items []int
}

// Додати елемент до стека
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// Видалити верхній елемент та повернути його
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value, true
}

// Повернути верхній елемент без видалення
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// Перевірити, чи стек порожній
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
