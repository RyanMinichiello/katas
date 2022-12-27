package thief

import (
	"testing"
)

func TestSteal(t *testing.T) {
	testcases := []struct {
		in []string
		want int
	}{
		{[]string{"*.*.*.*.*.", "...*..**..", "**.**...*.", "**..**..**"}, 13},
		{[]string{"*.*.*.*.*."}, 5},
		{[]string{"...*..**.."}, 2},
		{[]string{"**.**...*."}, 3},
		{[]string{"**..**..**"}, 3},
		{[]string{"**********"}, 5},
		{[]string{".........."}, 0},
	}
	for _, tc := range testcases {
		reps := Steal(tc.in)
		if reps != tc.want {
			t.Errorf("Want: %d Got: %d", tc.want, reps)
		}
	}
}