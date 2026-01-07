package main

import "fmt"

func testProg() (hello string) {
	hello = "Hello World"
	fmt.Println(hello)
	defer func() {
		hello = "Hello For Closer defer"
		fmt.Println(hello)
	}()

	defer func() {
		hello = "Hello   defer2"
		fmt.Println(hello)
	}()

	var arrData = []int{1, 2, 3, 4, 5}
	for _, value := range arrData {
		func() {
			defer fmt.Println("I love you", value)
		}()
	}
	arrData = append(arrData, 6, 7, 8)
	arrData = make([]int, 10, 11)

	arrData[2] = 3
	arrData = arrData[:3]
	fmt.Println(arrData)

	fmt.Println("sliced", arrData)

	fmt.Println("Array Length:", len(arrData))
	fmt.Println("Array Capacity:", cap(arrData))

	hello = "Hello after Closer defer"
	fmt.Println(hello)
	return

}
func main() {
	testProg()
}
