package main

import (
	"fmt"
)

// priority queue implementation
// uses a min heap as an underlying data structure

type Process struct {
	priority int
	data int
}

type PriorityQ []Process


var pq PriorityQ

func init() {
    // initialize priority queue
    pq = PriorityQ{{-1, -1}}
}

func swap(pos_i, pos_j int) {

	temp := pq[pos_i]
	pq[pos_i] = pq[pos_j]
	pq[pos_j] = temp

}

func reorderHeap(index int) {
	// base case, root
	if index == 1 {
		return
	}
	
	var parentIndex int
	
	if index % 2 == 0 {
		parentIndex = index / 2
	} else {
		parentIndex = (index - 1) / 2
	}
	
	if pq[index].priority < pq[parentIndex].priority {
		swap(index, parentIndex)
	}
	
	// call reorderHeap recursively till the root
	reorderHeap(parentIndex)
}

func Enqueue(p Process) bool {

	// dynamic insertion into the maxheap
	pq = append(pq, p)

	// reorder elements to form a heap
	reorderHeap(len(pq) - 1)

	return true
}

func heapify(index int) {
	
	// find smallest of left and right child
	var smallest int = index
	var left int = 2 * index
	var right = 2 * index
	
	if left < len(pq) && pq[left].priority < pq[smallest].priority {
		smallest = left
	}
	
	if right < len(pq) && pq[right].priority < pq[smallest].priority {
		smallest = right
	}
	
	if smallest != index {
		swap(index, smallest)
		
		// remove last element
		pq = pq[:len(pq) - 1]
		
		// recursively call heapify
		heapify(smallest)
	}

	return
}

func Dequeue() bool {
	// to dequeue swap the root with the last element
	swap(1, len(pq) - 1)
	heapify(1)
	return true
}

func main() {
	
	// test enqueue
	for i := 5; i > 0; i-- {
		Enqueue(Process{i, 0})
	}
	
	fmt.Println("After Enqueue:  ", pq)
	
	// test dequeue
	Dequeue()
	
	fmt.Println("After Dequeue: ", pq)
}

