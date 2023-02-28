package leetcode

import "fmt"

func dicesProbability(n int)(res []float64){
	dp := make([][]int,n+1)
	for i:=1;i<=n;i++{
		dp[i] = make([]int,6*n+1)
	}
	//投掷一枚的和出现的次数都是1
	for i:=1;i<=6;i++{
		dp[1][i]=1
	}
	//从第二次投掷开始
	for i:=2;i<=n;i++{
		//每次投掷可能为1-6
		for j:=1;j<=6;j++{
			for k:=i;k<len(dp[i]);k++{
				if k-j<=0{continue}
				dp[i][k] += dp[i-1][k-j] //投掷i次出现k的次数等于投掷i-1次出现k-j的次数
			}
		}
	}
	for i:=1;i< len(dp[n]);i++{
		if dp[n][i]!=0 {
			fmt.Println(i,dp[n][i])
			res = append(res,float64(dp[n][i])/float64(GetPow(n)) )
		}
	}
	return
}
func GetPow(n int)int{
	sum:=1
	for i:=1;i<=n;i++{
		sum*=6
	}
	return sum
}