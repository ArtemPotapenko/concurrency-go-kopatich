// ЗАДАЧА 8: ExecuteTaskWithTimeout
// Есть функция executeTask, которая во время исполнения может зависнуть на неопределенно долгое время.
// Реализуйте функцию-обертку executeTaskWithTimeout, которая:
// - исполняет executeTask
// - принимает аргументом контекст
// - завершается либо в результате исполнения executeTask, либо в результате отмены контекста. В последнем случае
//   вернуть ошибку контекста.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 100 * time.Millisecond

func main() {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	// напишите ваш код здесь
	return nil
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
