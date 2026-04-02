// Задача 8 – error и nil-интерфейсы
// как отработает код?

package main

import "fmt"

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error // ?
	checkErr(e1) //  ?

	var e2 *errorString // ?
	checkErr(e2)        //  ?

	e2 = &errorString{} //  ?
	checkErr(e2)        //  ?

	var e3 *errorString = nil // ?
	checkErr(e3)              //  ?
}
