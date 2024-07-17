package priorityQueue

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type Queue[T any] interface {
	Push(T)
	Pop() (T, error)
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

func NewPriorityQueue[T constraints.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{data: make([]T, 0)}
}
