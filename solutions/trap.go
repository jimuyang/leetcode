package solutions

import (
	"container/list"
	"fmt"
)

type Height struct {
	I int
	H int
}

// Trap 供 main 或外部调用的导出包装
func Trap(height []int) int {
	return trap(height)
}

func trap(height []int) int {
	res := 0
	// 单调减的栈
	stack := list.New()

	for i, h := range height {
		if h == 0 && stack.Len() == 0 {
			continue
		}
		// 不特殊处理高度为0

		height := &Height{i, h}

		top := stack.Front()

		// 更矮的放入
		if top == nil || h < top.Value.(*Height).H {
			stack.PushFront(height)
		} else {
			// 遇到更高或者相等的 弹出并计算雨水 直到更矮
			for top != nil && h >= top.Value.(*Height).H {
				stack.Remove(top)
				bot := top.Value.(*Height).H
				top = stack.Front()
				if top != nil {
					// 计算雨水
					min := h
					if min > top.Value.(*Height).H {
						min = top.Value.(*Height).H
					}
					area := (min - bot) * (i - top.Value.(*Height).I - 1)
					res += area
					fmt.Println(area, i, h)
				}

			}
			stack.PushFront(height)
		}

	}
	return res

}

// func trap(height []int) int {
// 	res := 0
// 	// 单调减的栈
// 	stack := list.New()

// 	for i, h := range height {
// 		if h == 0 {
// 			continue
// 		}

// 		height := &Height{i, h}

// 		top := stack.Front()

// 		// 遇到更高的 弹出并计算雨水 直到更矮或相等
// 		bot := 0
// 		for top != nil && h > top.Value.(*Height).H {
// 			stack.Remove(top)
// 			// 计算雨水
// 			area := (top.Value.(*Height).H - bot) * (i - top.Value.(*Height).I - 1)
// 			res += area
// 			fmt.Println(area, i, h)

// 			// 下次计算雨水的底
// 			bot = top.Value.(*Height).H
// 			top = stack.Front()
// 		}

// 		// 更矮或者相等的放入
// 		if top == nil || h <= top.Value.(*Height).H {
// 			if top != nil {
// 				// 算一次雨水
// 				area := (h - bot) * (i - top.Value.(*Height).I - 1)
// 				res += area
// 				fmt.Println(area, i, h)
// 			}
// 			stack.PushFront(height)
// 		}
// 	}

// 	return res
// }
