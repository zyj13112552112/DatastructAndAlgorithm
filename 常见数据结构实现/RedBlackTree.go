package main

import (
	"fmt"
	"math/rand"
)
const (
	red = 0
	black = 1
)
//删除todo
type Color int
//红黑树节点
type RedBlackNode struct {
	data int
	color Color
	l *RedBlackNode
	r *RedBlackNode
	f *RedBlackNode
}
//红黑树
type RedBlackTree struct {
	root *RedBlackNode
}
//红黑树插入
func (tree *RedBlackTree)INSERT(v int){
	if tree.root==nil {
		tree.root = &RedBlackNode{
			data:  v,
			color: black,
			l:     nil,
			r:     nil,
			f:     nil,
		}
		return
	}else{
		tree.root = INSERT(tree.root,v)
		tree.root.f = nil
		tree.root.color = black
	}
}
//插入
func INSERT(root *RedBlackNode,v int)*RedBlackNode{
	if root == nil {
		return &RedBlackNode{
			data:  v,
			color: red,
			l:     nil,
			r:     nil,
			f:     nil,
		}
	}
	if root.data == v {
		return root
	}
	if root.data<v {
		root.r = INSERT(root.r,v)
		root.r.f = root

		//调整
		if root.r.r!=nil{
			if root.r.color == red && root.r.r.color == red {
				root = RBT_RR(root)
			}
		}
		if root.r.l!=nil{
			if root.r.color == red && root.r.l.color == red {
				root = RBT_RL(root)
			}
		}
	}
	if root.data>v{
		root.l = INSERT(root.l,v)
		root.l.f = root
		//调整
		if root.l.l!=nil{
			if root.l.color==red&&root.l.l.color==red{
				root = RBT_LL(root)
			}
		}
		if root.l.r!=nil{
			if root.l.color==red&&root.l.r.color==red{
				root = RBT_LR(root)
			}
		}

	}
	return root
}
//右右红红
func RBT_RR(root *RedBlackNode)*RedBlackNode{ //root.f,root,root.r黑红红
	rootR,rootL := root.r,root.l
	if rootL!=nil {
		if rootL.color == red { //叔叔节点存在且是红色，变色即可
			root.color = red
			rootR.color, rootL.color = black, black
			return root
		}
	}
	//叔叔节点是黑色或无，旋转
	rootRL := root.r.l
	root.color = red
	rootR.color = black
	rootR.l = root
	root.r = rootRL
	root.f = rootR
	if rootRL!=nil {
		rootRL.f = root
	}
	return rootR

}
//右左红红
func RBT_RL(root *RedBlackNode)*RedBlackNode{
	rootR,rootL := root.r,root.l
	if rootL!=nil {
		if rootL.color == red { //叔叔存在且为红色
			root.color = red
			root.r.color = black
			root.l.color = black
			return root
		}
	}
	//叔叔存在为黑或者不存在
	rootRL := root.r.l
	rootRLR := rootRL.r
	rootRLL := rootRL.l

	root.color = red
	rootRL.color = black

	rootRL.l = root
	root.f = rootRL
	rootRL.r = rootR
	rootR.f = rootRL

	root.r = rootRLL
	if rootRLL!=nil {
		rootRLL.f = root
	}
	rootR.l = rootRLR
	if rootRLR!=nil {
		rootRLR.f = rootR
	}
	return rootRL
}
//左左红红
func RBT_LL(root *RedBlackNode)*RedBlackNode{
	rootL,rootR:=root.l,root.r
	if rootR!=nil{
		if rootR.color == red{ //叔叔节点存在且是红色，变色，不用旋转
			root.color = red
			root.l.color = black
			root.r.color = black
			return root
		}
	}

	//叔叔不存在或存在为黑色
	rootLR := root.l.r
	root.color = red
	rootL.color = black
	root.l = rootLR
	rootL.r = root
	root.f = rootL
	if rootLR!=nil {
		rootLR.f = root
	}
	return rootL
}
//左右红红
func RBT_LR(root *RedBlackNode)*RedBlackNode{
	rootR,rootL := root.r,root.l
	if rootR!=nil{
		if rootR.color == red { //叔叔存在且为红色
			root.color = red
			root.r.color = black
			root.l.color = black
			return root
		}
	}

	//叔叔不存在或存在为黑色
	rootLR := root.l.r
	root.color = red
	rootLR.color = black

	rootLRR,rootLRL := rootLR.r,rootLR.l

	root.l = rootLRR
	if rootLRR!=nil{
		rootLRR.f = root
	}
	rootL.r = rootLRL
	if rootLRL!=nil{
		rootLRL.f = rootL
	}

	rootLR.r = root
	root.f = rootLR
	rootLR.l = rootL
	rootL.f = rootLR

	return rootLR
}
//中序
func (tree *RedBlackTree)MID(){
	f := func(root *RedBlackNode) {}
	f = func(root *RedBlackNode) {
		if root == nil {return}
		f(root.l)
		fmt.Print(root.data,"--->",root.color," ")
		f(root.r)
	}
	f(tree.root)
}
func SUM(root *RedBlackNode,a int){
	if root==nil {
		fmt.Println(a)
		return
	}
	if root.color == black {
		SUM(root.r,a+1)
		SUM(root.l,a+1)
	}else{
		SUM(root.r,a)
		SUM(root.l,a)
	}
} //验证红黑树是否正确
//红黑树删除
func (tree *RedBlackTree)Delete(v int){
	 DELETE(tree.root,v)
}
//删除
func DELETE(root *RedBlackNode,v int){
	  if root.data == v{
		  if root.r==nil&&root.l==nil{
				if root.color == red { //红节点直接删除
					if root == root.f.l { //这个节点可能是左节点，也可能是右节点
						root.f.l = nil
					}else if root == root.f.r {
						root.f.r = nil
					}
				}
				if root.color == black {
					if root == root.f.l { //要删除的是左孩子
						if root.f.r.color == red { //兄弟节点是红色,父节点必是黑色

						}
					}else if root == root.f.r { //要删除的是右孩子
						if root.f.l.color == red { //兄弟节点是红色

						}
					}
				}
		  }else if root.r==nil&&root.l!=nil{
				//要删除的节点只有左子节点，这个左子节点必是叶子节点，且为红
			  root.data = root.l.data
			  root.l = nil //交换数据删除
		  }else if root.l==nil&&root.r!=nil{
				//要删除的节点只有右子节点，这个右子结点必是叶子节点，且为红
			  	root.data = root.r.data
				  root.r = nil //交换数据，删除
		  }else{
			    //寻找后继节点
				temp := root.r
				for{
					if temp.l == nil {break}
					temp = temp.l
				}
				root.data = temp.data
				DELETE(temp,temp.data)
		  }
	  }else if root.data >v {
		  DELETE(root.l,v)
	  }else if root.data <v {
		  DELETE(root.r,v)
	  }
}
func RedBrother(root *RedBlackNode){
		if root == root.f.l { //当前是左孩子
			root.l = &RedBlackNode{
				data:  root.data,
				color: black,
				l:     nil,
				r:     nil,
				f:     root,
			}

			rootf := root.f
			brother := rootf.r

			root.r = brother.l
			brother.l.f = root

			root.data = rootf.data
			root.color = red

			rootf.data = brother.data

			rootf.r = brother.r
			brother.r.f = rootf

			//转化成黑兄红父的情况
			BlackBrotherRedFather(root.l)
		}else{ //当前是右孩子
			root.r = &RedBlackNode{
				data:  root.data,
				color: black,
				l:     nil,
				r:     nil,
				f:     root,
			}
			rootf := root.f
			brother := root.f.l

			root.l = brother.r
			brother.r.f = root

			root.color = red

			root.data = rootf.data

			rootf.data = brother.data

			rootf.l = brother.l
			brother.l.f = rootf
			//转化成黑兄红父的情况
			BlackBrotherRedFather(root.r)
		}
}
func BlackBrotherRedFather(root *RedBlackNode){ //黑兄红父
	if root == root.f.l {
		root.f.color = black
		root.f.r.color = red
		root.f.l = nil // 删除
	}else{
		root.f.color = black
		root.f.l.color = red
		root.f.r = nil
	}
}
func BlackRightRed(root *RedBlackNode){
	if root == root.f.l {
		rootf := root.f
		brother := root.f.r

		root.data = rootf.data

		root.r = brother.l
		brother.l.f = root

		rootf.data = brother.data

		rootf.r = brother.r
		brother.r.f = rootf
		rootf.r.color = black
	}else{
		rootf := root.f
		brother := rootf.l

		root.data = rootf.data
		rootf.data = brother.r.data

		brother.r = nil
	}
}
func BlackLeftRed(root *RedBlackNode){
	if root == root.f.l {
		root.data = root.f.data
		root.f.data = root.f.r.l.data
		root.f.r.l = nil
	}else{
		brother := root.f.l
		rootf := root.f

		root.data = rootf.data
		root.l = brother.r
		brother.r.f = root

		rootf.data = brother.data

		rootf.l = brother.l
		brother.l.color = black
		brother.l.f = rootf
	}
}
func BrotherNoSun(root *RedBlackNode){
	if root == root.f.l {
		root.f.r.color = red
		root.f.l = nil
	}else {
		root.f.l.color = red
		root.f.r = nil
	}
	//向上递归
	//todo
}

func main() {
	rbt := &RedBlackTree{}
	for i:=0 ; i<100000;i++{
		rbt.INSERT(rand.Intn(10000))
	}
	SUM(rbt.root,0)
}