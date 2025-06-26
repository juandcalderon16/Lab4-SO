// producer_consumer.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const bufferSize = 5

var (
	buffer  = make([]int, bufferSize)
	in, out int
	mutex   sync.Mutex
	empty   = make(chan struct{}, bufferSize)
	full    = make(chan struct{}, bufferSize)
)

func produceItem() int {
	return rand.Intn(100)
}

func consumeItem(item int) {
	fmt.Printf("Consumed: %d\n", item)
}

func producer() {
	for i := 0; i < 10; i++ {
		item := produceItem()
		<-empty // espera espacio libre
		mutex.Lock()
		buffer[in] = item
		fmt.Printf("Produced: %d\n", item)
		in = (in + 1) % bufferSize
		mutex.Unlock()
		full <- struct{}{}
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer() {
	for i := 0; i < 5; i++ {
		<-full // espera ítem disponible
		mutex.Lock()
		item := buffer[out]
		out = (out + 1) % bufferSize
		mutex.Unlock()
		empty <- struct{}{} // libera espacio
		consumeItem(item)
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Inicializamos los espacios vacíos del buffer
	for i := 0; i < bufferSize; i++ {
		empty <- struct{}{}
	}

	go producer()
	go consumer()
	go consumer()

	time.Sleep(5 * time.Second)
}
