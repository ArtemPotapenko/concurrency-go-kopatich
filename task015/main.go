// ЗАДАЧА 15: RingBuffer
// Реализуйте методы структуры ringBuffer так, что main отрабатывал без паники

package main

import (
	"fmt"
	"reflect"
)

type ringBuffer struct {
	// напишите ваш код
}

func newRingBuffer(size int) *ringBuffer {
	// напишите ваш код
	return nil
}

func (b *ringBuffer) write(v int) {
	// напишите ваш код
}

func (b *ringBuffer) close() {
	// напишите ваш код
}

func (b *ringBuffer) read() (v int, ok bool) {
	// напишите ваш код
	return

}

func main() {
	buff := newRingBuffer(3)

	for i := 1; i <= 6; i++ {
		buff.write(i)
	}

	buff.close()

	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}

	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
