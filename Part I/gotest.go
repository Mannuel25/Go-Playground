package main

import (
	"fmt"
	"strconv"
)

func displayGrreetings() {
	fmt.Printf("Hello, Emmanuel 👋\n")
}

// single-line function
func message() { fmt.Println("Hello!"); fmt.Println("My name is John."); };

func loopFunc(){
	// ctr := 0
	// for ctr < 10 {
	// 	fmt.Println("ctr:", ctr)
	// 	ctr +=1
	// }

	nums := []int{1, 2, 3, 4, 5} // slice
	for index, value := range nums{
		fmt.Printf("Index is %d and Value is %d\n", index, value)
	}
	for i := 1; i < 11; i++ {
		// fmt.Println(i)
		if i % 2 == 0 {
			fmt.Printf("%d is an even number\n", i)
		}else {
			fmt.Printf("%d is an odd number\n", i)
		}
	}
}


func getUserInfo(){
	print("What's your name?: ")
	var my_name string
	fmt.Scanln(&my_name)
	msg := fmt.Sprintf("You name is %s", my_name)
	fmt.Print(msg)
	print("\nWhat's the amount?: ")
	var amount int
	fmt.Scan(&amount)
	amount_to_float := float64(amount)
	fmt.Print(amount, amount_to_float)
	fmt.Println("\nEnter your first and last name: ")
	var first_name, last_name string
	fmt.Scan(&first_name, &last_name)
	message := fmt.Sprintf("Your first name is %s and your last name is %s", first_name, last_name)
	fmt.Println(message)
	/* 
		fmt.Scan reads space-separated values from input. It stops reading when it encounters a space, tab, or newline.
		fmt.Scanln reads input until it encounters a newline (\n) character.
	*/
	full_name := first_name + " " + last_name
	fmt.Println("Full Name: ", full_name)


}


func arithmetic(){
	var first_number, second_number string
	fmt.Println("Enter two numbers: ")
	fmt.Scan(&first_number, &second_number)
	int_first_number, _ := strconv.Atoi(first_number)
	int_second_number, _ := strconv.Atoi(second_number)
	addition := int_first_number + int_second_number
	addition_msg := fmt.Sprintf("Addition of %d and %d is %d", int_first_number, int_second_number, addition)
	subtraction := int_first_number - int_second_number
	subtraction_msg := fmt.Sprintf("Subtraction of %d and %d is %d", int_first_number, int_second_number, subtraction)
	division := int_first_number / int_second_number
	division_msg := fmt.Sprintf("Division of %d and %d is %d", int_first_number, int_second_number, division)
	multiplication := int_first_number * int_second_number
	multiplication_msg := fmt.Sprintf("Multiplication of %d and %d is %d", int_first_number, int_second_number, multiplication)
	fmt.Println(addition_msg)
	fmt.Println(subtraction_msg)
	fmt.Println(multiplication_msg)
	fmt.Println(division_msg)
	quantity := 100
	fmt.Printf("quantity data type: %T", quantity)
	// strconv.ParseFloat, strconv.Itoa
}

// static-typed variable
var name string = "Emmanuel"
var is_cold = false

func main() {
	fmt.Println("Hello, world!")
	// displayGrreetings()
	// message()
	loopFunc()
	fmt.Println("My name is", name)

	// dynamic-typed variable
	// age := 10
	// fmt.Println("Age: ", age)
	// fmt.Println("is_cold: ", is_cold)
	// amount, gender := 25, "male"
	// var decimal, number float32 = 2.2222222, 1.0090
	// fmt.Println(amount, gender, number, decimal)
	// message := fmt.Sprintf("The amount is %d and the decimal is %f", amount, decimal)
	// fmt.Println(message)
	// fmt.Println("Memory Address of message is", &message)
	// getUserInfo()
	// arithmetic()
}

/*
	ways to run the file: (1) complile and run --> go build gotest.go, then gotest, (2) run directly --> go run gotest.go
	to format my codes --> go fmt gotest.go
	:= --> short variable declaration, var age int = 10 --> explicit declaration
*/

