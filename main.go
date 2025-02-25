package main

import "fmt"

// GO best practice:
// 1. Do only operations in the functions and return values
// 2. The logging and printing happens in the main() func

func main() {
	welcomeMessage := "Hello World"
	fmt.Println(welcomeMessage)

	/*
		// GO function and return values Practice
		// getPersonData
		person := getPersonData("Robert", 31, 177)
		fmt.Println(person)

		// getSliceData
		length, sum, err := getSliceData(50, 13, 88)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Printf("The slice is %d characters long, and this is the sum of the slice: %d\n", length, sum)
		}

		// returnThreeValues
		productName, productPrice, isSuccess, err2 := orderConfirmation("Vitamins", 22.70, true)
		if err2 != nil {
			fmt.Println("Error: ", err2)
		} else {
			fmt.Printf("Your Order is Confirmed!\n Product: %s\n Price: $%.2f\n Status: %t\n", productName, productPrice, isSuccess)
		} */

	// Structs and Methods - Bank Account simulation
		account := &BankAccount{
			Owner:         "Ali",
			Balance:       1000,
			AccountNumber: "ACC001",
		}
		account.Deposit(250)
		balanceAfterDeposit := account.Info()
		fmt.Println("Current balance after deposit", balanceAfterDeposit)
		account.Withdraw(100)
		balanceAfterWithdraw := account.Info()
		fmt.Println("Current balance after deposit\n", balanceAfterWithdraw)
}
