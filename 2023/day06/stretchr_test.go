package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput2(t *testing.T) {
	input := `Time:        40     81     77     72
Distance:   219   1012   1365   1089`

	gottt, gotdd := parseInput(input)
	wanttt := []int{40, 81, 77, 72}
	wantdd := []int{219, 1012, 1365, 1089}
	_ = wantdd
	_ = gotdd

	assert.Equal(t, len(gottt), len(wanttt), "times incorrect length")

	for i := 0; i < len(gottt); i++ {
		assert.Equal(t, gottt[i], wanttt[i], "times incorrect at index %d", i)
	}

	assert.Equal(t, len(gotdd), len(wantdd), "distances incorrect length")

	for i := 0; i < len(gotdd); i++ {
		assert.Equal(t, gotdd[i], wantdd[i], "distances incorrect at index %d", i)
	}

}

func TestGetNumberOfWaysToWin2(t *testing.T) {
	time := 7
	distance := 9

	got := getNumberOfWaysToWin(time, distance)
	want := 4

	assert.Equal(t, got, want, "numberOfWays")

	if got != want {
		t.Errorf("result was incorrect: got %d, wanted %d", got, want)
	}
}

func TestQuadraticFormul2(t *testing.T) {
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
			assert.Equal(t, gotPositive, tt.wantPositive, "positive")
			assert.Equal(t, gotNegative, tt.wantNegative, "negative")
		})
	}
}
