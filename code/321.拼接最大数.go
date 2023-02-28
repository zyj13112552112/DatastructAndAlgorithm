package leetcode

func maxNumber(nums1 []int, nums2 []int, k int) (res []int) {
	start := k
	if len(nums1) < k {
		start = len(nums1)
	}
	for i := start; i >= 0; i-- {
		if len(nums2) < k-i {
			continue
		}
		temp := Merge(GetMax(nums1, i), GetMax(nums2, k-i))
		if len(res) == 0 {
			res = temp
		}
		if pd(temp, res) {
			res = temp
		}
	}
	return
}
func pd(s1, s2 []int) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return s1[i] > s2[i]
		}
	}
	return false
}

// 单调栈选取最大递减数列
func GetMax(s []int, k int) (res []int) {
	for i, v := range s {
		if len(res) == 0 && len(res) < k {
			res = append(res, v)
		} else {
			for len(res) > 0 && len(res)+len(s)-i > k && res[len(res)-1] < v {
				res = res[:len(res)-1]
			}
			if len(res) < k {
				res = append(res, v)
			}
		}
	}
	return res
}

// 拼接
func Merge(s1, s2 []int) (res []int) {
	res = make([]int, len(s1)+len(s2))
	for i := range res {
		if Max(s1, s2) {
			res[i] = s1[0]
			s1 = s1[1:]
		} else {
			res[i] = s2[0]
			s2 = s2[1:]
		}
	}
	return
}
func Max(a, b []int) bool { //判断选取哪个数
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return len(a) > len(b)
}
