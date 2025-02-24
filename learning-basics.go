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

// FUNCTIONS BASICS
func greet(name string) string {
	// My first solution
	// fmt.Println("Hi", name)
	return fmt.Sprintf("Hi %s", name) // // Using Sprintf to return a string instead of printing directly
}

func calculateArea(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func findMax(nums ...int) (int, error) {
	// For error handling, "error" has to be defined in function declaration
	if len(nums) == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}
	largestNum := nums[0]
	for _, num := range nums {
		if num > largestNum {
			largestNum = num
		}
	}
	return largestNum, nil
}

func getPersonData(name string, age int, height float64) string {
	return fmt.Sprintf("Hi my name is %q\n, I am %d\n years old and my height is %.2f\n cm", name, age, height)
}

func getSliceData(nums ...int) (int, int, error) {
	if len(nums) < 0 {
		return 0, 0, fmt.Errorf("Cannot get Slice Data, slice is Empty")
	}

	sliceLength := len(nums)
	sliceSum := 0
	for _, num := range nums {
		sliceSum += num
	}
	return sliceLength, sliceSum, nil
}

func main() {
	welcomeMessage := "Hello World"
	fmt.Println(welcomeMessage)
	// personalData()
	// loopsAndIfs()
	// slicesAndMaps()

	// GO best practice:
	// 1. Do only operations in the functions and return values
	// 2. The logging and printing happens in the main() func

	greeting := greet("Robert")
	fmt.Println(greeting)

	area, perimeter := calculateArea(10, 11)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", area, perimeter)
	maxNum, err := findMax(1, 500, 100)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Largest number: %d\n", maxNum)
	}

	person := getPersonData("Robert", 31, 177)
	fmt.Println(person)

	length, sum, err := getSliceData(50, 13, 88)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("The slice is %d characters long, and this is the sum of the slice: %d\n", length, sum)
	}

}
