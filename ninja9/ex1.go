package main

import (
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	fmt.Println("I am foo!!!")
	wg.Done()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func bar() {
	fmt.Println("I am bar!!!")
	wg.Done()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
