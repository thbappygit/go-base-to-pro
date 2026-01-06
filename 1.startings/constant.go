package main


import (
	"fmt"
	"math"
)

const name = "Tanvir"
const s string = "constant"

// Circle ফাংশন যা বৃত্তের ক্ষেত্রফল (area) ও পরিধি (circumference) রিটার্ন করবে
func Circle(radius float64) (float64, float64) {
	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius
	return area, circumference
}

func main() {
	fmt.Println(name, s)
	const n = 5000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Tan(n))

	radius := 10.0
	// ফাংশন কল করে মান গ্রহণ
	area, circumference := Circle(radius)

	// প্রিন্ট করা হচ্ছে
	fmt.Println("Radius:", radius)
	fmt.Println("Area of Circle:", area)
	fmt.Println("Circumference of Circle:", circumference)
}
