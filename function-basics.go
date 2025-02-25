package main

import "fmt"

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

func orderConfirmation(productName string, productPrice float64, isSuccess bool) (string, float64, bool, error) {
	if productName == "" {
		return "", 0, false, fmt.Errorf("Cannot process empty name")
	} else if productPrice == 0 {
		return "", 0, false, fmt.Errorf("Cannot process free product")
	}

	productNameString := fmt.Sprintf("The product you ordered is %s", productName)
	priceAfterTax := productPrice + (productPrice * 0.17)

	return productNameString, priceAfterTax, isSuccess, nil
}
