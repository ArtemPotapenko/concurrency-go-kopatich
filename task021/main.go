// Задача 1 – Буферизированный канал и select
// как отработает код?
package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan string, 3)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c <- "Goroutine " + strconv.Itoa(i)
		}(i)
	}

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		}
	}

	wg.Wait()
	close(c)
}
