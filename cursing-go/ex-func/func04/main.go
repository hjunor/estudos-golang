package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func main() {
	fmt.Println("Goroutine")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("cpu", runtime.NumCPU())
	fmt.Println("threads was", runtime.NumGoroutine())
	waitgroup.Add(2)
	fmt.Println("threads ago", runtime.NumGoroutine())

	go primary()
	go secundary()
	waitgroup.Wait()

}
func primary() {
	for i := 0; i < 100; i++ {
		println("p", i)

		time.Sleep(20000)
	}
	waitgroup.Done()
}

func secundary() {
	for i := 0; i < 100; i++ {
		println("s", i)
		time.Sleep(20000)

	}
	waitgroup.Done()
}
