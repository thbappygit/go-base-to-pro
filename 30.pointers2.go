package main

import "fmt"

type Developer struct {
	Name      string
	Age       int
	Stack     string
	Salary    int
	Country   string
	City      string
	Skills    []string
	Projects  []string
	Interests []string
	Languages []string
	Hobbies   []string
	Education []string
}

func printF(data *[5]int) {
	fmt.Println(data)
}

func developerInfo(test *Developer) {

	test.Skills = append(test.Skills, "Java", "Laravel")
	fmt.Println(test.Stack, test.Skills)
}

func main() {

	examMark := 100
	addressOfExamMark := &examMark // this is the variable to take the address/pointer of examMark Variable // & is pointer operator

	valueOfTheAddress := *addressOfExamMark

	fmt.Println("Address of examMark : ", addressOfExamMark, "")
	fmt.Println("value of examMark : ", valueOfTheAddress, "")

	nArr := [5]int{1, 2, 3, 4, 5}
	printF(&nArr)

	myDev := Developer{
		Name:      "Tanvir",
		Age:       28,
		Stack:     "Go",
		Salary:    100000,
		Country:   "Bangladesh",
		City:      "Dhaka",
		Skills:    []string{"Go", "Python", "C++"},
		Projects:  []string{"Go-lang", "Python-lang", "C++-lang"},
		Interests: []string{"Gaming", "Reading", "Traveling"},
		Languages: []string{"English", "Bangla"},
		Hobbies:   []string{"Reading", "Gaming", "Traveling"},
		Education: []string{"Bachelor of Science in Computer Science"},
	}

	developerInfo(&myDev)

}
