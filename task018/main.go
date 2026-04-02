// ЗАДАЧА 18: Produce и Main с сигналами
// Напишите функции produce и main.
// Функция produce:
// - на вход получает канал pipe
// - бесконечно пишет целые числа в канал pipe, начиная с 0
// - по сигналу от main должна завершать работу
// - при завершении должна заснуть на 3 секунды, после чего напечатать "produce finished"
// Функция main:
// - должна создать канал pipe
// - запустить produceCount функций produce и начать читать из канала pipe, печатая каждое число
// - при получении числа produceStop из pipe должна перестать читать новые числа из канала и должна отправить
//   сигнал в функции produce, заверщающий их работу
// - должна дождаться всех сообщений "produce finished" и напечатать "main finished"
// Для реализации требований допускается добавить дополнительные аргументы в функцию produce
//
// Последние 4 строки вывода программы:
// produce finished
// produce finished
// produce finished
// main finished

package main

import (
	"sync"
)

const (
	produceCount = 3
	produceStop  = 10
)

func produce(pipe chan<- int, wg *sync.WaitGroup) { // допускается добавить доп. аргументы
	// напишите свой код здесь

}

func main() {
	// напишите свой код здесь

}
