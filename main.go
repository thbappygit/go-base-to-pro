package main

import "fmt"

func main() {
	var testData []int

	testData = make([]int, 5)

	fmt.Println(testData, cap(testData))

	testData[0] = 5
	testData = append(testData, 10, 20)
	testData = append(testData, 30, 40, 50)
	newTestData := testData[4:6]

	fmt.Println(testData, newTestData)

	//myData := testData[:2]

}
