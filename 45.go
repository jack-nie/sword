package algo

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type stringArray []string

func (s stringArray) Len() int {
	return len(s)
}

func (s stringArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stringArray) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}

func minNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(stringArray(strs))
	var sb strings.Builder
	for _, v := range strs {
		sb.WriteString(v)
	}
	result, err := strconv.Atoi(sb.String())
	if err != nil {
		log.Fatal(err)
	}
	return result
}
