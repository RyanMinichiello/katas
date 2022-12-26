package wave

import (
	"reflect"
	"testing"
)

func TestWave(t *testing.T) {
	testcases := []struct {
		in string
		want []string
	}{
		{"hello", []string{"Hello", "hEllo", "heLlo", "helLo", "hellO"}},
		{"hello me", []string{"Hello me", "hEllo me", "heLlo me", "helLo me", "hellO me", "hello Me", "hello mE"}}, 
	}

	for _, tc := range testcases {
		w := Wave(tc.in)
		if !reflect.DeepEqual(w, tc.want) {
			t.Errorf("Wave: %q, want %q", w, tc.want)
		}
	}
}