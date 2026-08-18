type MinHeap [][2]int 

func (h MinHeap) Len() int {
	return len(h)
}
func(h MinHeap)Less(i, j int)bool{
	return h[i][0] < h[j][0]
}
func(h MinHeap)Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}
func(h *MinHeap)Push(x interface{}){
	*h = append(*h, x.([2]int))
}
func(h *MinHeap)Pop()interface{}{
	old := *h
	oldLen := len(old)
	x := old[oldLen - 1] 
	*h = old[:oldLen - 1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	freqCount := make(map[int]int)
	for _, num := range nums {
		freqCount[num]++
	}

	//Initiate Our Min-Heap
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	//iterate over our freq map to get the key,val i.e number,freq
	for num, freq := range freqCount {
		heap.Push(minHeap, [2]int{freq, num})

		if minHeap.Len() > k{
			heap.Pop(minHeap)
		}
	}

	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(minHeap).([2]int)[1]
	}

	return result
}
