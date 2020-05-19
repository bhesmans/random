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

func TestPermutation(t *testing.T) {
	tests := []struct {
		s1, s2 string
		perm   bool
	}{
		{"", "", true},
		{"eazr", "fdsqsqfdq", false},
		{"azerty", "zrtyea", true},
		{"azerty", "Azerty", false},
		{"azert", "azerty", false},
		{"éé@èè", "èé@éè", true},
	}
	for _, c := range tests {
		t.Logf("Testing %v ...", c)
		if permutation(c.s1, c.s2) != c.perm {
			t.Errorf("%s ... %s, should be %v", c.s1, c.s2, c.perm)
		}
	}
}

func TestEscape(t *testing.T) {
	tests := []struct {
		s1, s2 string
	}{
		{"", ""},
		{"abc", "abc"},
		{"aa bb", "aa%20bb"},
		{" a ", "%20a%20"},
	}
	for _, c := range tests {
		t.Logf("Testing %v ...", c)
		if escapeSpaces(c.s1) != c.s2 {
			t.Errorf("%s != %s", c.s1, c.s2)
		}
	}
}
