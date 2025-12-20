package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "Intro", "slice"}
	fmt.Println(arr)

	//slice is a reference to an array
	//slice = pointer + length + capacity

	sl := arr[1:4] //[is a go]

	//pointer =  &arr[1] means the slice starts at index 1 of the array
	//length = 4-1 = 3 means the slice has 3 elements
	//capacity = len(arr) - 1 = 6 - 1 = 5 means the slice can grow up to 5 elements before needing to allocate new memory

	fmt.Println(sl)

	sl1 := sl[1:2]

	fmt.Println(sl1)
	fmt.Println(len(sl1)) //length
	fmt.Println(cap(sl1)) //capacity
	fmt.Println(*&sl1[0])

}
