package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}
	que1 := []*TreeNode{}
	que2 := []*TreeNode{}
	que1 = append(que1, root)
	i := 1
	for len(que1) != 0 {
		t := []int{}
		for len(que1) != 0 {
			temp := que1[0]
			que1 = que1[1:]
			if i%2 == 1 {
				t = append(t, temp.Val)
			} else {
				t = append([]int{temp.Val}, t...)
			}
			if temp.Left != nil {
				que2 = append(que2, temp.Left)
			}
			if temp.Right != nil {
				que2 = append(que2, temp.Right)
			}
		}
		res = append(res, t)
		que1 = que2
		que2 = []*TreeNode{}
		i++
	}
	return res
}
