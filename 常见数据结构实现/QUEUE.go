package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data any
	next *Node
}
type QUEUE struct {
	start,end *Node //队首，队尾
	len int //队列长度
}
//返回一个队列
func NewQueue()QUEUE{
	return QUEUE{
		len : 0,
		start : nil,
		end : nil,
	}
}
func (que *QUEUE)PUSH(el any){
	if que.start == nil {
		que.start = new(Node)
		que.start.data = el
		que.end = que.start
	}else{
		que.end.next = new(Node)
		que.end.next.data = el
		que.end = que.end.next
	}
	que.len++
}
func (que *QUEUE)Front()any{
	if que.len == 0 {
		panic(errors.New("Queue has no element"))
	}
	que.len--
	temp := que.start.data
	que.start = que.start.next
	return temp
}
func main() {
	queue := NewQueue()
	queue.PUSH(1)
	queue.PUSH(2)
	fmt.Println(queue.Front())
	fmt.Println(queue.Front())
	fmt.Println(queue.Front())
}
