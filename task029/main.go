// Задача 9 – Указатели в срезе и переменная цикла
// как отработает код?

package main

import "fmt"

func main() {
	numbers := []*int{}
	for i := 0; i < 5; i++ {
		numbers = append(numbers, &i)
	}
	for _, p := range numbers {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
}
