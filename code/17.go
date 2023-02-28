package 中等
//递归
var ph map[string]string = map[string]string{
	"2":"abc",
	"3":"def",
	"4":"ghi",
	"5":"jkl",
	"6":"mno",
	"7":"pqrs",
	"8":"tuv",
	"9":"wxyz",
}
var str []string
var digit string
func Dfs(i int,res string){
	if i == len(digit){
	str = append(str, res)
	return
	}
	//i是第几个数字，j是第几个字母
	var s string = ph[string(digit[i])]
	for j:=0;j< len(s);j++{
	Dfs(i+1,res+string(s[j]))
}

}
func letterCombinations(digits string) []string {
	str = []string{}
	if digits == "" {return str}
	digit = digits
	Dfs(0,"")
	return str
}
