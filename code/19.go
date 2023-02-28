package 中等

// ListNode /**
 type ListNode struct {
	     Val int
	     Next *ListNode
 }
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p *ListNode = head
	var len = 0
	for ; p!=nil;p = p.Next{
		len++
	}
	if len == 1 {return nil}
	if len == n {return head.Next}
	n = len - n + 1; //正序第n个删除
	p = head
	for i:=1;i<=n-2;i++{
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}