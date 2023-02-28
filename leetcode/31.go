package 中等
/**
找到下一个排列（比当前排列大）
我们需要找到一个左边靠右的小数和一个右边的大数交换
首先从后向前查找第一个顺序对 (i,i+1)(i,i+1)，满足 a[i] < a[i+1]a[i]<a[i+1]。这样「较小数」即为 a[i]a[i]。此时 [i+1,n)[i+1,n) 必然是下降序列。
如果找到了顺序对，那么在区间 [i+1,n)[i+1,n) 中从后向前查找第一个元素 jj 满足 a[i] < a[j]a[i]<a[j]。这样「较大数」即为 a[j]a[j]。
交换 a[i]a[i] 与 a[j]a[j]，此时可以证明区间 [i+1,n)[i+1,n) 必为降序。我们可以直接使用双指针反转区间 [i+1,n)[i+1,n) 使其变为升序，而无需对该区间进行排序。
 */
func nextPermutation(nums []int) {
	var left int = -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			left = i;
			break
		}
	}
	//已经最大
	if left == -1 {
		reverse(nums[:])
	} else {
		var right int = 0
		for i := len(nums) - 1; i >= left; i-- {
			if nums[i] > nums[left] {
				right = i;
				break
			}
		}
		nums[left], nums[right] = nums[right], nums[left]
		reverse(nums[left+1:])
	}
}
//翻转
func reverse(S []int){
	for i,n:=0,len(S);i<n/2;i++{
		S[i],S[n-i-1] = S[n-1-i],S[i]
	}
}