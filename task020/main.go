// ЗАДАЧА 20: Or функция
// Напишите функцию or, которая возвращает канал, который закрывается либо как только какой-либо из каналов
// channels стал доступен для чтения, либо как только какой-либо из каналов channels закрыли
// В функции main написан код, который при корректной работе функции or выведет число очень близкое к секунде.

package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	// напишите ваш код здесь
	return nil
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
