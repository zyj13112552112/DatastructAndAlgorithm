package leetcode

// 单调队列
func maxSlidingWindow(nums []int, k int) (res []int) {
	que = []int{} //存储下标
	num = nums
	for i, _ := range nums {
		push(i)
		if que[0] <= i-k {
			que = que[1:]
		}
		if i >= k-1 {
			res = append(res, num[que[0]])
		}
	}
	return res
}

var que []int
var num []int

func push(i int) {
	if len(que) == 0 {
		que = append(que, i)
		return
	}
	if num[i] > num[que[0]] {
		que = []int{i}
	} else {
		for num[i] > num[que[len(que)-1]] && len(que) != 0 {
			que = que[:len(que)-1]
		}
		que = append(que, i)
	}
}
