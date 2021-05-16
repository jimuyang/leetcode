package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	runes := []rune(expression)
	res := make([]int, 0)
	for i, ru := range runes {
		switch ru {
		case '+':
			left := diffWaysToCompute(string(runes[:i]))
			right := diffWaysToCompute(string(runes[i+1:]))
			for _, l := range left {
				for _, r := range right {
					res = append(res, l+r)
				}
			}
		case '-':
			left := diffWaysToCompute(string(runes[:i]))
			right := diffWaysToCompute(string(runes[i+1:]))
			for _, l := range left {
				for _, r := range right {
					res = append(res, l-r)
				}
			}
		case '*':
			left := diffWaysToCompute(string(runes[:i]))
			right := diffWaysToCompute(string(runes[i+1:]))
			for _, l := range left {
				for _, r := range right {
					res = append(res, l*r)
				}
			}
		default:
			continue
		}
	}
	if len(res) == 0 {
		num, err := strconv.Atoi(expression)
		if err != nil {
			fmt.Println(err)
		} else {
			res = append(res, num)
		}
	}
	return res

}
