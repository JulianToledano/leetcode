package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

func minPathSum(grid [][]int) int {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j < len(grid[0])-1 {
				grid[i][j] += grid[i][j+1]
			}
			if j == len(grid[0])-1 && i < len(grid)-1 {
				grid[i][j] += grid[i+1][j]
			}
			if i < len(grid)-1 && j < len(grid[0])-1 {
				grid[i][j] += min(grid[i][j+1], grid[i+1][j])
			}
		}
	}
	return grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
