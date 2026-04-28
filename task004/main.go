// ЗАДАЧА 4: RepeatFn и Take
// Напишите функции repeatFn и take.
// Функция repeatFn бесконечно вызывает функцию fn и пишет результат ее работы в возвращаемый канал.
// Прекращает работу раньше, если контекст отменен.
// Функция take читает не более чем num из канала in, пока in открыт, и пишет значение в возвращаемый канал.
// Прекращает работу раньше, если контекст отменен.

package main

import (
	"context"
	"fmt"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			v := fn()
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()
	return ch
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{}, num)
	go func() {
		defer close(out)
		for range num {
			select {
			case v := <-in:
				out <- v
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	randi := func() interface{} { return rand.Intn(10) }

	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, randi), 3) {
		res = append(res, num)
	}

	fmt.Println(res)

	if len(res) != 3 {
		panic("wrong code")
	}
}
