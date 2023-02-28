package 中等
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
//var BB *TreeNode
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B==nil{
		return false
	}
	return recur(A,B)||isSubStructure(A.Left,B)||isSubStructure(A.Right,B)
}
func recur(A* TreeNode,B*TreeNode)bool{
	if B==nil {
		return true
	}
	if A==nil {
		return false
	}
	if A.Val == B.Val {
		return recur(A.Left,B.Left) && recur(A.Right,B.Right)
	}else {
		return false
	}

}