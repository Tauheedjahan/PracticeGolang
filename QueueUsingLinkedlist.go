package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(item int) {
	newNode := &Node{data: item}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}

	item := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func main() {
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)

	fmt.Println("Data dequeued from the queue:")
	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		fmt.Println(item)
	}
}
