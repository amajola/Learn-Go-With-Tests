package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("[0,1,2,4,5,3]", func(t *testing.T) {
		numbers := []int{0, 2, 1, 5, 3, 4}

		got := buildArray(numbers)
		want := []int{0, 1, 2, 4, 5, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("[5,0,1,2,3,4]", func(t *testing.T) {
		numbers := []int{5, 0, 1, 2, 3, 4}

		got := buildArray(numbers)
		want := []int{4, 5, 0, 1, 2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("[0,1,2,4,5,3] Best from LeetCode", func(t *testing.T) {
		numbers := []int{0, 2, 1, 5, 3, 4}

		got := buildArrayBestOnLeetCode(numbers)
		want := []int{0, 1, 2, 4, 5, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("[5,0,1,2,3,4] Best from LeetCode", func(t *testing.T) {
		numbers := []int{5, 0, 1, 2, 3, 4}

		got := buildArrayBestOnLeetCode(numbers)
		want := []int{4, 5, 0, 1, 2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
