package main

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	can := make([]bool, len(nums))
	can[0] = true

	for i := 0; i < len(nums); i++ {
		if !can[i] {
			continue
		}
		for j := i + 1; j <= i+nums[i] && j < len(can); j++ {
			can[j] = true
		}
	}
	return can[len(nums)-1]
}

func canJump1(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	far := 0
	for i := 0; i <= far; i++ {
		if nums[i]+i > far {
			far = nums[i] + i
		}
		if far >= len(nums)-1 {
			return true
		}
	}
	return false
}
