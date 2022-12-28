package war

import (
	"reflect"
	"testing"
)

func TestGame(t *testing.T) {
	testcases := []struct {
		in [][]int
		want int
	}{
		{[][]int{{2, 3}, {1, 4}}, 0},
		{[][]int{{1, 4}, {2, 3}}, 1},
		{[][]int{{1, 4, 2, 3, 4, 5}, {2, 3, 2, 3, 4, 6}}, 0},
	}
	for _, tc := range testcases {
		winner := PlayGame(tc.in)
		if winner != tc.want {
			t.Errorf("Want Player: %d Got Player: %d", tc.want, winner)
		}
	}
}

func TestWar(t *testing.T) {
	testcases := []struct {
		in [][]int
		want [][]int
	}{
		{[][]int{{2, 3, 4, 5}, {2, 3, 4, 6}}, [][]int{{}, {2, 3, 4, 5, 2, 3, 4, 6}}},
	}
	for _, tc := range testcases {
		results := War(tc.in)
		if !reflect.DeepEqual(results, tc.want) {
			t.Errorf("Want: %q Got: %q", tc.want, results)
		}
	}
}