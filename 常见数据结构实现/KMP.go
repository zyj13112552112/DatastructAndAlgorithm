package main

import "fmt"
//KMP
var (
	text string
	mode string
	next = make([]int,1000007)
)
func Next(){
	l := len(mode)

	next[0] = 0
	j:=0
	for i:=1;i< l;i++{
		for;j>0&&mode[i]!=mode[j];{
			j = next[j-1]
		}
		if mode[i]==mode[j]{
			next[i] = j+1
			j++ //下一次开始比较的位置
		}
	}
}
func Compare(){
	l := len(mode)
	ll := len(text)
	j := 0
	for i:=0;i< ll;i++{
		for ;j>0&&text[i]!=mode[j];{
			j = next[j-1]
		}
		if text[i] == mode[j] {
			j++
		}
		if j == l {
			fmt.Println(i-j+2)
			j = next[j-1]
		}
	}
}
func main() {
	fmt.Scan(&text,&mode)
	Next()
	Compare()
	for i:=0;i< len(mode);i++{
		fmt.Printf("%d ",next[i])
	}
}
