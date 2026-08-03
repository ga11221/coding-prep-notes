package main

import "fmt"

/*
Given an n x n binary matrix grid, return the length of the shortest clear path
in the matrix. A clear path has a length of the number of visited cells and
only uses cells with value 0. The top-left cell and bottom-right cell must both
be 0. You can move 8-directionally (up, down, left, right, and the four
diagonals).

Return -1 if no such path exists.
*/

func main() {
	//grid := [][]int{{0, 1}, {1, 0}}
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Printf("shortest path for grid: %v is %v", grid, shortestPathBinaryMatrix(grid))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 || grid[len(grid)-1][len(grid[0])-1] != 0 {
		return -1
	}
	if len(grid) == 1 {
		return 1
	}
	return _shortestPath(grid, 0, 0)
}

func _shortestPath(grid [][]int, i, j int) int {
	queue := [][]int{}
	visited := map[[2]int]uint{[2]int{i, j}: 1}
	queue = append(queue, neighbors(grid, i, j, &visited)...)
	pathLength := 1
	x := len(grid) - 1
	for len(queue) > 0 {
		pathLength++
		nextLevel := [][]int{}
		for i := 0; i < len(queue); i++ {
			neighbor := queue[i]
			if neighbor[0] == x && neighbor[1] == x {
				return pathLength
			}
			nextLevel = append(nextLevel, neighbors(grid, neighbor[0], neighbor[1], &visited)...)
		}
		queue = nextLevel
	}
	return -1
}

func neighbors(grid [][]int, i, j int, visited *map[[2]int]uint) [][]int {
	neighbors := [][]int{}
	var rightInbounds, downInbounds, upInbounds, leftInbounds bool
	if i+1 < len(grid) {
		downInbounds = true
		if _, ok := (*visited)[[2]int{i + 1, j}]; !ok && grid[i+1][j] == 0 {
			neighbors = append(neighbors, []int{i + 1, j})
			(*visited)[[2]int{i + 1, j}] = 1
		}
	}
	if j+1 < len(grid[0]) {
		rightInbounds = true
		if _, ok := (*visited)[[2]int{i, j + 1}]; !ok && grid[i][j+1] == 0 {
			neighbors = append(neighbors, []int{i, j + 1})
		}
		(*visited)[[2]int{i, j + 1}] = 1
	}
	if i-1 >= 0 {
		upInbounds = true
		if _, ok := (*visited)[[2]int{i - 1, j}]; !ok && grid[i-1][j] == 0 {
			neighbors = append(neighbors, []int{i - 1, j})
		}
		(*visited)[[2]int{i - 1, j}] = 1
	}
	if j-1 >= 0 {
		leftInbounds = true
		if _, ok := (*visited)[[2]int{i, j - 1}]; !ok && grid[i][j-1] == 0 {
			neighbors = append(neighbors, []int{i, j - 1})
		}
		(*visited)[[2]int{i, j - 1}] = 1
	}
	if upInbounds && rightInbounds {
		if _, ok := (*visited)[[2]int{i - 1, j + 1}]; !ok && grid[i-1][j+1] == 0 {
			neighbors = append(neighbors, []int{i - 1, j + 1})
			(*visited)[[2]int{i - 1, j + 1}] = 1
		}
	}
	if upInbounds && leftInbounds {
		if _, ok := (*visited)[[2]int{i - 1, j - 1}]; !ok && grid[i-1][j-1] == 0 {
			neighbors = append(neighbors, []int{i - 1, j - 1})
			(*visited)[[2]int{i - 1, j - 1}] = 1
		}
	}
	if downInbounds && leftInbounds {
		if _, ok := (*visited)[[2]int{i + 1, j - 1}]; !ok && grid[i+1][j-1] == 0 {
			neighbors = append(neighbors, []int{i + 1, j - 1})
			(*visited)[[2]int{i + 1, j - 1}] = 1
		}
	}
	if downInbounds && rightInbounds {
		if _, ok := (*visited)[[2]int{i + 1, j + 1}]; !ok && grid[i+1][j+1] == 0 {
			neighbors = append(neighbors, []int{i + 1, j + 1})
			(*visited)[[2]int{i + 1, j + 1}] = 1
		}
	}
	return neighbors
}
