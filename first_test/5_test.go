package first_test

import (
	"reflect"
	"testing"
)

func Test5(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		got := SliceOpterator()
		want := []int{1, 2, 4, 5, 7, 8, 10, 11, 13, 14, 16, 17, 19, 20, 22, 23, 25, 26, 28, 29, 31, 32, 34, 35, 37, 38, 40, 41, 43, 44, 46, 47, 49, 50, 666}
		if !reflect.DeepEqual(got, want) {
			t.Error("error")
		}
	})
}
