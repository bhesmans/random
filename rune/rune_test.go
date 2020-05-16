package main

import "testing"

type uf func(string) bool

func testUnique(t *testing.T, f uf) {
	tests := []struct {
		s      string
		unique bool
	}{
		{"a", true},
		{"abc", true},
		{"été", false},
		{"", true},
		{"éè@#!", true},
		{"__", false},
	}
	for _, c := range tests {
		t.Logf("Testing %v ...", c)
		if f(c.s) != c.unique {
			t.Errorf("%s, should be %v", c.s, c.unique)
		}
	}
}

func TestUnique(t *testing.T) {
	testUnique(t, unique)
}

func TestUnique2(t *testing.T) {
	testUnique(t, unique2)
}
