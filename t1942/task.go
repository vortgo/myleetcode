package t1942

import (
	"container/heap"
)

const (
	eventInt = iota
	eventOut
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type EventHeap []event

func (h EventHeap) Len() int { return len(h) }
func (h EventHeap) Less(i, j int) bool {
	if h[i].time == h[j].time {
		return h[i].eventType > h[j].eventType
	}
	return h[i].time < h[j].time
}

func (h EventHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x interface{}) {
	*h = append(*h, x.(event))
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type event struct {
	friendNumber int
	eventType    int
	time         int
}

func smallestChair(times [][]int, targetFriend int) int {

	fTime := times[targetFriend][0]

	var events = &EventHeap{}
	heap.Init(events)

	for i := 0; i < len(times); i++ {
		if times[i][0] <= fTime {
			heap.Push(events, event{friendNumber: i, eventType: eventInt, time: times[i][0]})
		}
		if times[i][1] <= fTime {
			heap.Push(events, event{friendNumber: i, eventType: eventOut, time: times[i][1]})
		}
	}

	var friendToChair = make(map[int]int)

	var chairs = &IntHeap{}
	heap.Init(chairs)
	for i := 0; i < len(times)*2; i++ {
		heap.Push(chairs, i)
	}

	for events.Len() > 0 {
		event := heap.Pop(events).(event)

		if event.eventType == eventInt {
			selectedChair := heap.Pop(chairs).(int)
			friendToChair[event.friendNumber] = selectedChair

			if event.friendNumber == targetFriend {
				return selectedChair
			}
		}

		if event.eventType == eventOut {
			heap.Push(chairs, friendToChair[event.friendNumber])
			delete(friendToChair, event.friendNumber)
		}
	}

	return -1
}

type I int

func f() {
	var b I = 23
	println(b)
}
