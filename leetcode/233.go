package 困难

//思路
//计算所有的1，只需要计算每个位数上1出现的次数，加起来即可
//考虑个，十，百。。。位置上1出现的情况
func countDigitOne(n int) int {
	var count = 0
	var m = n
	var j = 1
	var i = 10
	for ; n!=0 ;{
		p := n/10
		count += p*j //j位上1出现的次数

		t2 := m%i //尾数
		if t2-j>=j{ //尾数-j>=j , +j
			count += j
		}else if t2-j>=0&&t2-j<j{
			count += t2 - j + 1
		}


		j*=10
		i*=10
		n/=10
	}

	return count
}
