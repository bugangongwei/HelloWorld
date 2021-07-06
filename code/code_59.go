package code

import "fmt"

// [[1,2,3],[8,9,4],[7,6,5]]
func GenerateMatrix(n int) [][]int {
	var (
		// 顺时针方向数组
		vers = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		// 方向下标
		v = 0
		// 结果集
		result = make([][]int, n)
		// 二维数组下标
		i, j = 0, 0
		// 二位数组的值
		num = 1
	)
	// 结果集初始化
	for k := 0; k < n; k++ {
		result[k] = make([]int, n)
	}

	for num <= n*n {
		fmt.Println("before", i, j, num, v)

		result[i][j] = num

		nextI, nextJ := i+vers[v%4][0], j+vers[v%4][1]
		if nextI >= n || nextI < 0 || nextJ >= n || nextJ < 0 || result[nextI][nextJ] > 0 {
			v++
			nextI, nextJ = i+vers[v%4][0], j+vers[v%4][1]
		}

		i, j = nextI, nextJ
		num++
	}

	return result
}

func generateMatrix(n int) [][]int {
	var (
		result  = make([][]int, n)
		flagMap = make(map[[2]int]bool)
		row     int
		col     int

		directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		dirIdx     int
	)

	for i := 0; i < len(result); i++ {
		result[i] = make([]int, n)
	}

	for k := 1; k <= n*n; k++ {
		result[row][col] = k // add to result
		flagMap[[2]int{row, col}] = true
		nextRow, nextCol := directions[dirIdx][0]+row, directions[dirIdx][1]+col
		if nextRow >= n || nextCol < 0 || nextCol >= n || flagMap[[2]int{nextRow, nextCol}] {
			dirIdx = (dirIdx + 1) % 4
			nextRow, nextCol = directions[dirIdx][0]+row, directions[dirIdx][1]+col
		}

		row, col = nextRow, nextCol
	}
	return result
}
