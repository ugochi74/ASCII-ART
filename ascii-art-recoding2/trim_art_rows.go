package main

import "strings"

func TrimArtRows(rows []string) []string {
	result := make([]string, 0, len(rows))

	for _, i := range rows {
		//result[i] = strings.TrimRight(rows[i], " ")
		trim := strings.TrimRight(i, " ")
		result = append(result, trim)
	}
	return result
}
