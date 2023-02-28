package 中等


type pair struct{
	x, y int
}
func exist(board [][]byte, word string) bool {
	direct := []pair{{1,0},{-1,0},{0,1},{0,-1}}

	h := len(board)
	w := len(board[0])

	visit := make([][]bool,h)
	for i:=range visit{
		visit[i] = make([]bool,w)
	}

	var check func(i,j,k int)bool
	check = func(i, j, k int) bool {
		if board[i][j]!=word[k] {return false}
		if k==len(word)-1 {return true}
		visit[i][j] = true
		defer func() {visit[i][j] = false}()

		for _,d:=range direct{
			if I,J:=d.x+i,d.y+j;I>=0&&J>=0&&I<=h-1&&J<=w-1&&!visit[I][J]{
				if check(I,J,k+1){
					return true
				}
			}
		}
		return false
	}

	for i,row:=range board{
		for j:=range row{
			if check(i,j,0){
				return true
			}
		}
	}
	return false
}