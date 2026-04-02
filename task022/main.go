// Задача 2 – «processed» и потерянные сообщения
// как отработает код?

package main

import (
	"fmt"
)

func run() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"
}

func main() { run() }
