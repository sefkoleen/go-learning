package main

import "fmt"

func personalData() {
	name := "Robert"
	age := 31

	fmt.Println(name)
	fmt.Println(age)
}

func loopsAndIfs() {
	score := 75

	if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	for i := 1; i <= score; i++ {
		fmt.Println(i)
	}
}

var testSlice []string
var testMap map[string]int

func slicesAndMaps() {
	testSlice := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(testSlice)

	// testMap := make(map[string]int)
	testMap := map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}
	value := testMap["Wednesday"]
	fmt.Println(value)

	testMap["Funday"] = 8
	fmt.Println(testMap)
}

func main() {
	welcomeMessage := "Hello World"
	fmt.Println(welcomeMessage)
	// personalData()
	// loopsAndIfs()
	slicesAndMaps()
}
