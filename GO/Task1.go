// Description: A thread-safe queue implementation in Go using mutexes and condition variables
package main

import (
	"fmt"
	"sync"
	"time"
)

const maxSize = 10

type ThreadSafeQueue struct {
	items    []int
	front    int
	rear     int
	size     int
	lock     sync.Mutex
	notEmpty *sync.Cond
}

func NewQueue() *ThreadSafeQueue {
	q := &ThreadSafeQueue{
		items: make([]int, maxSize),
	}
	q.notEmpty = sync.NewCond(&q.lock)
	return q
}

func (q *ThreadSafeQueue) isEmpty() bool {
	return q.size == 0
}

func (q *ThreadSafeQueue) isFull() bool {
	return q.size == maxSize
}

func (q *ThreadSafeQueue) Enqueue(item int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.isFull() {
		fmt.Println("Queue is full, dropping item")
		return
	}
	q.items[q.rear] = item
	q.rear = (q.rear + 1) % maxSize
	q.size++
	fmt.Printf("Enqueued: %d\n", item)
	q.notEmpty.Signal()
}

func (q *ThreadSafeQueue) Dequeue() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.isEmpty() {
		q.notEmpty.Wait()
	}
	item := q.items[q.front]
	q.front = (q.front + 1) % maxSize
	q.size--
	fmt.Printf("Dequeued: %d\n", item)
	return item
}

func producer(q *ThreadSafeQueue) {
	for i := 0; i < 20; i++ {
		q.Enqueue(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(q *ThreadSafeQueue, id int) {
	for i := 0; i < 10; i++ {
		q.Dequeue()
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	q := NewQueue()
	go producer(q)
	go consumer(q, 1)
	go consumer(q, 2)
	time.Sleep(5 * time.Second)
}
