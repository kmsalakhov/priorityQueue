package priorityQueue

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type Queue[T any] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	Len() int
	IsEmpty() bool
}

type PriorityQueue[T constraints.Ordered] struct {
	data []T
}

func getParentId(id int) int {
	return (id+1)/2 - 1
}

func getLeftSon(id int) int {
	return (id+1)*2 - 1
}

func getRightSon(id int) int {
	return (id + 1) * 2
}

func (q *PriorityQueue[T]) up(id int) {
	if id == 0 {
		return
	}

	parentId := getParentId(id)
	for id > 0 && q.data[parentId] > q.data[id] {
		q.data[parentId], q.data[id] = q.data[id], q.data[parentId]
		id = parentId
		parentId = getParentId(id)
	}
}

func (q *PriorityQueue[T]) down(id int) {
	for {
		leftSonId := getLeftSon(id)
		rightSonId := getRightSon(id)
		minSonId := -1

		if leftSonId < len(q.data) {
			minSonId = leftSonId

			if rightSonId < len(q.data) && q.data[rightSonId] < q.data[leftSonId] {
				minSonId = rightSonId
			}
		}

		if minSonId != -1 && q.data[minSonId] < q.data[id] {
			q.data[minSonId], q.data[id] = q.data[id], q.data[minSonId]
			id = minSonId
		} else {
			break
		}
	}
}

func (q *PriorityQueue[T]) Push(val T) {
	q.data = append(q.data, val)

	q.up(len(q.data) - 1)
}

func (q *PriorityQueue[T]) Pop() (top T, err error) {
	if len(q.data) == 0 {
		return top, errors.New("empty queue")
	}

	top = q.data[0]

	q.data[0], q.data[len(q.data)-1] = q.data[len(q.data)-1], q.data[0]
	q.data = q.data[:len(q.data)-1]

	q.down(0)

	return
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.data)
}

func (q *PriorityQueue[T]) Peek() (top T, err error) {
	if len(q.data) == 0 {
		return top, errors.New("empty queue")
	}

	return q.data[0], nil
}

func New[T constraints.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{data: make([]T, 0)}
}

func FromSortedSlice[T constraints.Ordered](slice []T) *PriorityQueue[T] {
	data := make([]T, len(slice))
	copy(data, slice)

	return &PriorityQueue[T]{data: data}
}

func Sort[T constraints.Ordered](slice []T) {
	queue := FromSortedSlice(slice)

	for i := range slice {
		slice[i], _ = queue.Pop()
	}
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}
