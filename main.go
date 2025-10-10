package main

import (
	"fmt"
	"sync"
)

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	mu := sync.Mutex{}

	result := 999999

	var res [2]int

	go func() {
		wg.Done()
		mu.Lock()
		res[0] = n1
		if result > n1 {
			result = n1
		}
		mu.Unlock()

	}()

	go func() {
		wg.Done()
		mu.Lock()
		res[1] = n2
		if result > n2 {
			result = n2
		}
		mu.Unlock()
	}()
	wg.Wait()
	fmt.Println(result)
}
