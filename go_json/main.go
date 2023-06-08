package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func marshal() {
	/*
		marshal
		object -> byte
	*/

	s := Student{
		Name: "park",
		Age:  20,
	}
	// object -> byte
	ret, _ := json.Marshal(s)
	// byte -> string
	ret_string := string(ret)
	fmt.Println(ret_string)
}

func unmarshal() {
	/*
		unmarshal
		byte -> object
	*/
	s := Student{
		Name: "park",
		Age:  20,
	}
	// object -> byte
	ret, _ := json.Marshal(s)
	// byte -> object
	res := &Student{}
	error := json.Unmarshal(ret, res)
	if error != nil {
		panic(error)
	}
	fmt.Println(res.Age)
	fmt.Println(res.Name)
	/*
		byte -> map
		역따옴표(backtick, `) 방식, 역따옴표로 열리고 닫힌 string은 Raw string이라 한ㄷ,
	*/
	res_map := map[string]any{}

	data := []byte(`{"num":6.13, "str": ["a", "b"]}`)
	json.Unmarshal(data, &res_map)

	val := res_map["num"]
	fmt.Println(val)
}
func main() {
	/*
		marshal : object(string, struct, ...) -> byte
		unmarshal : []byte -> object(string, struct, ......)
	*/
	marshal()
	unmarshal()
}
