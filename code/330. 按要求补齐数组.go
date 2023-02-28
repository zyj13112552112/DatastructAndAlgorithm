package leetcode

//贪心
//对于x,如果1~x-1都能被覆盖，那么加上x之后可以覆盖到1~2x-1
//所以，我们每次都需要找到最小的x
func minPatches(nums []int, n int)(res int) {
	for i,x:=0,1;x<=n;{ //i是nums数组下标索引，x表示x-1之前的数都能被覆盖，x需要被覆盖
		if i< len(nums)&&x>=nums[i]{
			x+=nums[i]
			i++
		}else{
			x*=2 //nums[i]>x ,则需要补充 ，补充一个数 ： x
			res++
		}
	}
	return
}
