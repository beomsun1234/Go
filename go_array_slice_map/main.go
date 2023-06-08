package main

import (
	"fmt"
)

func array() {
	/*
		array init
	*/
	arr_int1 := [5]int{0, 1, 2, 3, 4}
	var arr_int2 [5]int

	arr_size := len(arr_int1)
	/*
		array search
	*/
	for i := 0; i < arr_size; i++ {
		fmt.Println(i, arr_int1[i])
		arr_int2[i] = i
	}
	for idx, val := range arr_int1 {
		fmt.Println(idx, val)
	}

}

func slice() {
	//슬라이스는 가변길이의 배열이다.
	/*
		slice init ----------------------------------------------
	*/
	slice_int1 := []int{1, 2, 3, 4, 5}
	slice_int2 := make([]int, 5)
	slice_int3 := []int{}
	// 확장
	slice_int1 = append(slice_int1, 3)
	fmt.Println("len slice_int1 = ", len(slice_int1)) //6
	fmt.Println("len slice_int2 = ", len(slice_int2)) //5
	fmt.Println("len slice_int3 = ", len(slice_int3)) //0
	// search
	for _, val := range slice_int1 {
		fmt.Println(val)
	}
}
func _map() {
	/*
		map init ----------------------------------------------
	*/
	map_string := map[string]int{}
	map_string["p"] = 1
	map_string["b"] = 2
	map_string["s"] = 3

	map_string2 := map[string][]string{}
	map_string2["p"] = []string{"p"}
	val := map_string2["p"]
	val = append(val, "b")
	val = append(val, "s")
	map_string2["p"] = val
	// search
	for k, v := range map_string {
		fmt.Println("key = ", k)
		fmt.Println("val = ", v)
	}
	for k, v := range map_string2 {
		fmt.Println("key = ", k)
		for _, v := range v {
			fmt.Println(v)
		}
	}
}
func main() {
	fmt.Println("----")
	array()
	slice()
	_map()
}
