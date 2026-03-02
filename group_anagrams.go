package main

import (
	"fmt"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	// 统计每个str中的字母出现次数

	stats := make(map[string][]string)
	for _, str := range strs {
		stat := make([]int, 26)
		for _, char := range []rune(str) {
			stat[char-'a']++
		}
		// 将stat转换为字符串 作为每个str的唯一标识

		sb := strings.Builder{}
		for _, count := range stat {
			sb.WriteString(fmt.Sprintf("%d", count))
			sb.WriteString("-")
		}
		statStr := sb.String()
		stats[statStr] = append(stats[statStr], str)
	}

	res := make([][]string, 0)
	for _, strs := range stats {
		res = append(res, strs)
	}
	return res
}
