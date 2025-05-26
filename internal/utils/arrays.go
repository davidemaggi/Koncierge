package utils

func DistinctStrings(input []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, val := range input {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}

	return result
}
