package 中等

import "strconv"

var strs [31]string = [31]string{}
func init(){
	strs[1] = "1"
	for i:=2;i<=30;i++{
		strs[i]=""
	}
	DiGui(2)
}
func countAndSay(n int) string {
	return strs[n]
}
func DiGui(n int){
	if n==31{return}
	s := strs[n-1]
	t := 1
	num := s[0]

	for i:=1;i<len(s);i++{
		if num==s[i]{
			t++
		}else{
			strs[n] = strs[n] + strconv.Itoa(t) + strconv.Itoa(int(num))
			t = 1
			num = s[i]
		}
	}
	strs[n] = strs[n] + strconv.Itoa(t) + strconv.Itoa(int(num))

	DiGui(n+1)
}