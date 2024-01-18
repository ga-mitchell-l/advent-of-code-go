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

func TestQuadraticFormula2(t *testing.T) {
	a, b, c := 1.0, -3.0, 2.0
	p, n := QuadraticFormula(a, b, c)

	expectedp := 2.0
	expectedn := 1.0

	if expectedp != p {
		t.Errorf("p was incorrect: got %f, want %f", p, expectedp)
	}

	if expectedn != n {
		t.Errorf("n was incorrect: got %f, want %f", n, expectedn)
	}
}

func TestGetNumberOfWaysToWin(t *testing.T) {
	time := 7
	distance := 9

	result := getNumberOfWaysToWin(time, distance)
	expectedResult := 4

	if result != expectedResult {
		t.Errorf("result was incorrect: got %d, wanted %d", result, expectedResult)
	}
}
