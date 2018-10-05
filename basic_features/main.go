package main

import (
	"strconv"
	"errors"
	"fmt"
)

func main() {
	doSomething("2.34")
	doSomething("abcd")

	checkCityForTW("Manchester")
	checkCityForTW("Paris")
}

func checkCityForTW(city string) {
	fmt.Printf("Does %s have a TW office? A: %v\n", city, isItATWOffice(city))
}

func isItATWOffice(city string) bool {
	//array definition
	twOffices := []string{"Manchester", "London", "Hamburg", "Barcelona"}

	//range
	for _, office := range twOffices {
		if office == city {
			return true
		}
 	}

 	return false
}

func doSomething(input string) {
	defer func() {fmt.Println("Thanks for using the doSomething method!! See you soon")}()
	amount, err := parseAmount(input)

	if err != nil {
		fmt.Printf("There was an issue: %s\n", err)
		return
	}

	//strings formatting
	fmt.Printf("The amount is: %.2f\n", amount)
}

//multiple returns and error handling
func parseAmount(input string) (float64, error) {
	amount, parseErr := strconv.ParseFloat(input, 64)

	if parseErr != nil {
		return 0, parseErr
	}

	if  amount < 0 {
		return 0, errors.New("amount needs to be positive")
	}

	return amount, nil
}