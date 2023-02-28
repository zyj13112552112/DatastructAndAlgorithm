package 困难

import (
	"container/heap"
	"sort"
)

//小根堆，
//大根堆只需要将存入小跟队的数取反，取出来再取反即可
type H struct {
	 sort.IntSlice
}

func (h *H)Push(v interface{}){
	h.IntSlice = append(h.IntSlice,v.(int))
}

func (h *H)Pop()interface{}{
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
type MedianFinder struct {
	MAX,MIN H
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
   return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
	Ma,Mi := &this.MAX,&this.MIN
	if len(Mi.IntSlice)==0 || num >=Mi.IntSlice[0]{
		heap.Push(Mi,num)
		if len(Ma.IntSlice)+1< len(Mi.IntSlice){
			heap.Push(Ma,-(heap.Pop(Mi)).(int))
		}
	}else{
		heap.Push(Ma,-num)
		if len(Ma.IntSlice)>len(Mi.IntSlice){
			heap.Push(Mi,-heap.Pop(Ma).(int))
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	Ma,Mi := &this.MAX,&this.MIN
	l1 := len(Ma.IntSlice)
	l2 := len(Mi.IntSlice)
	if (l1+l2)%2==0{
		return float64(Mi.IntSlice[0]-Ma.IntSlice[0])/2
	}else{
		return float64(Mi.IntSlice[0])
	}
}




