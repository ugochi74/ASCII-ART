package main

func StackTwo(top []string, bottom []string) []string {
	result := make([]string, 0, len(top)+len(bottom))

	result = append(result, top...)
	result = append(result, bottom...)
	return result
}

func StackAll(blocks [][]string) []string {
	if len(blocks) == 0 {
		return []string{}

	}
	var result []string
	for i := 0; i < len(blocks); i++ {
		result = StackTwo(result, blocks[i])
	}
	return result

}
