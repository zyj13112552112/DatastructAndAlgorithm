package main

import (
	"fmt"
	"math/rand"
)

//平衡二叉树的插入和删除
type AVLTree struct {
	Root *AVLNode
}
func (tree *AVLTree)Insert(v int){
	tree.Root=Insert(tree.Root,v)
}
func (tree *AVLTree)Delete(v int){
	tree.Root=Delete(tree.Root,v)
}
type AVLNode struct {
	data int
	left * AVLNode
	right * AVLNode
	height int

}
func Max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
func GetHeight(node *AVLNode)int{
	if node == nil {
		return 0
	}else{
		return node.height
	}
}
func Insert(root *AVLNode,v int)*AVLNode{
	if root == nil {
		root = &AVLNode{
			data:   v,
			left:   nil,
			right:  nil,
			height: 1,
		}
		return root
	}
	if v == root.data {
		return root
	}
	if v>root.data { //往右插入
		root.right=Insert(root.right,v)
		root.height = Max(GetHeight(root.left),GetHeight(root.right))+1 //计算高度
		if GetHeight(root.right)-GetHeight(root.left)==2{ //需要调整
			if root.right.data<v { //右右
				root = RR(root)
			}else if root.right.data>v{ //右左
				root = RL(root)
			}
		}
	}
	if v<root.data { //往左插入
		root.left=Insert(root.left,v)
		root.height = Max(GetHeight(root.left),GetHeight(root.right))+1 //计算高度
		if GetHeight(root.right)-GetHeight(root.left)==-2{ //需要调整
			if root.left.data>v{ //左左
				root = LL(root)
			}else if root.left.data<v{ //左右
				root = LR(root)
			}
		}
	}
	return root
}
func LL(node *AVLNode)*AVLNode{
	tempL := node.left
	tempLR := node.left.right
	node.left = tempLR
	tempL.right = node
	node.height = Max(GetHeight(node.left),GetHeight(node.right)) + 1
	tempL.height = Max(GetHeight(tempL.left),GetHeight(tempL.right)) + 1
	return tempL
}
func LR(node *AVLNode)*AVLNode{
	tempL := node.left
	tempLR := node.left.right
	tempLRL := node.left.right.left
	tempLRR := node.left.right.right
	//旋转节点
	tempLR.left = tempL
	tempLR.right = node
	node.left = tempLRR
	tempL.right= tempLRL
    //重新计算高度
	node.height =  Max(GetHeight(node.left),GetHeight(node.right)) + 1
	tempLR.height =  Max(GetHeight(tempLR.left),GetHeight(tempLR.right)) + 1
	tempL.height =  Max(GetHeight(tempL.left),GetHeight(tempL.right)) + 1
	return tempLR
}
func RR(node *AVLNode)*AVLNode{
	tempR := node.right
	tempRL := node.right.left
	node.right = tempRL
	tempR.left = node
	node.height = Max(GetHeight(node.left),GetHeight(node.right)) + 1
	tempR.height = Max(GetHeight(tempR.left),GetHeight(tempR.right)) + 1
	return tempR
}
func RL(node *AVLNode)*AVLNode{
	tempR := node.right
	tempRL := node.right.left
	if tempRL == nil{
		fmt.Print("-----------")
	}
	tempRLL := node.right.left.left
	tempRLR := node.right.left.right

	tempRL.left = node
	tempRL.right = tempR
	node.right = tempRLL
	tempR.left = tempRLR

	node.height =  Max(GetHeight(node.left),GetHeight(node.right)) + 1
	tempR.height =  Max(GetHeight(tempR.left),GetHeight(tempR.right)) + 1
	tempRL.height =  Max(GetHeight(tempRL.left),GetHeight(tempRL.right)) + 1
	return tempRL
}
func (tree *AVLTree)MID(){
	//中序遍历
	var f = func(root *AVLNode) {}
	f = func(root *AVLNode) {
		if root==nil {return}
		f(root.left)
		fmt.Print(root.data," ")
		f(root.right)
	}
	f(tree.Root)
}
func Delete(root *AVLNode,v int)*AVLNode{
 	//删除
	if root.data == v{
		if root.left==nil&&root.right==nil{return nil} //没有左右孩子
		if root.left==nil&&root.right!=nil{return root.right} //只有右孩子
		if root.left!=nil&&root.right==nil{return root.left}  //只有左孩子
		//有左右孩子
		//将当前节点数据 与 当前节点的右子树的最左边的叶子节点的数据交换，删除叶子节点
		temp:=root.right
		for {
			if temp.left == nil {break}
			temp = temp.left
		}
		root.data = temp.data
		root.right = Delete(root.right,temp.data)
	}else if root.data < v {
		root.right = Delete(root.right,v)
	}else{
		root.left = Delete(root.left,v)
	}
	root.height = Max(GetHeight(root.right),GetHeight(root.left)) + 1
	if GetHeight(root.right)-GetHeight(root.left) == 2 { //右子树-左子树==2
		if GetHeight(root.right.right) > GetHeight(root.right.left) { //右子树的右子树更高
			root = RR(root)
		}else { //右子树的左子树更高
			root = RL(root)
		}
	}else if GetHeight(root.right)-GetHeight(root.left) == -2 { //左子树-右子树==2
		if GetHeight(root.left.left)>GetHeight(root.left.right){ //左子树的左子树更高
			root = LL(root)
		}else{ //左子树的右子树更高
			root = LR(root)
		}
	}
	return root
}

func (tree *AVLTree)Level(){
	S := []*AVLNode{tree.Root}
	for{
		if len(S)==0 {
			break
		}
		temp:= []*AVLNode{}
		for i:=0;i< len(S);i++{
			fmt.Print(GetHeight(S[i].left)-GetHeight(S[i].right)," ")
			if S[i].left!=nil {
				temp = append(temp, S[i].left)
			}
			if S[i].right!=nil {
				temp = append(temp, S[i].right)
			}
		}
		S = temp
		fmt.Println()
	}
}

func main() {
	tree := &AVLTree{}
	for i:=0;i<10000;i++{
		item := rand.Intn(10000)
		//fmt.Println(i)
			tree.Insert(item)
		}
	tree.Level()
}
