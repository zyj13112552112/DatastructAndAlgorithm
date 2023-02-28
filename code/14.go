package 中等
func cuttingRope(n int) int {
	if n<4 {return n-1}
	if n==4 {return 4}
	res := 1
	for {
		if n<=4 {break}
		res = res * 3 % 1000000007   //分为3段最乘积最大
		n -= 3
	}
	return res*n%1000000007
}


