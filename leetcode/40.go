package 中等

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var m = map[int]int{}
	var c1 = []int{}

	sort.Ints(candidates)
	m[candidates[0]]=1
	c1 = append(c1,candidates[0])
	for i:=1;i< len(candidates);i++ {
		if candidates[i] == candidates[i-1] {
			m[candidates[i]]++
		} else {
			m[candidates[i]] = 1
			c1 = append(c1, candidates[i])
		}
	}

	var ans = [][]int{}

	var f = func(idx,tar int,t []int) {}

	f = func(idx, tar int, t []int) {
		if tar == 0 {
			ans= append(ans, append([]int(nil),t...))
			return
		}
		for i:=idx;i<len(c1);i++{
			for j:=1;j<=m[c1[i]];j++{ //当前数在序列中有几次
				if tar-j*c1[i]>=0{
					for k:=1;k<=j;k++{
						t = append(t, c1[i])
					}
					f(i+1,tar-j*c1[i],t)
					t = t[:len(t)-j]
				}
			}
		}
	}
	f(0,target,[]int{})
	return ans
}