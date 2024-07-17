package priorityQueue

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pushPop[T constraints.Ordered](t *testing.T, val T) {
	queue := NewPriorityQueue[T]()

	t.Run("TestPushPop", func(t *testing.T) {
		queue.Push(val)
		actual, err := queue.Pop()
		if assert.NoError(t, err) {
			assert.Equal(t, val, actual)
		}
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

func testSort[T constraints.Ordered](t *testing.T, testId int, elems []T) {
	queue := NewPriorityQueue[T]()

	t.Run("#"+strconv.Itoa(testId), func(t *testing.T) {
		expected := make([]T, len(elems))
		copy(expected, elems)

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

		for _, elem := range elems {
			queue.Push(elem)
		}

		actual := make([]T, len(elems))

		for i := range elems {
			var err error
			actual[i], err = queue.Pop()

			if assert.NoError(t, err) {
				assert.Equal(t, expected[i], actual[i])
			}
		}
	})
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
		BitTestSize = 1_000_000
	)

	arr := make([]int, BitTestSize)

	for i := range arr {
		arr[i] = rand.Int()
	}

	testSort(t, 1, arr)
}
