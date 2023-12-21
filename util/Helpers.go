package util

import (
	"strconv"
)

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func ConvertToIntSlice(ss []string) (ret []int) {
	filter := func(s string) bool { return s != "" }
	ssNotEmpty := Filter(ss, filter)

	intSlice := make([]int, len(ssNotEmpty))
	for i, line := range ssNotEmpty {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		intSlice[i] = n
	}
	return intSlice
}

func ConvertToStringSlice(ss []int) (ret []string) {
	ret = make([]string, len(ss))
	for i, line := range ss {
		ret[i] = strconv.Itoa(line)
	}
	return ret
}

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)

	result := make([][]string, xl)

	for i := range result {
		result[i] = make([]string, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func CopySlice(input [][]string) [][]string {
	copyWhole := make([][]string, len(input))
	for index, row := range input {
		copyRow := make([]string, len(row))
		copy(copyRow, input[index])
		copyWhole[index] = copyRow
	}
	return copyWhole
}
