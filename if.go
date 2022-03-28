package main

import "fmt"

func main() {
	number1 := 17
	number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	fmt.Println(resultMessage)
	fmt.Println(number1)
	fmt.Println(number2)
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

	if number1 == number2 {
		fmt.Println("17 and 24")
	} else {
		fmt.Println("Not equal")
	}

}
