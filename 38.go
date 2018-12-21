package algo

import "strings"

func permutation(str string) []string {
	if len(str) == 0 {
		return nil
	}
	bytesStr := []byte(str)
	var result [][]byte
	var current []byte
	var isVisited []bool
	for i := 0; i < len(bytesStr); i++ {
		isVisited = append(isVisited, false)
	}
	permutationCore(&result, &current, isVisited, bytesStr)
	var sb strings.Builder
	var results []string
	for _, v := range result {
		for _, iv := range v {
			sb.WriteString(string(iv))
		}
		results = append(results, sb.String())
		sb.Reset()
	}
	return results
}

func permutationCore(result *[][]byte, current *[]byte, isVisited []bool, bytesStr []byte) {
	if len(*current) == len(bytesStr) {
		*result = append(*result, *current)
		return
	}
	for i := 0; i < len(bytesStr); i++ {
		if !isVisited[i] {
			isVisited[i] = true
			*current = append(*current, bytesStr[i])
			permutationCore(result, current, isVisited, bytesStr)
			isVisited[i] = false
			*current = (*current)[: len(*current)-1 : len(*current)-1]
		}
	}
}
