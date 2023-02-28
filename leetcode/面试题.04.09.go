package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//思路：根节点必须在子节点之前出现，进行层次遍历，队列中的元素每次可以任取
//func BSTSequences(root *TreeNode)(res [][]int){
//
//	if root == nil {
//		return [][]int{{}}
//	}
//
//	dp := []*TreeNode{root}
//
//	var dfs func(dp []*TreeNode,path []int)
//
//	dfs = func(dp []*TreeNode, path []int) {
//		if len(dp) == 0 {
//			res = append(res, append([]int(nil), path...))  //注意这里进行解包操作，不然后期修改path会改变res的值
//			return
//		}
//		size := len(dp)
//		for i:=0;i<size;i++{
//			cur := dp[0] //每次选择其中一个
//			dp = dp[1:]
//
//			path = append(path, cur.Val)
//
//			if cur.Left!=nil{
//				dp = append(dp, cur.Left)
//			}
//
//			if cur.Right!=nil{
//				dp = append(dp,cur.Right)
//			}
//
//			dfs(dp,path)
//
//			dp = dp[:size-1]
//			dp = append(dp, cur) //添加回去
//
//			path = path[:len(path)-1]
//		}
//	}
//	dfs(dp,[]int{})
//	return res
//}

func BSTSequences(root *TreeNode) (res [][]int) {
	if root == nil {
		return [][]int{{}}
	}
	dp := []*TreeNode{root}

	var dfs func(dp []*TreeNode, path []int)
	dfs = func(dp []*TreeNode, path []int) {
		if len(dp) == 0 {
			res = append(res, append([]int{}, path...)) //注意解包
			return
		}
		size := len(dp)

		for i := 0; i < size; i++ {

			cur := dp[0]
			dp = dp[1:]

			path = append(path, cur.Val)

			if cur.Left != nil {
				dp = append(dp, cur.Left)
			}

			if cur.Right != nil {
				dp = append(dp, cur.Right)
			}

			dfs(dp, path)

			dp = dp[:size-1]     //恢复dp
			dp = append(dp, cur) //添加之前的cur

			path = path[:len(path)-1]
		}

	}

	dfs(dp, []int{})

	return
}
