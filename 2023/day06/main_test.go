package main

import "testing"

func TestParseInput(t *testing.T) {
	input := `Time:        40     81     77     72
Distance:   219   1012   1365   1089`

	tt, dd := parseInput(input)
	expectedtt := []int{40, 81, 77, 72}
	expecteddd := []int{219, 1012, 1365, 1089}
	_ = expecteddd
	_ = dd

	if len(tt) != len(expectedtt) {
		t.Errorf("times was incorrect length: got %d, want %d", len(tt), len(expectedtt))
	}

	for i := 0; i < len(tt); i++ {
		if tt[i] != expectedtt[i] {
			t.Errorf("times was incorrect at index %d: got %d, want %d", i, tt[i], expectedtt[i])
		}
	}

	if len(dd) != len(expecteddd) {
		t.Errorf("distances was incorrect length: got %d, want %d", len(dd), len(expecteddd))
	}

	for i := 0; i < len(dd); i++ {
		if dd[i] != expecteddd[i] {
			t.Errorf("distances was incorrect at index %d, got %d, want %d", i, dd[i], expecteddd[i])
		}
	}

}
