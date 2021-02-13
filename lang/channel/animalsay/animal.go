package main

import (
	"fmt"
	"sync"
)

var ab = make(chan struct{}, 0)
var bc = make(chan struct{}, 0)
var ca = make(chan struct{}, 0)

var wg = sync.WaitGroup{}

func cat() {
	i := 0
	for {
		select {
		case <-ca:
			i++
			if i > 100 {
				wg.Done()
				return
			}
			fmt.Println("cat")
			ab <- struct{}{}
		}
	}
}

func dog() {
	for {
		select {
		case <-ab:
			fmt.Println("dog")
			bc <- struct{}{}
		}
	}
}

func fish() {
	for {
		select {
		case <-bc:
			fmt.Println("fish")
			ca <- struct{}{}
		}
	}
}

func main() {
	wg.Add(1)
	go cat()
	go dog()
	go fish()

	ca <- struct{}{}

	wg.Wait()
}
