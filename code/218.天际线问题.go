package leetcode

import (
	"container/heap"
	"sort"
)

/*
优先队列

	用一个优先队列记录当前的建筑物高度，以及它们的结束位置（right），
	如果优先队列顶端的建筑物的right小于等于当前位置，就把它pop掉，
	如果队列为空，就说明当前位置天际线为0
*/
type pair struct {
	x, y int
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].y > h[j].y }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() interface{} { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func getSkyline(buildings [][]int) (res [][]int) {
	x := []int{}
	for _, v := range buildings {
		x = append(x, v[0], v[1])
	} //保存所有可能影响结果的横坐标
	sort.Ints(x) //排序
	h := hp{}
	maxh := 0 //存储高度，当高度改变时，记录结果
	index := 0
	n := len(buildings)
	for _, v := range x { //遍历横坐标
		for index < n && buildings[index][0] <= v { //l在横坐标左边的将r,h入队列
			heap.Push(&h, pair{buildings[index][1], buildings[index][2]})
			index++
		}
		for len(h) > 0 && h[0].x <= v { //r在横坐标左边的出队列
			heap.Pop(&h)
		}
		if len(h) == 0 { //当前高度为0
			if maxh == 0 {
				continue
			}
			res = append(res, []int{v, 0})
			maxh = 0
		} else if maxh != h[0].y {
			maxh = h[0].y
			res = append(res, []int{v, maxh})
		}
	}
	return
}
