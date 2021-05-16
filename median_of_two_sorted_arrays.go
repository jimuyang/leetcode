package main

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	numK, left1, left2 := findKAndLeftInSortedArrays(nums1, nums2, (total+1)/2)

	mid := float64(numK)
	if total%2 == 0 {
		numK1, _, _ := findKAndLeftInSortedArrays(left1, left2, 1)
		mid = (mid + float64(numK1)) / 2
	}
	return mid
}

// 第k大的数
func findKAndLeftInSortedArrays(nums1 []int, nums2 []int, k int) (num int, left1 []int, left2 []int) {
	if len(nums2) == 0 {
		return nums1[k-1], nums1[k:], nil
	}
	if len(nums1) == 0 {
		return nums2[k-1], nil, nums2[k:]
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0], nums1[1:], nums2
		} else {
			return nums2[0], nums1, nums2[1:]
		}
	}

	halfK := k / 2 // 每个数组看第halfK个数
	// 不够halfK个数 用最后一个
	k1 := halfK
	if len(nums1) < halfK {
		k1 = len(nums1)
	}
	t1 := nums1[k1-1]

	k2 := halfK
	if len(nums2) < halfK {
		k2 = len(nums2)
	}
	t2 := nums2[k2-1]

	if t1 < t2 {
		return findKAndLeftInSortedArrays(nums1[k1:], nums2, k-k1)
	} else {
		return findKAndLeftInSortedArrays(nums1, nums2[k2:], k-k2)
	}
}
