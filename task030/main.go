// Задача 10 – JSON и неэкспортируемое поле
// как отработает код?

package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Age  int    `json:"age"`
	name string `json:"name"`
}

func testData() {
	in := Data{33, "Ivan"}
	fmt.Printf("%#v\n", in) // ?

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) //  ?

	var out Data
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) //  ?
}

func main() { testData() }
