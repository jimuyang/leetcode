package main

import "container/list"

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//说明：你不能倾斜容器。

// 输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

type Height struct {
	I int
	H int
}

func maxArea1(height []int) int {
	// 双指针法 每次移动较矮的指针
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		if height[i] < height[j] {
			area := height[i] * (j - i)
			if area > max {
				max = area
			}
			i++
		} else {
			area := height[j] * (j - i)
			if area > max {
				max = area
			}
			j--
		}
	}
	return max
}

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	var res int

	// 单调增 栈
	stack := list.New()
	stack.PushFront(Height{0, height[0]})

	for i := 1; i < len(height); i++ {
		h := height[i]

		// 当前值作为右边（不如右边更高的
		if !(i < len(height)-1 && h <= height[i+1]) {
			// 计算与栈内左边的面积
			for e := stack.Front(); e != nil; e = e.Next() {
				temp := e.Value.(Height)
				var area int
				if temp.H > h {
					area = h * (i - temp.I)
				} else {
					area = temp.H * (i - temp.I)
				}
				if area > res {
					res = area
				}
			}
		}

		// 当前值作为左边（不如栈里更高的-栈顶
		if h > stack.Front().Value.(Height).H {
			stack.PushFront(Height{i, h})
		}
	}
	return res
}
