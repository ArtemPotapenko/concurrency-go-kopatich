// ЗАДАЧА 19: Once с каналами
//
// Реализуйте структуру once, функцию new и потокобезопасный метод do.
// Реализиция once и new должна использовать каналы, не используйте пакет sync.
// Функция new возвращает указатель на структуру once
// Метод do:
// - получает на вход функцию f
// - исполняет f только в том случае, если do вызывается в первый раз для этого экземпляра once. В противном случае
//   ничего не делает
//
// Функция main должна вывести call в консоль ровно один раз.

package main

import (
	"fmt"
	"sync"
)

const goroutinesNumber = 10

type once struct {
	// напишите ваш код здесь
}

func new() *once {
	// напишите ваш код здесь
	return nil
}

func (o *once) do(f func()) {
	// напишите ваш код здесь
}

func funcToCall() {
	fmt.Printf("call")
}

func main() {
	wg := sync.WaitGroup{}
	so := new()

	wg.Add(goroutinesNumber)
	for i := 0; i < goroutinesNumber; i++ {
		go func(f func()) {
			defer wg.Done()
			so.do(f)
		}(funcToCall)
	}

	wg.Wait()
}
