package 中等
func myPow(x float64, n int) float64 {
	if n > 0 {
		return mul(x,n)
	}
	return 1.0/mul(x,-n)
}

func mul(x float64, n int)float64 {
	mul_x := x
	ans := 1.0
	for n>0 {
		if n%2==1 {
			ans *= mul_x
		}
		mul_x*=mul_x
		n/=2
	}
	return ans
}


