package main

//给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量。
//子矩阵 x1, y1, x2, y2 是满足 x1 <= x <= x2 且 y1 <= y <= y2 的所有单元 matrix[x][y] 的集合。
//如果 (x1, y1, x2, y2) 和 (x1', y1', x2', y2') 两个子矩阵中部分坐标不同（如：x1 != x1'），那么这两个子矩阵也不同。

//链接：https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	//
	sums := make([][]int, len(matrix))

	for i, line := range matrix {
		sums[i] = make([]int, len(line))

		curLineSum := 0
		for j, val := range line {
			curLineSum += val
			sums[i][j] = curLineSum
			if i > 0 {
				sums[i][j] += sums[i-1][j]
			}
		}
	}

	res := 0
	// 每个起点
	for i1, line := range matrix {
		for j1 := range line {
			// 每个终点
			for i2 := i1; i2 < len(matrix); i2++ {
				for j2 := j1; j2 < len(line); j2++ {
					if target == calcWithSums(i1, j1, i2, j2, sums) {
						res++
					}
				}
			}
		}
	}
	return res
}

// (i1, j1) - (i2, j2)
func calcWithSums(i1, j1, i2, j2 int, sums [][]int) int {
	sum := sums[i2][j2]
	if j1 > 0 {
		sum -= sums[i2][j1-1]
	}
	if i1 > 0 {
		sum -= sums[i1-1][j2]
	}
	if i1 > 0 && j1 > 0 {
		sum += sums[i1-1][j1-1]
	}
	return sum
}
