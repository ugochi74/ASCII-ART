package main

import (
	"strings"
)

func GenerateArt(text string, banner map[rune][]string) string {
	if text == "" {
		return ""
	}
	input := strings.ReplaceAll(text, "\n", "\\n")
	if input == "\\n" {
		return "\n"
	}

	if strings.ReplaceAll(text, "\n", "") == "" {
		return strings.Repeat("\n", len(input)/2*8)
	}

	part := SplitInput(text)
	var write strings.Builder
	for i, parts := range part {
		if parts == "" {
			if i == len(part)-1 {
				write.WriteString(strings.Repeat("\n", 8))
			} else {
				write.WriteString("\n")
			}
		} else {
			row := RenderLine(parts, banner)
			for _, rows := range row {
				write.WriteString(rows + "\n")
			}
		}
	}
	return write.String()
}
