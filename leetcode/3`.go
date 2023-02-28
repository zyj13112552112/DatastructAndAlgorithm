package 中等
func validateStackSequences(pushed []int, popped []int) bool {
	stack := [1001]int{}
	stack[0] = -99999
	i:=0
	j:=0 //跑pushed数组
	for _,v1:=range popped{
		for stack[i]!=v1{
			if j == len(pushed){
				break
			}
				i++
			stack[i] = pushed[j]
			j++
		}
		if stack[i] == v1{
			i--
		}
	}
	return i==0
}

