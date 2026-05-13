package main

func StackTwo(top []string, bottom []string) []string {
	var result []string

	result = append(result, top...)
	result = append(result, bottom...)
	return result
}

func StackAll(blocks [][]string) []string {
	var result []string
	for i := 0; i < len(blocks); i++ {
		result = StackTwo(result, blocks[i])
	}
	return result
}
