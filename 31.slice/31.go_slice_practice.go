package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello Slices")
	var fruitList = []string{"Apple", "Banana", "Orange"}
	fmt.Printf("Type of fruitList is  %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Grapes")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList) // ["Banana", "Orange"]

	highScores := make([]int, 5)

	highScores[0] = 865
	highScores[1] = 645
	highScores[2] = 756
	highScores[3] = 800
	highScores[4] = 532
	highScores = append(highScores, 782)

	fmt.Println(highScores, len(highScores), cap(highScores))

	sort.Ints(highScores)

	fmt.Println(highScores, sort.IntsAreSorted(highScores))

}
