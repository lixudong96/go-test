package first_test

import "testing"

func TestIsALeapYear(t *testing.T) {
	t.Run("普通闰年 2004", func(t *testing.T) {
		got := IsleapYear(2004)
		want := true
		if got != want {
			t.Errorf("got '%v' want '%v'", got, true)
		}
	})

	t.Run("不是闰年 2002", func(t *testing.T) {
		got := IsleapYear(2002)
		want := false
		if got != want {
			t.Errorf("got '%v' want '%v'", got, true)
		}
	})

	t.Run("特殊闰年 2000", func(t *testing.T) {
		got := IsleapYear(2000)
		want := true
		if got != want {
			t.Errorf("got '%v' want '%v'", got, true)
		}
	})

	t.Run("特殊闰年 1900", func(t *testing.T) {
		got := IsleapYear(1900)
		want := false
		if got != want {
			t.Errorf("got '%v' want '%v'", got, true)
		}
	})
}
