package main

import "fmt"

func displayGrreetings() {
	fmt.Printf("Hello, Emmanuel 👋\n")
}

// single-line function
func message() { fmt.Println("Hello!"); fmt.Println("My name is John."); };

func loopFunc(){
	ctr := 0
	for ctr < 10 {
		fmt.Println("ctr:", ctr)
		ctr +=1
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


}

// static-typed variable
var name string = "Emmanuel"
var is_cold = false

func main() {
	fmt.Println("Hello, world!")
	displayGrreetings()
	message()
	loopFunc()
	fmt.Println("My name is", name)

	// dynamic-typed variable
	age := 10
	fmt.Println("Age: ", age)
	fmt.Println("is_cold: ", is_cold)
	amount, gender := 25, "male"
	var decimal, number float32 = 2.2222222, 1.0090
	fmt.Println(amount, gender, number, decimal)
	message := fmt.Sprintf("The amount is %d and the decimal is %f", amount, decimal)
	fmt.Println(message)
	fmt.Println("Memory Address of message is", &message)
	getUserInfo()
}

/*
	ways to run the file: (1) complile and run --> go build gotest.go, then gotest, (2) run directly --> go run gotest.go
	to format my codes --> go fmt gotest.go
	:= --> short variable declaration, var age int = 10 --> explicit declaration
*/

