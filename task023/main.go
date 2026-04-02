// Задача 3 – Канал-стопер done
// как отработает код?

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println("processed:", msg)
			case <-done:
				fmt.Println("stopped")
				return
			}
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"
	done <- struct{}{} // корректное завершение
}
