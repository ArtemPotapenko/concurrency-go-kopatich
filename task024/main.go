// Задача 4 – GOMAXPROCS(1) и бесконечный цикл
// как отработает код?

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() { fmt.Println("Hello world!") }()

	for {
	}
}
