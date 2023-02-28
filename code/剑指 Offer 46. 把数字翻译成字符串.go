package leetcode

import "strconv"

func translateNum(num int) int {
	str := strconv.Itoa(num)
	dp := make([]int, len(str)+1)
	dp[1] = 1
	dp[0] = 1
	for i:=1;i< len(str);i++{
		a,_ := strconv.Atoi(string(str[i]))
		b,_ := strconv.Atoi(string(str[i-1]))
		c := b*10+a

		if c <= 25&&b!=0 {
			dp[i+1] = dp[i-1] + dp[i]
		}else{
			dp[i+1] = dp[i]
		}
	}
	return dp[len(str)]
}