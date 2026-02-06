package main

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums) // 从小到大

	// 从小遍历一个数n 从更大的数组中 找2个数求和等于-n
	// 为啥是更大的数组？因为更小的数都已经处理过了

	for i, n := range nums {
		// n重复 答案就重复了
		if i > 0 && n == nums[i-1] {
			continue
		}

		ret := find2Sum(nums[i+1:], -n)
		if len(ret) > 0 {
			for _, tuple := range ret {
				result = append(result, []int{n, tuple[0], tuple[1]})
			}
		}

	}
	return result
}

func find2Sum(nums []int, n int) [][]int {
	res := make([][]int, 0)

	// 2侧夹逼
	lo, hi := 0, len(nums)-1
	for lo < hi {
		// 移动后相同 也会重复
		if lo > 0 && nums[lo] == nums[lo-1] {
			lo++
			continue
		}
		if hi < len(nums)-1 && nums[hi] == nums[hi+1] {
			hi--
			continue
		}

		sum := nums[lo] + nums[hi]

		switch {
		case sum == n:
			res = append(res, []int{nums[lo], nums[hi]})
			lo++
			hi--
		case sum < n:
			lo++
		case sum > n:
			hi--
		}
	}
	return res
}
