package first_test

import (
	"reflect"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	t.Run("100 内", func(t *testing.T) {
		got := PrimeNumber(100)
		want := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("匹配结果不正确")
		}
	})
}
