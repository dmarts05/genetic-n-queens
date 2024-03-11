package util

import "math/rand"

func Sample[T any](s []T, n int) []T {
	length := len(s)
	if n > length {
		n = length
	}

	result := make([]T, n)
	perm := rand.Perm(length)[:n]

	for i, j := range perm {
		result[i] = s[j]
	}

	return result
}
