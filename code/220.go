package 困难

//桶排，每个桶中只有一个元素true，两个则false
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := map[int]int{}

	for i,v:=range nums{
		id := getID(v,t+1)
		if _,has:=m[id];has{
			return true
		}
		if a,has:=m[id-1];has&&abs(a-v)<=t{
			return true
		}
		if a,has:=m[id+1];has&&abs(a-v)<=t{
			return true
		}
		m[id]=v
		if i>=k{
			delete(m,getID(nums[i-k],t+1))
		}
	}
	return false
}

func getID(x,w int)int{
	if x>=0{
		return x/w
	}
	return (x+1)/w-1
}

func abs(x int)int{
	if x<0{
		return -x
	}
	return x
}