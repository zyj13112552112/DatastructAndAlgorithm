package 中等

import "math"

func divide(dividend int, divisor int) int {
	var sign bool = divisor>0&&dividend>0 || divisor<0&&dividend<0 //true为正
	res := 0
	len := 1
	temp := divisor
	if dividend<0 {dividend*=-1}
	if divisor<0 {divisor*=-1}
	for; dividend >= divisor;{
		res += len
		dividend -= divisor
		divisor=divisor<<1 //*2 log(n)
		len=len<<1 //*2
	}
	for; dividend >= temp ;{
		divisor=divisor>>1//右移
		len=len>>1
		if dividend >= divisor{
			dividend -= divisor
			res += len
		}
	}
	if !sign {res*=-1}
	if res < math.MinInt32 || res >math.MaxInt32 {return math.MaxInt32}
    return res
}
