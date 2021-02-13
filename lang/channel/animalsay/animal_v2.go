package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan string
	done func()
}

func createWorker(wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan string),
		done: func() {
			wg.Done()
		},
	}

	go doWork(w)
	return w
}

func doWork(w worker) {
	for v := range w.in {
		fmt.Println(v)
		w.done()
	}
}

func sayDemo() {
	var wg sync.WaitGroup

	worker := createWorker(&wg)

	wg.Add(300)

	for i := 0; i < 100; i++ {
		worker.in <- "cat"
		worker.in <- "dog"
		worker.in <- "fish"
	}

	wg.Wait()
}

func main() {
	sayDemo()
}
