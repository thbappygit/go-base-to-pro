package main

import "fmt"

type Person struct {
	FullName string
	Age      int
	Gender   string
}

// receiver function
func (person Person) printDetails() {
	var data string
	if person.Gender == "Male" {
		data = "His"
	} else {
		data = "Her"
	}

	fmt.Println("The name of the person is: ", person.FullName)

	fmt.Println(data, "Age", person.Age, " Years Old")
}

func main() {
	person1 := Person{
		FullName: "Tanvir Hossen Bappy",
		Age:      28,
		Gender:   "Male",
	}

	person2 := Person{
		FullName: "Ruhana Al Rumi",
		Age:      27,
		Gender:   "Female",
	}

	person1.printDetails()
	person2.printDetails()
}

func init() {
	fmt.Println("Welcome to struct and receiver func  learning ")
}
