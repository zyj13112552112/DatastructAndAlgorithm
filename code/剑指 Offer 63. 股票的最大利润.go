package leetcode

func maxProfit(prices []int) int {
	sum := 0
	cur := 0
	for i:=1;i< len(prices);i++{
		if prices[i]-prices[cur]>sum{
			sum=prices[i]-prices[cur]
		}else{
			if prices[i]<prices[cur]{
				cur = i
			}
		}
	}

	return sum
}