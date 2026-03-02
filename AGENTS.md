# LeetCode 刷题项目 - Agent 指南

本项目为 Go 语言实现的 LeetCode 算法题解仓库。以下约定用于保持代码风格一致。

## 项目结构

- 所有题解位于 `solutions/` 目录
- 每题一个独立 `.go` 文件，使用 `package solutions`
- 可选：为题目编写 `*_test.go` 测试文件（与题解同目录）
- 根目录 `main.go` 为入口，可调用 `solutions` 包进行验证

## 文件命名

- **统一使用 snake_case**：`solutions/container_with_most_water.go`、`solutions/group_anagrams.go`
- LeetCode 题号可简写：`solutions/l15.go`、`solutions/l2.go`

## 代码风格

### 函数命名

- 主解法函数：小写驼峰，与 LeetCode 函数签名一致（如 `trap`、`groupAnagrams`、`maxArea`）
- 辅助/备选实现：可加数字后缀（如 `maxArea1`、`canJump1`）

### 变量命名

- 结果变量：`res`、`result`
- 循环索引：`i`、`j`、`l`、`r`、`lo`、`hi`
- 保持简短，在算法题中可接受单字母变量

### 注释

- 使用**中文注释**说明算法思路、边界条件、关键逻辑
- 复杂题目可在文件顶部用注释描述题意或示例
- 可保留注释掉的旧实现，便于对比思路演进

### 结构体

- 辅助结构体首字母大写（如 `Height`）
- 字段名简短且含义明确（如 `I` 表示索引，`H` 表示高度）

## 测试规范

- 使用 Go 标准库 `testing` 包
- 采用**表驱动测试**，`tests` 为 `[]struct{ name, input, want }`
- 测试用例 `name` 使用**中文**描述场景（如 "空数组"、"LeetCode 示例1"、"单调递增-无积水"）
- 用 `t.Run(tt.name, ...)` 运行子测试
- 断言失败时使用 `t.Errorf` 输出输入、实际结果和期望结果

## 算法实现偏好

- 优先考虑时间/空间复杂度
- 边界条件（空数组、单元素等）提前返回
- 可保留多种解法（如暴力、优化版）便于对比
