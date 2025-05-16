package queue

type Queue struct {
	items []int
}

// Додати елемент в кінець черги
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// Видалити перший елемент та повернути його
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value, true
}

// Повернути перший елемент без видалення
func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

// Перевірити, чи черга порожня
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
