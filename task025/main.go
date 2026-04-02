// Задача 5 – Срезы и race condition с append
// Исправьте код чтобы он работал корректно
// go run -race main.go запусти и посмотри

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	data := []int{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			data = append(data, val)
		}(i)
	}

	wg.Wait()
	fmt.Println("Length:", len(data))
	fmt.Println("Expected: 100")
}
