package main

import "fmt"

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
