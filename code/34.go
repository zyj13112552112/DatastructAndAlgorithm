package 中等

import "fmt"

//在排序数组中查找元素的第一个和最后一个位置
//偏左二分，偏右二分
func searchRange(nums []int, target int) []int {
	res := [2]int{}
	r,l:=0,len(nums)-1
	if l<0 {return []int{-1,-1}}
	mid:= (r+l)/2
	//偏左二分
	for;l-r>1;{
		if nums[mid]>=target{
			l = mid
		}else if nums[mid]<target{
			r = mid+1
		}
		mid = (l+r)/2
	}
	if nums[r]!=target&&nums[l]!=target {return []int{-1,-1}}
	if nums[r]==target {res[0]=r}else{res[0]=l}//右边界
	r = res[0]
	l = len(nums) - 1
	mid = (r+l)/2
	fmt.Println(r,l)
	//靠右二分,只有等于和大于两种情况
	for;l-r>1;{
		fmt.Println(nums[mid])
		if nums[mid] > target {
			l = mid - 1
		}else if nums[mid] == target{
			r = mid + 1
			if nums[r]!=target{
				r = mid
				break
			}
		}
		mid = (l+r)/2
	}
	fmt.Println(r,l)
	if nums[l] == target {res[1]=l}else {res[1]=r}
	return res[:]
}