package main

import (
	"fmt"
	"strings"
)

func message() {
	var username, message string
	fmt.Print("Hello, what's your name?: ")
	fmt.Scanln(&username)
	message = fmt.Sprintf("Hi %s, how're you doing?", username)
	fmt.Println(message)
	// fmt.Println("Hi %s, how're you doing?")

}

func room_area() {
	var length, breadth float64
	fmt.Print("What's the rectangle's length?: ")
	fmt.Scanln(&length)
	fmt.Print("What's the rectangle's breadth?: ")
	fmt.Scanln(&breadth)
	area := length * breadth
	// fmt.Println("The Area of the rectangle is:", area)
	fmt.Printf("The Area of the rectangle is: %.2f", area)
}

func positive_intgers() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	sum_of_num := (number * (number + 1) / 2)
	fmt.Println("The Sum of the First n PositiveIntegers:", sum_of_num)
}

func even_odd_num(){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	if (num % 2 == 0){
		fmt.Println("\n\nEven number")
	}else{
		fmt.Println("\n\nOdd number")
	}
}

func vowel_consonant(){
	var letter string
	fmt.Print("Enter a letter: ")
	fmt.Scanln(&letter)
	letter = strings.ToLower(letter)
	letters := []string{"a", "e", "i", "o", "u"}
	is_vowel := false
	for _, value := range letters{
		if value == letter{
			is_vowel = true
			break
		}
	}

	if is_vowel{
		fmt.Printf("%s is a vowel.", letter)
	}else{
		fmt.Printf("%s is a consonant.", letter)
	}

}


func shapes(){
	var side int
	for {
		fmt.Print("Enter a side: ")
		_, error := fmt.Scanln(&side)
		if error == nil{
			break
		}
		var discard string
		fmt.Scanln(&discard)

		fmt.Println("Enter a valid integer")
	}

	shapes_sides := map[int]string{
		3:  "Triangle",
		4:  "Quadrilateral",
		5:  "Pentagon",
		6:  "Hexagon",
		7:  "Heptagon",
		8:  "Octagon",
		9:  "Nonagon",
		10: "Decagon",
		11: "Hendecagon",
		12: "Dodecagon",
	}

	shape_name, exist := shapes_sides[side]
	if exist{
		fmt.Printf("A shape with %d sides is a %s", side, shape_name)
	}else{
		fmt.Print(side, " doesn't exist")
	}
}



func main() {
	// message()
	// room_area()
	// positive_intgers()
	// even_odd_num()
	// vowel_consonant()
	shapes()
}
