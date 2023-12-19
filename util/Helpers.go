package util

import "strconv"

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
