// Задача 6 – Порядок defer
// как отработает код?

package main

import "fmt"

func main() {
	a := 0
	defer fmt.Println(a)
	defer func() {
		a++
		fmt.Println(a)
	}()
}
