package main

import "strings"

func TrimArtRows(rows []string) []string {
	var result []string

	for i := 0; i < 8; i++ {
		//result[i] = strings.TrimRight(rows[i], " ")
		trim := strings.TrimRight(rows[i], " ")
		result = append(result, trim)
	}
	return result
}
