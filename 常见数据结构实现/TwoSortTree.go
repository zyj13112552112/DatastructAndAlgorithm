package main

import "fmt"

type Tree struct {
	root *node
}
func (T *Tree)Insert(v int){
	if T.root==nil{
		T.root=&node{
			data:  v,
			left:  nil,
			right: nil,
			father: nil,
		}
	}else{
		T.root.Insert(v)
	}
} //插入
func (T *Tree)Delete(v int){
	if v == T.root.data{
		l := T.root.left
		r := T.root.right
		if r!=nil{
			T.root = r
			r.father = nil
			for{
				if r.left==nil{
					r.left = l
					l.father = r
					break
				}
				r = r.left
			}
		}else{
			T.root = l
			l.father = nil
		}
	}else {
		T.root.Delete(v)
	}
} //删除
func (T *Tree)MID(){
	T.root.MID()
} //中序遍历
type node struct {
	data int
	left *node
	right *node
	father *node
}
func (t *node)Insert(v int){
	if v == t.data {
		return
	}else if v> t.data{
		if t.right==nil{
			t.right=&node{
				data:  v,
				left:  nil,
				right: nil,
				father: t,
			}
		}else{
			t.right.Insert(v)
		}
	}else{
		if t.left==nil{
			t.left=&node{
				data:  v,
				left:  nil,
				right: nil,
				father: t,
			}
		}else{
			t.left.Insert(v)
		}
	}
}
func (t *node)Delete(v int){
	if t.data==v{
		temp:=t.father
		if t.left!=nil&&t.right==nil{
			if temp.left==t{
				temp.left = t.left
			}else{
				temp.right = t.left
			}
			t.left.father=temp
		}else if t.left==nil&&t.right!=nil{
			if temp.left==t{
				temp.left = t.right
			}else{
				temp.right = t.right
			}
			t.right.father=temp
		}else if t.left==nil&&t.right==nil{
			if temp.left==t{
				temp.left = nil
			}else{
				temp.right = nil
			}
		}else if t.left!=nil&&t.right!=nil{
			templeft := t.left
			tempRight := t.right
			if temp.left==t{
				temp.left = tempRight
			}else{
				temp.right = tempRight
			}
			tempRight.father=temp
			for{
				if tempRight.left==nil{
					tempRight.left = templeft
					templeft.father = tempRight
					break
				}
				tempRight = tempRight.left
			}
		}
	}else{
		if v>t.data{
			t.right.Delete(v)
		}else{
			t.left.Delete(v)
		}
	}
}
func (t *node)MID(){
	if t.left!=nil{
		t.left.MID()
	}
	fmt.Print(t.data," ")
	if t.right!=nil{
		t.right.MID()
	}
}
func main(){
	tree := &Tree{}
	tree.Insert(10)
	tree.Insert(-1)
	tree.Insert(9)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(23)
	tree.Insert(13)
	tree.Insert(-2)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(99)
	tree.MID()
}



















































































































































































