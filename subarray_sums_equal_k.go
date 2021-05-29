package main

//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums))
	sum := 0
	for i, val := range nums {
		sum += val
		sums[i] = sum
	}

	res := 0
	for i := range sums {
		for j := i; j < len(sums); j++ {

			var s int
			if i > 0 {
				s = sums[j] - sums[i-1]
			} else {
				s = sums[j]
			}

			if k == s {
				res++
			}
		}
	}
	return res
}
