package 中等
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode)(res []int) {
	if root == nil {
		return nil
	}
	res = []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue)!=0 {
		temp := queue[0]
		res = append(res, temp.Val)
		queue = queue[1:]
		if temp.Left!=nil{
			queue = append(queue, temp.Left)
		}
		if temp.Right!=nil{
			queue = append(queue, temp.Right)
		}

	}
	return
}
