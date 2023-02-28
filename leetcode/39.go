package 中等


func combinationSum(candidates []int, target int)([][]int) {
	var ans = [][]int{}
	var dfs = func(idx,tar int,combine []int){}
	dfs = func(idx, tar int,combin []int) {
		if tar==0 {
			ans = append(ans, combin)
			return
		}
		for i:=idx;i<len(candidates);i++{
			if tar-candidates[i]>=0{
				combin = append(combin, candidates[i])
				dfs(i,tar-candidates[i],combin)
				combin = combin[:len(combin)-1]
			}
		}
	}
		dfs(0,target,[]int{})
	return ans
}
