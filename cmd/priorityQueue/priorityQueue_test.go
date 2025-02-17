package priorityQueue

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pushPop[T constraints.Ordered](t *testing.T, val T) {
	queue := New[T]()

	t.Run("TestPushPop", func(t *testing.T) {
		queue.Push(val)

		assert.Equal(t, false, queue.IsEmpty())
		assert.Equal(t, 1, queue.Len())

		peek, err := queue.Peek()
		if assert.NoError(t, err) {
			assert.Equal(t, val, peek)
		}

		actual, err := queue.Pop()
		if assert.NoError(t, err) {
			assert.Equal(t, val, actual)
		}

		assert.Equal(t, 0, queue.Len())
		assert.Equal(t, true, queue.IsEmpty())
	})
}

func TestPushPopInt(t *testing.T) {
	pushPop[int](t, 5)
}

func TestPushPopFloat(t *testing.T) {
	pushPop[float32](t, 0.2)
}

func TestPushPopUInt(t *testing.T) {
	pushPop[uint](t, uint(5))
}

func TestPushPopString(t *testing.T) {
	pushPop[string](t, "Hello world!")
}

func testSort[T int | string | float64](t *testing.T, testId int, elems []T) {
	queue := New[T]()

	t.Run("#"+strconv.Itoa(testId), func(t *testing.T) {
		expected := make([]T, len(elems))
		copy(expected, elems)

		sortSlice(expected)

		for _, elem := range elems {
			queue.Push(elem)
		}

		assert.Equal(t, len(expected), queue.Len())

		actual := make([]T, len(elems))

		for i := range elems {
			var err error
			actual[i], err = queue.Pop()

			if assert.NoError(t, err) {
				assert.Equal(t, expected[i], actual[i])
			}
		}

		assert.Equal(t, 0, queue.Len())
	})
}

func sortSlice[T int | string | float64](expected []T) {
	switch v := any(expected[0]).(type) {
	case int:
		expected := any(expected).([]int)
		sort.Ints(expected)
	case string:
		expected := any(expected).([]string)
		sort.Strings(expected)
	case float64:
		expected := any(expected).([]float64)
		sort.Float64s(expected)
	default:
		panic(fmt.Sprintf("unknown type %T", v))
	}
}

func TestSortSmallInts(t *testing.T) {
	testSort(t, 1, []int{3, 1, 2})
	testSort(t, 2, []int{1, 2, 3})
	testSort(t, 3, []int{2, 1, 3})
	testSort(t, 4, []int{2, 1, 3, -5})
	testSort(t, 5, []int{7, 15, 8, -32, 1})
}

func TestSortSmallFloats(t *testing.T) {
	testSort(t, 1, []float64{3.5, 1.2, 2.7})
	testSort(t, 2, []float64{0.15, -1.85, 2.995, -0.1172})
}

func TestSortSmallStrings(t *testing.T) {
	testSort(t, 1, []string{"B", "C", "A", "K"})
	testSort(t, 2, []string{"Hello world!", "What?", "Be!", "Lol"})
}

func TestSortBigInts(t *testing.T) {
	const (
		BigTestSize = 1_000_000
	)

	arr := make([]int, BigTestSize)

	for i := range arr {
		arr[i] = rand.Int()
	}

	testSort(t, 1, arr)
}

func TestEmptyQueuePop(t *testing.T) {
	queue := New[int]()

	t.Run("1", func(t *testing.T) {
		_, err := queue.Pop()

		assert.Equal(t, errors.New("empty queue"), err)

		assert.Equal(t, 0, queue.Len())
	})
}

func TestPushPopPushPop(t *testing.T) {
	queue := New[int]()

	t.Run("1", func(t *testing.T) {
		queue.Push(1)
		queue.Push(2)
		assert.Equal(t, 2, queue.Len())

		actual, err := queue.Pop()
		if assert.NoError(t, err) {
			assert.Equal(t, 1, actual)
		}

		assert.Equal(t, 1, queue.Len())

		queue.Push(3)
		assert.Equal(t, 2, queue.Len())

		actual, err = queue.Pop()
		if assert.NoError(t, err) {
			assert.Equal(t, 2, actual)
		}

		assert.Equal(t, 1, queue.Len())

		actual, err = queue.Pop()
		if assert.NoError(t, err) {
			assert.Equal(t, 3, actual)
		}

		assert.Equal(t, 0, queue.Len())
	})
}

func testFromSortedSlice[T constraints.Ordered](t *testing.T, testId int, slice []T) {
	t.Run("#"+strconv.Itoa(testId), func(t *testing.T) {
		queue := FromSortedSlice(slice)
		assert.Equal(t, len(slice), queue.Len())

		for _, expected := range slice {
			actual, err := queue.Pop()

			if assert.NoError(t, err) {
				assert.Equal(t, expected, actual)
			}
		}

		assert.Equal(t, 0, queue.Len())
		assert.Equal(t, true, queue.IsEmpty())
	})
}

func testSortSlice[T int | string | float64](t *testing.T, testId int, slice []T) {
	t.Run("#"+strconv.Itoa(testId), func(t *testing.T) {
		expected := make([]T, len(slice))
		copy(expected, slice)

		actual := make([]T, len(slice))
		copy(actual, slice)

		sortSlice(expected)
		Sort(actual)

		assert.Equal(t, expected, actual)
	})
}

func TestSort(t *testing.T) {
	testSortSlice(t, 1, []int{5, 3, 2, 1})
	testSortSlice(t, 2, []int{5, 3, 2, 1, 7, -13})
	testSortSlice(t, 3, []int{3, 3, 3, 3})
	testSortSlice(t, 4, []int{-11, -12, 5, 77, 33})
	testSortSlice(t, 5, []string{"B", "A", "F", "C"})
	testSortSlice(t, 6, []string{"BRO", "liKE", "Type", "HELLO wOrLd!!", "*"})
}

func TestFromSortedSlice(t *testing.T) {
	testFromSortedSlice(t, 1, []int{1, 2, 3, 4, 5})
	testFromSortedSlice(t, 2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	testFromSortedSlice(t, 3, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	testFromSortedSlice(t, 4, []string{"a", "b", "c", "d"})
}
