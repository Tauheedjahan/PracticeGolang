package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(item int) {
	q.data = append(q.data, item)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
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
