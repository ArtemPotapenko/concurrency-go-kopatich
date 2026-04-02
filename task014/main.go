// ЗАДАЧА 14: Run конкурентных функций
// Напишите функцию Run, которая запускает конкурентное выполнение функций fs и дожидается их окончания. Если одна или
// несколько функций из fs завершились с ошибкой, Run возвращает любую из них.

package main

import (
	"errors"
)

type fn func() error

func main() {
	expErr := errors.New("error")

	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}

	if err := Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}
}

func Run(fs ...fn) error {
	// напишите ваш код здесь
	return nil
}
