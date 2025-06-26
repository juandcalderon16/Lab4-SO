// philosophers.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

var (
	forks [N]sync.Mutex
	room  = make(chan struct{}, N-1) // solo 4 filósofos a la vez
)

func think(id int) {
	fmt.Printf("Philosopher %d is thinking\n", id)
	time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
}

func eat(id int) {
	fmt.Printf("Philosopher %d is eating\n", id)
	time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
}

func philosopher(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		think(id)
		room <- struct{}{} // intenta entrar al comedor
		forks[id].Lock()
		forks[(id+1)%N].Lock()
		eat(id)
		forks[id].Unlock()
		forks[(id+1)%N].Unlock()
		<-room // sale del comedor
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go philosopher(i, &wg)
	}
	wg.Wait()
}
