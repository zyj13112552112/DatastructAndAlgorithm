package leetcode

import (
	"fmt"
	"sort"
)

//建图,深搜，排序
func findItinerary(tickets [][]string)(res []string) {
	Map := make(map[string][]string)

	visit := make(map[Pair]int) // 标记机票是否使用

	for i := 0; i < len(tickets); i++ { //建图
		if Map[tickets[i][0]] == nil {
			Map[tickets[i][0]] = []string{}
		}
		Map[tickets[i][0]] = append(Map[tickets[i][0]], tickets[i][1])
		visit[Pair{
			start: tickets[i][0],
			end:   tickets[i][1],
		}]++   //可以有重复边
	}



	for k:=range Map{
		sort.Strings(Map[k])
		fmt.Println(Map[k])
	}//排序
	res = append(res,"JFK" )

	var f func(now string,total int)
	flag := false
	f = func(now string,total int) {
		if total == len(tickets) {
			flag = true
			return
		}
		for _,v:=range Map[now] {
			if visit[Pair{
				start: now,
				end:   v,
			}]>0{
				visit[Pair{
					start: now,
					end:   v,
				}]--
				res = append(res, v)
				f(v,total+1)
				if flag {
					return
				}
				res = res[:len(res)-1]
				visit[Pair{
					start: now,
					end:   v,
				}]++
			}
		}
	}
	f("JFK",0)
	return
}
type Pair struct {
	start,end string
}