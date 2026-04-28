// ЗАДАЧА 5: Worker Pool
// Необходимо написать worker pool: нужно выполнить параллельно numJobs заданий, используя numWorkers горутин,
// которые запущены единожды за время выполнения програмы.
// Для этого напишите функции worker и main.
// Функция worker:
// - на вход получает функцию для выполнения f, канал для получения аргументов jobs и канал для записи результатов results
// - читает из jobs и записывает результат выполнения f(job) в results.
// Функция main:
// - запускает функцию worker в numWorkers горутинах;
// - в качестве первого аргумента worker использует функцию multiplier;
// - пишет числа от 1 до numJobs в канал jobs;
// - читает и выводит полученные значения из канала results, паралелльно работе воркеров

package main

import (
	"fmt"
	"sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	go func() {
		for n := range jobs {
			v := f(n)
			results <- v
		}
		wg.Done()
	}()
}

const numJobs = 5
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}
	multiplier := func(x int) int {
		return x * 10
	}

	go func() {
		defer close(jobs)
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
	}()
	wg.Add(numWorkers)

	for range numWorkers {
		worker(multiplier, jobs, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for v := range results {
		fmt.Println(v)
	}

}
