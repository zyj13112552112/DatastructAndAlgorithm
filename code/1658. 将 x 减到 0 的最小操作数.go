package leetcode

// 将问题转化为求中间连续序列和为sum - x(最长),采用双指针，o(n)即可
func minOperations(nums []int, x int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	s -= x //中间要得到的和

	i, j := 0, 0
	sum := 0
	ans := len(nums) + 1
	for i < len(nums) && sum < s {
		sum += nums[i]
		i++
		for sum > s {
			sum -= nums[j]
			j++
		}
		if sum == s {
			ans = min(ans, j-1+len(nums)-s)
		}
	}

	if ans == len(nums)+1 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
