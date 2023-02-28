package leetcode
//所有的0都是由2*5得到的，可以用计算器试一下.
//分解出来的5肯定比2少，所以只需要统计5的个数 10--2*5 15--3*5 20--4*5
func trailingZeroes(n int)(res int){
	for i:=1;i<=n;i++{
		for x:=i;x%5==0;x/=5{ //能分解出5
			res++
		}
	}
	return res
}