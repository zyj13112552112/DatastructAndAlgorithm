package leetcode

// 归并排序计算逆序对
func countSmaller(nums []int) (res []int) {
	n := len(nums)
	idx = make([]int, n)
	index = make(map[int]int)
	//index2 = make(map[int]int)
	count = make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
		//index[i] = i
		//index2[i] = i
	}
	MergeSort(nums, 0, n-1)

	for i := 0; i < n; i++ {
		res = append(res, count[index[i]])
		//fmt.Println(nums[i])
	}

	return
}

var (
	idx []int
	//index2 map[int]int //存储原来下标到最终下标的映射
	index map[int]int //存储最终下标到原来下标的映射
	count []int       //跟随着归并过程交换
)

func MergeSort(nums []int, l, r int) {
	if r-l == 0 {
		return
	}
	m := l + (r-l)/2
	MergeSort(nums, l, m)
	MergeSort(nums, m+1, r)

	temp := []int{}
	tempC := []int{}
	tempIdx := []int{}
	i, j := l, m+1
	add := 0
	now := l
	for i <= m && j <= r {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			count[i] += add
			tempC = append(tempC, count[i])
			tempIdx = append(tempIdx, idx[i])
			index[idx[i]] = now
			i++
		} else { //nums[i]>nums[j]
			temp = append(temp, nums[j])
			tempC = append(tempC, count[j])
			tempIdx = append(tempIdx, idx[j])
			index[idx[j]] = now
			add++ //i后面的那些数都比这个大
			j++
		}
		now++
	}

	for ; i <= m; i++ {
		tempIdx = append(tempIdx, idx[i])
		index[idx[i]] = now
		temp = append(temp, nums[i])
		count[i] += add
		tempC = append(tempC, count[i])
		now++
	}

	for ; j <= r; j++ {
		index[idx[j]] = now
		tempIdx = append(tempIdx, idx[j])
		temp = append(temp, nums[j])
		tempC = append(tempC, count[j])
		now++
	}
	i = l
	for _, v := range temp {
		nums[i] = v
		i++
	}
	i = l
	for _, v := range tempC {
		count[i] = v
		i++
	}
	i = l
	for _, v := range tempIdx {
		idx[i] = v
		i++
	}
}
