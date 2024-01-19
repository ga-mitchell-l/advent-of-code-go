package main

import (
	_ "embed"
	"testing"
)

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

func TestGetNumberOfWaysToWin(t *testing.T) {
	time := 7
	distance := 9

	result := getNumberOfWaysToWin(time, distance)
	expectedResult := 4

	if result != expectedResult {
		t.Errorf("result was incorrect: got %d, wanted %d", result, expectedResult)
	}
}

func TestQuadraticFormula(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name         string
		args         args
		wantPositive float64
		wantNegative float64
	}{
		{name: "foo", args: args{a: 1.0, b: -3.0, c: 2.0}, wantPositive: 2, wantNegative: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPositive, gotNegative := QuadraticFormula(tt.args.a, tt.args.b, tt.args.c)
			if gotPositive != tt.wantPositive {
				t.Errorf("QuadraticFormula() gotPositive = %v, want %v", gotPositive, tt.wantPositive)
			}
			if gotNegative != tt.wantNegative {
				t.Errorf("QuadraticFormula() gotNegative = %v, want %v", gotNegative, tt.wantNegative)
			}
		})
	}
}
