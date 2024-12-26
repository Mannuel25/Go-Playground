package main

import (
	"fmt"
	"strconv"
	"math"
	"math/rand"
	// "time"
	"strings"
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
			// fmt.Printf("%d is an even number\n", i)
			goto MESSAGE1
		}else {
			// fmt.Printf("%d is an odd number\n", i)
			goto MESSAGE2
		}
	}
	MESSAGE1:
		fmt.Println("This is an even number")
	MESSAGE2: 
		fmt.Println("This is an odd number")
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
	fmt.Printf("quantity data type: %T\n", quantity)
	// strconv.ParseFloat, strconv.Itoa
	fmt.Println(math.Pow(float64(int_first_number), 3))
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(11)) // random number from 0 - 10
	num := 23.4567

	// Round to 2 decimal places
	rounded := math.Round(num*100) / 100

	// Store in a float variable
	var roundedFloat float64 = rounded

	fmt.Println("Rounded float:", roundedFloat)
}

// static-typed variable
var name string = "Emmanuel"
var is_cold = false


func validateGuess(user_guess int){
	computer_number := rand.Intn(11)
	// check if the user's guess is correct
	if user_guess == computer_number{
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect. The number was", computer_number)
	}
}

func randomNumberGame(){
	for {
		fmt.Print("Guess a random number from 1 - 10: ")
		var user_number int
		_, err := fmt.Scanln(&user_number)
		if err != nil {
			fmt.Println("Invalid input, kindly enter an integer from 1 - 10")
			// clear the buffer to prevent infinite loop on invalid input
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if (user_number < 0 || user_number > 10){
			fmt.Println("Invalid input, enter a random number from 1 - 10")
			randomNumberGame()
		} else{
			validateGuess(user_number)
			break
		}
		}
}


func generateTable(number int){
	// display the multiplication table
	fmt.Printf("Displaying Multiply Table %d\n", number)
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number * i)
	}
}

func switchCase() {
	// color := "blue"
	color, number := "green", 6
	// switch (color) {
	// 	case "red":
	// 		fmt.Println("It's red not blue")
	// 	case "blue":
	// 		fmt.Println("It's blue")
	// 	default:
	// 		fmt.Println("Uhmmm..no match found") // executes when no case matches
	// }

	switch (number) {
		case 7:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough // ensures the execution continues to the next case
		case 6:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough // ensures the execution continues to the next case
		case 5:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough 
		case 4:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough 
		case 3:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough 
		case 2:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough 
		case 1:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough 
		case 0:
			fmt.Println("...", number, "...")
			number -= 1
			fallthrough 
		default:
			fmt.Println("no match found")
	}

	// using multiple expressions in a case
	switch (color) {
		case "red", "blue", "yellow":
			fmt.Println(color, "is a primary color")
		case "orange", "green", "violet":
			fmt.Println(color, "is a secondary color")
		default:
			fmt.Println(color, "is not a primary or secondary color.")
	}

}


func multiplicationTable(){
	// ask the user for a number
	// prints out multiplication table for that number
	for {
		fmt.Printf("Enter a number to generate its multiplication table: ")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Invalid, enter a valid integer")
			var discard string
			fmt.Scanln(&discard)
		}
		// call function to display the multiplication table
		generateTable(number)
		break
	}
}


func reverseString() {
	// Ask the user for a word to reverse
	var word string
	reversedWord := ""
	fmt.Print("Enter a word to reverse: ")
	fmt.Scanln(&word)

	// Iterate through the string in reverse order
	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}

	// Print the reversed word
	fmt.Println("Reversed word:", reversedWord)
}

func urlShortner(){
	// ask the user for a URL to shorten
	fmt.Printf("Enter a URL to shorten: ")
	var url, identifier string

	fmt.Scanln(&url)
	// get the last part of the url
	url_parts := strings.Split(url, "/")
	page_number := url_parts[len(url_parts)-1]
	if len(page_number) == 1{
		identifier = string(page_number[0]) + string(page_number[0])
	} else if len(page_number) >= 2{
		identifier = string(page_number[0]) + string(page_number[len(page_number)-1])
	} else {
		identifier = "mm" // dummy identifier
	}
	random_number := rand.Intn(99)
	shortened_url := fmt.Sprintf("https://mannuel.com/%s/%d", identifier, random_number)
	fmt.Println("Your shortened URL is:", shortened_url)

}


func main() {
	fmt.Println("Hello, world!")
	// multiplicationTable()
	// switchCase()
	// reverseString()
	urlShortner()

	// displayGrreetings()
	// message()
	// loopFunc()
	// fmt.Println("My name is", name)

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
	// randomNumberGame()
	// mannually handle seeding
	// ns := rand.NewSource(time.Now().UnixNano())
	// generator := rand.New(ns)
	// fmt.Println(generator.Intn(100))
	// fmt.Println(rand.Intn(100))

}


/*
	ways to run the file: (1) complile and run --> go build gotest.go, then gotest, (2) run directly --> go run gotest.go
	to format my codes --> go fmt gotest.go
	:= --> short variable declaration, var age int = 10 --> explicit declaration
*/

