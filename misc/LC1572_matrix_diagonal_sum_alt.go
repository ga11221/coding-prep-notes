package main

import "fmt"

func main() {
	//mat := [][]int{
	//{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	//}
	//mat := [][]int{
	//{1, 1, 1, 1},
	//{1, 1, 1, 1},
	//{1, 1, 1, 1},
	//{1, 1, 1, 1},
	//}
	mat := [][]int{
		{5},
	}
	fmt.Printf("sum for matrix: %v is %v", mat, diagonalSum(mat))
}

func diagonalSum(mat [][]int) int {
	length := len(mat)
	return sum("UL->LR", mat, 0, 0) + sum("UR->LL", mat, 0, length-1)
}

func sum(direction string, mat [][]int, i, j int) int {
	if i == len(mat) {
		return 0
	}
	if direction == "UR->LL" && len(mat)%2 == 1 {
		if i == j {
			i++
			j--
			return sum(direction, mat, i, j)
		}
	}
	var nexti, nextj = i, j
	if direction == "UL->LR" {
		i++
		j++
	} else {
		i++
		j--
	}
	return mat[nexti][nextj] + sum(direction, mat, i, j)
}
