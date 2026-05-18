package main

import (
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {
	var result []string
	for i := 0; i < 8; i++ {
		var write strings.Builder

		for _, ch := range input {
			write.WriteString(banner[ch][i])
		}
		result = append(result, write.String())
	}
	return result
}
