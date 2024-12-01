package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"container/heap"
	"strconv"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	content, err := os.Open("1.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	left := &IntHeap{}
	right := &IntHeap{}
	heap.Init(left)
	heap.Init(right)

	re := regexp.MustCompile(`(\d+)\s+(\d+)`)

	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindStringSubmatch(text)
		left1, _  := strconv.Atoi(matches[1])
		right1, _  := strconv.Atoi(matches[2])
		heap.Push(left, left1)
		heap.Push(right, right1)
	}
	fmt.Println(day1(left, right))
}

func day1(l, r *IntHeap) int {
	total := 0
	for l.Len() > 0 {
		l1 := heap.Pop(l).(int)
		r1 := heap.Pop(r).(int)
		temp := l1-r1
		if temp > 0 {
			total += temp
		} else {
			total -= temp
		}
	}
	return total
}
