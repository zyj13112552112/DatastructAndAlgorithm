package leetcode

func singleNumber(nums []int) int {
	hash := map[int]int{}
	for _,v:=range nums{
		hash[v]++
	}
	for k,v:=range hash{
		if v == 1 {
			return k
		}
	}
	return 0
}