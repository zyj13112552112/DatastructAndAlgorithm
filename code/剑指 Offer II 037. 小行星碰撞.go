package leetcode


//用栈模拟即可
func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	stack = append(stack, asteroids[0])
	for _,v:=range asteroids{
		if len(stack)==0 {
			stack = append(stack,v)
		}
		//-> ->                            <- <-                          <- ->
		if stack[len(stack)-1]>0 && v>0 ||stack[len(stack)-1]<0 && v<0 ||stack[len(stack)-1]<0&&v>0{
			stack = append(stack, v)
		}else{
			//-> <-
			for{
				if len(stack)==0 {
					stack = append(stack, v)
					break
				}
				if stack[len(stack)-1]<0 {
					stack = append(stack, v)
					break
				}
				if stack[len(stack)]<v*-1 {
					stack = stack[:len(stack)-1]
				}else if stack[len(stack)]==v*-1{
					stack = stack[:len(stack)-1]
					break
				}else {
					break
				}
			}
		}
	}
	return stack
}
