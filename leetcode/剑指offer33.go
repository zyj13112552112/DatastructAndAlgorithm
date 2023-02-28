package leetcode

func verifyPostorder(p []int) bool {
	if len(p) == 1 || len(p) == 2 || len(p) == 0 {
		return true
	}
	now := p[len(p)-1]
	mid := -1
	for i, v := range p {
		if v > now {
			mid = i
			break
		}
	}
	if mid == -1 {
		return verifyPostorder(p[:len(p)-1])
	} else {
		for i := mid; i < len(p)-1; i++ {
			if p[i] < now {
				return false
			}
		}
		return verifyPostorder(p[:mid]) && verifyPostorder(p[mid:len(p)-1])
	}
}
