package 困难

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val int
    Next *ListNode
}
//var sum = 0//节点总数
//var Stack [5001]*ListNode
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	p := head
//	for;p!=nil;p=p.Next{
////		sum++
//	}
//	p = head
//	var N *ListNode
//	var b bool = false
//	for;sum>=k;{
//		sum-=k
//		high:=0
//		for ;high!=k;{
//			high++
//			Stack[high]=p//入栈
//			p=p.Next
//		}
//		var cur *ListNode = Stack[high]
//		if !b {
//			head = cur
//			b = true
//		}else{
//			N.Next = cur
//		}
//		high--
//		for ;high!=0;{
//			cur.Next = Stack[high]
//			cur = cur.Next
//			high--
//		}
//		N = cur
//	}
//	N.Next = p
//	return head
//}
//递归，O（1）空间
func reverseKGroup(head *ListNode, k int) *ListNode {
	var myhead,dgHead,dgTail,dgNext *ListNode
	flag := false
	cur := head
	dgHead,dgTail = head,head
	j := 0
	sum :=0
	for;cur!=nil;{
		sum++;
		cur=cur.Next
	}
	cur = head
	for;cur!=nil;{
		if(sum<k){
			break
		}
		j++
		dgTail = cur
		// fmt.Println(cur.Val)
		cur = cur.Next
		if j==k {
			sum-=k
			if !flag {
				flag = true
				myhead = dgTail
			}else{
				dgNext.Next = dgTail

			}
			j = 0
			dgNext = DG(dgHead, dgTail)
			fmt.Println(dgNext.Val)
			dgHead = cur
		}
	}
	dgNext.Next = cur
	return myhead
}
func DG(head,sign *ListNode)*ListNode{
	if head == sign {return head}
	temp := head
	dg := DG(head.Next, sign)
	dg.Next = temp
	return temp
}