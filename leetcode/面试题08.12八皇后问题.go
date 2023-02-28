package leetcode

func solveNQueens(n int) [][]string {
	row = make(map[int]bool)
	line1 = make(map[int]bool)
	line2 = make(map[int]bool)
	N = n
	res = [][]string{}
	dfs(1, []string{})
	return res
}

// 规律如下
var (
	row   map[int]bool //列
	line1 map[int]bool //行+列
	line2 map[int]bool //行-列
	N     int
	res   [][]string
)

func dfs(j int, path []string) { // 第j行
	for i := 1; i <= N; i++ { //第i列
		if !row[i] && !line1[i+j] && !line2[j-i] {
			str := ""
			for k := 1; k < i; k++ {
				str += "."
			}
			str += "Q"
			for k := i + 1; k <= N; k++ {
				str += "."
			}
			path = append(path, str)
			row[i] = true
			line1[i+j] = true
			line2[j-i] = true
			if j == N {
				res = append(res, append([]string(nil), path...))
				path = path[:j-1]
				row[i] = false
				line1[i+j] = false
				line2[j-i] = false
				return
			}
			dfs(j+1, path)
			row[i] = false
			line1[i+j] = false
			line2[j-i] = false
			path = path[:j-1]
		}
	}
}
