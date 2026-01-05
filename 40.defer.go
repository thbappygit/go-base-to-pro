package main

import "fmt"

func bc() {
	i := 10
	fmt.Println("First", i)
	defer fmt.Println("Second", i)
	i++
	fmt.Println("Third", i)
	defer fmt.Println("Fourth", i)

	return

}

// Named Return Values example
func namedReturnFunc(a int, b int) (response int) {
	response = a + b
	return
}

func main() {
	bc()
	sumC := namedReturnFunc(5, 7)
	fmt.Println(sumC)
	bfun()
}

func bfun() {
	for i := 0; i <= 4; i++ {
		defer fmt.Print(i, "\n")
	}
}
