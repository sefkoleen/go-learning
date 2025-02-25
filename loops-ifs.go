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
