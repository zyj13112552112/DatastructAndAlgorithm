package leetcode

import "fmt"

//dfs,保存每次结果动态更新即可

var (
	dir     = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	n, m    int
	visit   [][]int
	length  [][]int
	matrix2 [][]int
	r       int
)

func longestIncreasingPath(matrix [][]int) int {
	n, m = len(matrix), len(matrix[0])
	visit = make([][]int, n)
	length = make([][]int, n)
	for i := 0; i < n; i++ {
		visit[i] = make([]int, m)
		length[i] = make([]int, m)
	}
	matrix2 = matrix
	r = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r = MAX(r, DFS(i, j))
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(length[i][j], " ")
		}
		fmt.Println()
	}
	return r
}
func DFS(I, J int) int {
	if visit[I][J] == 1 {
		return length[I][J]
	}
	length[I][J] = 1
	for i := 0; i <= 3; i++ { //四个方向
		if I+dir[i][0] >= 0 && I+dir[i][0] < n && J+dir[i][1] >= 0 && J+dir[i][1] < m &&
			matrix2[I][J] < matrix2[I+dir[i][0]][J+dir[i][1]] {
			length[I][J] = MAX(length[I][J], 1+DFS(I+dir[i][0], J+dir[i][1]))
		}
	}
	visit[I][J] = 1
	return length[I][J]
}
func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
