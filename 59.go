package algo

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxInQueue(nums []int, size int) []int {
	if len(nums) == 0 {
		return nil
	}

	result := make([]int, 0)
	for i := 0; i <= len(nums)-size; i++ {
		h := &IntHeap{}
		heap.Init(h)
		for j := i; j < i+size; j++ {
			heap.Push(h, nums[j])
		}
		result = append(result, heap.Pop(h).(int))
	}
	return result
}
