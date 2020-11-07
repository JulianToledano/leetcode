package main

import "fmt"

func main() {
	islands := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(islands))
}

func dfs(grid [][]byte, i, j int) {
	ni := len(grid)
	nj := len(grid[0])

	if i < 0 || j < 0 || i >= ni || j >= nj || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	ni, nj := len(grid), len(grid[0])
	numIslands := 0
	for i := 0; i < ni; i++ {
		for j := 0; j < nj; j++ {
			if grid[i][j] == '1' {
				numIslands++
				dfs(grid, i, j)
			}
		}
	}
	return numIslands
}
