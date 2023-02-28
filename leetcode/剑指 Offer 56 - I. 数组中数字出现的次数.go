package leetcode

//相同的数异或结果为0，那么可以求出两个不同数的异或值
func singleNumbers(nums []int) []int {
	a:=0
	for _,v:=range nums{
		a^=v
	}
	t:=1
	//找出为1的位,这一位为1，说明两个不同数在这个位置不同       1  10 100 1000 ，用这个数将nums分成不同的两组
	for ;a&t==0; {
		t<<=1
	}
	a = 0
    b := 0
	for _,v:=range nums{
		if v&t==0 {
			a^=v
		}else {
			b^=v
		}
	}
	return []int{a,b}
}