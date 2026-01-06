package main

import "fmt"

func main() {

	var data []int

	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)

	data2 := data

	data = append(data, 4)
	data2 = append(data2, 5)

	data[0] = 100

	fmt.Println(data)
	fmt.Println(data2)
}
