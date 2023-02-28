package leetcode

func countDigitOne(n int) (count int) {
	j, i, m := 10, 1, n
	for n != 0 { //1111 -> 1110 + 1 -> 1100 + 11
		p := n / 10
		count += p * i

		w := m % j
		if w-i >= i {
			count += i
		} else if w-i >= 0 && w-i < i {
			count += w - i + 1
		}

		i *= 10
		j *= 10
		n /= 10
	}
	return
}
