package utils

func removeDuplicates(elements []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for _, v := range elements {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}