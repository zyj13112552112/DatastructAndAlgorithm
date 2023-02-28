package main

import "fmt"

var(
	n int
	array = [100001]int{}
)
func Sort(l,r int){
	var mid int = array[(l+r)/2] //每次取中间值
	var ll,rr int = l,r
	for{
		if ll>=rr {break}
		for{if array[ll]<mid{ll++}else{break}} //左边比中间值小，继续右移
		for{if array[rr]>mid{rr--}else{break}} //右边比中间值大，继续左移
		if ll<=rr { //是否交换
			array[ll],array[rr] = array[rr],array[ll]
			ll++
			rr--
		}
	}
	if ll<r {Sort(ll,r)}
	if l<rr {Sort(l,rr)}
}
func main() {
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		fmt.Scan(&array[i])
	}
	Sort(1,n)
	for i:=1;i<=n;i++{
		fmt.Print(array[i]," ")
	}
}
