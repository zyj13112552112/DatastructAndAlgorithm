package leetcode

// 逆序对，归并即可
func reversePairs(nums []int) int {
	return S(nums, 0, len(nums)-1)
}

func S(p []int, start, end int) int {
	if start >= end {
		return 0
	}
	m := (start + end) / 2
	cnt := S(p, start, m) + S(p, m+1, end)
	temp := []int{}
	i, j := start, m+1
	for i <= m && j <= end {
		if p[i] > p[j] {
			temp = append(temp, p[j])
			j++
			cnt += (m - i + 1)
		} else {
			temp = append(temp, p[i])
			i++
		}
	}
	for ; i <= m; i++ {
		temp = append(temp, p[i])
	}
	for ; j <= end; j++ {
		temp = append(temp, p[j])
	}
	l := 0
	for I := start; I <= end; I++ {
		p[I] = temp[l]
		l++
	}
	return cnt
}
