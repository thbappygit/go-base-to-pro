package main

import "fmt"

type Country struct {
	name       string
	population int
	phoneCode  int
}

func main() {
	var country1 Country

	country1 = Country{
		name:       "China",
		population: 1400000000,
		phoneCode:  91,
	}
	country2 := Country{
		name:       "Bangladesh",
		population: 160000000,
		phoneCode:  880,
	}
	fmt.Println("1st country : ", country1.name, " 2nd country : ", country2.name, "")
	fmt.Println("1st country total populations : ", country1.population, " 2nd country total populations : ", country2.population, "")

}

func init() {
	fmt.Println("Welcome to struct learning")
}
