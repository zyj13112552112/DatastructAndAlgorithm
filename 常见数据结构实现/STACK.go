package main

import (
	"errors"
	"fmt"
)

type STACK struct {
	array []any
	len int
	cap int
}
//获取栈
func NewSTACK()STACK{
	return STACK{
		array: make([]any,1024),
		len:   0,
		cap:   1024,
	}
}
//栈扩容
func (s *STACK)Extend(){
	s.cap*=2
	temp := make([]any,s.cap)
	copy(temp,s.array)
	s.array = temp
}
//入栈
func (s *STACK)PUSH(el any){
	if s.len == s.cap {
		s.Extend()
	}
	s.array[s.len] = el
	s.len++
}
//出栈
func (s *STACK)POP()any{
	if s.len==0 {
		panic(errors.New("error : stack has no element"))
	}
	s.len--
	return s.array[s.len]
}
func main()  {
	s := NewSTACK()
	s.PUSH(1)
	s.PUSH("1")
	s.PUSH(3)
	fmt.Println(s.POP())
	fmt.Println(s.POP())
	fmt.Println(s.POP())
	fmt.Println(s.POP())
}