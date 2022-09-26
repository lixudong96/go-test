package first_test

import (
	"reflect"
	"testing"
)

func TestStrGetCount(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		got := GetStrShowCount("abcabcccd")
		want := map[string]int{"a": 2, "b": 2, "c": 4, "d": 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("error")
		}
	})
}
