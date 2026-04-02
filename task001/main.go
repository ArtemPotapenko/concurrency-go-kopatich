// ЗАДАЧА 1: merge и fillChan
// Напишите функции merge и fillChan.
//
// Функция fillChan:
// - на вход получает целое число n
// - возвращает канал
// - пишет в этот канал n чисел от 0 до n-1
//
// Функция merge:
// - полуает на вход массив каналов cs
// - возвращает канал
// - параллельно читает из каждого канала из cs и пишет полученное значение в возвращаемый канал
//
// Ожидаемый результат: все числа из a - [0, 1], b - [0, 1, 2] и c - [0, 1, 2, 3]

package main

import (
	"fmt"
	"sync"
)

// merge - соединяет каналы в один
func merge(cs ...<-chan int) <-chan int {
	mergeChan := make(chan int)
	var wg sync.WaitGroup
	go func() {
		defer close(mergeChan)
		for _, in := range cs {
			wg.Add(1)
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					mergeChan <- v
				}
			}(in)
		}
		wg.Wait()
	}()
	return mergeChan
}

// fillChan - заполняет канал числами от 0 до n-1
func fillChan(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := range n {
			c <- i
		}
	}()
	return c
}

func main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)
	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}
}
