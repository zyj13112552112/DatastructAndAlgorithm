package main

import (
	"fmt"
	//"sort"
)
/*
	二分查找
 */

func Search(S []int,a int)int{
	l,r := 0, len(S)-1
	for{
		if l>r {
			break
		}
		mid := (l+r)/2
		res:=comp(a,S[mid])
		if res == 0 { //a==S[mid]
			return mid
 		}else if res == 1{ //a>S[mid]
			l = mid + 1
		}else if res == -1{ //a<S[mid]
			r = mid - 1
		}
	}
	return -1 //不存在
}
func comp(a,b int)int{
	if a==b {
		return 0
	} else if a < b{
		return -1
	} else{
		return 1
	}
}

func main(){
	S := make([]int,0)
	S = append(S, 1,2,4,5,6,7,8)
	i := Search(S, (3))
	fmt.Println(i)
}
