package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	if width <= 0 {
		return rows
	}
	var result []string
	for ch := 0; ch < len(rows); ch++ {
		length := len(rows[ch])
		if length >= width {
			result = append(result, rows[ch])
		}
		space := width - length
		padding := rows[ch] + strings.Repeat(" ", space)
		result = append(result, padding)
	}
	return result

}
