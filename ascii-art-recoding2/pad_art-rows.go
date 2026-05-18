package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	if width <= 0 {
		return rows
	}
	result := make([]string, len(rows))
	for i, v := range rows{
		padding := width - len(v)
		if padding > 0{
			result[i] = v + strings.Repeat(" ", padding)
		} else {
			result[i] = v
		}
		//space := width - length
		//padding := rows[ch] + strings.Repeat(" ", space)
		//result = append(result, padding)
	}
	return result

}
