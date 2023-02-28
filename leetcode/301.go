package 困难

import "fmt"
//删除括号时剪枝
//减少重复删除 （（（（（）这种情况，删除任何一个（ 就行
var del = 0 //需要删除的括号数
var strs = []string{}
var S string
var m map[string]bool
var R = 0
var L = 0
func removeInvalidParentheses(s string) []string {
R = 0
L = 0
m = map[string]bool{}
strs = []string{}
S = ""
I := 0
del = 0
for i:=0;i< len(s);i++{
if s[i]=='('{
I++
}else if s[i]==')'{
I--
}
if I<0{ //右括号多余
R++
I = 0
}
}
L = I //左括号多余
flag := map[int]bool{}
for i:=0;i< len(s);i++{
if s[i]=='('{
break
}
if s[i]==')'{
flag[i]=true
R--
}
} // 预处理左边右括号

for i:=len(s)-1;i>=0;i--{
if s[i]=='('{
flag[i]=true
L--
}
if s[i]==')'{
break
}
}// 预处理右边左括号

for i:=0;i< len(s);i++{
if flag[i] {continue}
S += string(s[i])
}
fmt.Println(s)
s = S
dfs(s,R,L)
return strs
}

func check(str string)bool{
I := 0
for i:=0;i< len(str);i++{
if str[i]=='('{
I++
}else if str[i]==')'{
I--
}
if I<0{
return false
}
}
return I==0
}
func dfs(str string,r,l int){
if r==0&&l==0&&check(str)&&m[str]==false{
m[str] = true
strs = append(strs, str)
return
}
for i:=0;i< len(str);i++{
if i!=0 {
if str[i]==str[i-1]{
continue
}
}
if str[i]=='('{
if l>0{
temp := str[:i] + str[i+1:]
dfs(temp,r,l-1)
}
}else if str[i]==')'{
if r>0{
temp := str[:i] + str[i+1:]
dfs(temp,r-1,l)
}
}
}
}