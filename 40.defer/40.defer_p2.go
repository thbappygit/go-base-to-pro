package main

import "fmt"

func calculateVal() (result int) {
	fmt.Println("first", result)
	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}
	defer show()
	result = 5
	fmt.Println("second", result)
	return
}

func calcVal() int {
	result := 0
	fmt.Println("first", result)
	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}
	defer show()
	result = 5
	fmt.Println("second", result)
	return result
}

func main() {
	main1 := calculateVal()
	fmt.Println("main first", main1)

	main2 := calcVal()
	fmt.Println("main second", main2)
}

//Rules for both cases in defer in go
/*
  Named return values
---------------------------
1. All codes execute first
2. defer function store kora hobey magic box
3. return ->  just before return defer functions gula execute korbey
4. return korbey named variables gulor values


  Just return types
---------------------------
1. All codes execute first
2. defer function store kora hobey magic box e
3. return values are evaluated at this time , if defer values change that value, its wont change the return value
4.all defer functions gula execute korbey


*/
