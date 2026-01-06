package main

import "fmt"

func main() {
	s := []int{1, 2, 5} //slice literal

	fmt.Println("slice :", s, "length :", len(s), "capacity :", cap(s))

	//slice declare with make
	//make function with len

	lenSlice := make([]int, 5)

	lenSlice[4] = 10

	fmt.Println(lenSlice)

	//make function with len and capacity
	sliceData := make([]int, 3, 5) //here 3 is length and 5 is capacity and[] is slice literal

	sliceData[0] = 5
	sliceData[1] = 7
	sliceData[2] = 10
	fmt.Println(sliceData)

	//empty slice or nil slice
	var emptySlice []int

	fmt.Println(emptySlice)

	emptySlice = append(emptySlice, 1, 2, 3)

	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))

}

/*
1. slice from an existing array
2. slice from a slice
3. slice literal
4. make function with len
5. make function with len and capacity
6.empty slice or nil slice
*/
