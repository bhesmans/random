package matrix

import "testing"

func TestDet(t *testing.T) {
	tests := []struct {
		data [][]int
		det  int
	}{
		{
			[][]int{
				{2, 0, 4},
				{3, 1, 0},
				{4, 2, 2}},
			12,
		},
		{
			[][]int{
				{-2, 2, -3},
				{-1, 1, 3},
				{2, 0, -1}},

			18,
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},

			0,
		},
		{
			[][]int{
				{1, 3, 5, 9},
				{1, 3, 1, 7},
				{4, 3, 9, 7},
				{5, 2, 0, 9}},

			-376,
		},
	}
	for _, c := range tests {
		_, m := NewMatrix(c.data)
		ok, det := m.Det()
		if !ok {
			t.Errorf("det(%v) should be able to get a det", m)
		}
		if det != c.det {
			t.Errorf("det(%v) should be %v, got %v", m, c.det, det)
		} else {
			t.Logf("det(%v) is %v", m, det)
		}
	}
}
