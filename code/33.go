package 中等

import "fmt"

func search(nums []int, target int) int {
	if len(nums)==1 {
		if nums[0] == target{
			return 0
		}else{
			return -1
		}
	}
	//二分查找，找最左边第一个小于nums[0]的数
	var r,l,mid int
	r ,l  = 0, len(nums)-1
	mid = (r+l)/2
	for;r!=l;{
		if nums[mid] >= nums[r] {//必在右边
			r = mid + 1
			if nums[r] < nums[0] {break}//nums[r]是最小的
		}else if nums[mid] <nums[r] {
			l = mid
		}
		mid = (l+r)/2
	}
	l = len(nums) - 1
	temp := r-1
	//左右再次二分
	mid = (r+l)/2
	for;l-r>1;{
		if nums[mid] >target {
			l = mid
		}else if nums[mid] <target{
			r = mid
		}else{
			return mid
		}
		mid = (r+l)/2
	}
	if nums[r] == target {return r}
	if nums[l] == target {return l}
	l = temp
	r = 0
	mid = (l+r)/2
	fmt.Println(l)
	for;l-r>1;{
		if nums[mid] >target {
			l = mid
		}else if nums[mid] <target{
			r = mid
		}else{
			return mid
		}
		mid = (r+l)/2
	}
	fmt.Println(r,l)
	if nums[r] == target {return r}
	if nums[l] == target {return l}
	return -1
}
