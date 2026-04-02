// Задача 7 – Указатели и «смена адреса»
// как отработает код?

package main

import "fmt"

func changePointer(p *int) {
	v := 3
	p = &v
}

func main() {
	v := 5
	p := &v
	changePointer(p)
	fmt.Println(*p) // 5
}
