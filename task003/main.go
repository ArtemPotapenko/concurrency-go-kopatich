// ЗАДАЧА 3: Generator и Squarer с контекстом
// Напишите функции generator и squarer.
// Функция generator принимает на вход контекст и слайс целых чисел, элементы которого последовательно записываются в
// возвращаемый канал.
// Функция squarer принимает на вход контекст и канал целых чисел. Функция последовательно читает из канал числа,
// возводит их в квадрат и пишет в возвращаемый канал.
// Обе функции должны уметь завершаться по отмене контекста.

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	pipeline := squarer(ctx, generator(ctx, 1, 3, 2, 5))
	for x := range pipeline {
		if x == 4 {
			cancel()
		}
		fmt.Println(x)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	out := make(chan int, len(in))
	go func() {
		defer close(out)
		for _, val := range in {
			select {
			case <-ctx.Done():
				return
			case out <- val:
			}
		}
	}()
	return out
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case result <- v * v:
			}
		}
	}()
	return result
}
