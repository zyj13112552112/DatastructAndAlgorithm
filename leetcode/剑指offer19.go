package 困难
//
//func isMatch(s string, p string) bool {
//	m,n := len(s),len(p)
//	matches := func(i,j int)bool {
//		if i==0{
//			return false
//		}
//		if p[j-1]=='.'{
//			return true
//		}
//		return s[i-1]==p[j-1]
//	}
//	f := make([][]bool,m+1)
//	for i:=0;i<len(f);i++{
//		f[i] = make([]bool,n+1)
//	}
//	f[0][0] = true
//	for i:=0;i<=m;i++{
//		for j:=1;j<=n;j++{
//			if p[j-1]=='*'{
//				f[i][j] = f[i][j] || f[i][j-2]    //匹配0次
//				if matches(i,j-1){
//					f[i][j] = f[i][j] || f[i-1][j] //匹配1次，丢掉后还能继续匹配
//				}
//			}else if matches(i,j){
//				f[i][j] = f[i][j] || f[i-1][j-1]
//			}
//		}
//	}
//	return f[m][n]
//}


func isMatch(s string, p string) bool {
	n,m := len(s), len(p)

	f := make([][]bool,n+1)
	for i:=0;i<=n;i++{
		f[i] = make([]bool,m+1)
	}

	var Func = func(i,j int) bool{
		if i==0 {return false}
		if p[j-1]=='.' {return true}
		return p[j-1] == s[i-1]
	}

	f[0][0] = true

	for i:=0;i<=n;i++{
		for j:=1;j<=m;j++{
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if Func(i,j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			}else if Func(i,j){

					f[i][j] = f[i][j] || f[i-1][j-1]

			}
		}
	}

	return f[n][m]
}