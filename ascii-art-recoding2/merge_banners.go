package main

import "fmt"

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	// 1. Allocate a completely new map
	result := make(map[rune][]string)

	// 2. Deep copy all entries from the base map
	for k, v := range base {
		newSlice := make([]string, len(v))
		copy(newSlice, v)
		result[k] = newSlice
	}

	// 3. Deep copy all entries from priority (overwriting duplicates)
	for k, v := range priority {
		newSlice := make([]string, len(v))
		copy(newSlice, v)
		result[k] = newSlice
	}

	return result
}

func main() {
	// Quick verification
	base := map[rune][]string{'A': {"*", "*"}, 'B': {"#"}}
	priority := map[rune][]string{'A': {"@"}, 'C': {"%"}}

	merged := MergeBanners(base, priority)

	// Verify priority won over base for 'A'
	fmt.Println("Merged 'A':", merged['A']) // Outputs: [@]
	fmt.Println("Merged 'B':", merged['B']) // Outputs: [#]
	fmt.Println("Merged 'C':", merged['C']) // Outputs: [%]
}
