package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	counter := 0
	num := 100
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			m.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Current value:", counter)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End value:", counter)
}
