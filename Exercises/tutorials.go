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
	// fmt.Println("Hi ", username, " how're you doing?")

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
	fmt.Println("The Sum of the First n Positive Integers:", sum_of_num)
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


func name_triangle(){
	var side_1, side_2, side_3 float64
	fmt.Println("Enter side 1 of the triangle")
	fmt.Scanln(&side_1)
	fmt.Println("Enter side 2 of the triangle")
	fmt.Scanln(&side_2)
	fmt.Println("Enter side 3 of the triangle")
	fmt.Scanln(&side_3)
	if side_1 == side_2 && side_2 == side_3{
		fmt.Println("Equilateral triangle")
	}else if side_1 == side_2 || side_1 == side_3 || side_2 == side_3{
		fmt.Println("Isosceles triangle")
	// }else if side_1 != side_2 && side_1 != side_3 && side_2 != side_3{
	// 	fmt.Println("Scalene triangle")
	}else{
		fmt.Println("Scalene triangle")
	}
}


func grade_calculator(){
	var score float64
	for {
		fmt.Print("Enter your score: ")
		_, err := fmt.Scan(&score)
		if err == nil{
			break
		}
		var discard string
		fmt.Scanln(&discard)

		fmt.Println("Enter a valid score")
	}
	if score >= 70{
		fmt.Println("Grade A")
	}else if score > 60 && score <= 69{
		fmt.Println("Grade B")
	}else if score > 50 && score <= 59{
		fmt.Println("Grade C")
	}else if score > 45 && score <= 49{
		fmt.Println("Grade E")
	}else if score <= 44{
		fmt.Println("Grade F")
	}
}


func is_leap_year(){
	var year int
	fmt.Println("Enter a year: ")
	fmt.Scan(&year)
	if year % 400 == 0{
		fmt.Println("Leap Year")
	}else if year % 100 == 0{
		fmt.Println("Not a Leap Year")
	}else if year % 4 == 0{
		fmt.Println("Leap Year")
	}else {
		fmt.Println("Not a Leap Year")
	}
}


func is_palindrome(){
	var word, reversed_word string
	fmt.Println("Enter a word: ")
	fmt.Scan(&word)
	for i := len(word) - 1; i >= 0; i--{
		reversed_word += string(word[i])
	}
	if word == reversed_word{
		fmt.Print("Palindrome")
	}else{
		fmt.Print("Not a Palindrome")
	}
}


func is_prime(number int) bool {
	if number < 2{
		return false
	}

	for i := 2; i < number; i++{
		if number % i == 0{
			return false
		}
	}

	return true

}


func prime_func(){
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)
	is_prime := is_prime(number)
	if is_prime{
		fmt.Printf("%d is a Prime Number", number)
	}else{
		fmt.Printf("%d is not a Prime Number", number)
	}
}


func reverse_order(){
	var numbers[] int
	var input int

	for {
		fmt.Print("Enter a number (use 0 to stop): ")
		fmt.Scan(&input)

		if input == 0{
			break
		}
		numbers = append(numbers, input)
	}
	fmt.Println("Slice:", numbers)

	for i := len(numbers) - 1; i >= 0; i--{
		fmt.Println(numbers[i])
	}
}


func trial() {
	var nums [5]int
	for x := 4; x >= 1; x-- {
		fmt.Println(x)
		nums[4- x] = x
	}
	fmt.Println("Your array is:", nums)
}


// A pointer is variable that stores the memory address of another variable where the data is stored, instead of storing a value itself

func pointer_1(){
	// var a = 20 // holds value 20
	// z := &a // holds the memory address of a
	// var u *int
	// u = &a
	// fmt.Println(u)
	// println(a)
	// *z = 90 // changes the value of 20 to 90
	// println(z)
	// println(a)
	// println(*z)
	// var b *int
	// fmt.Println(b)
	// if b == nil{
	// 	fmt.Println("B is nil")
	// }else{
	// 	fmt.Println("B isn't nill")
	// }
	// fmt.Println(u == z) // comparing 2 pointers
	// p := &u
	// fmt.Println(p)
	// fmt.Println(*p)
	// fmt.Println(u)
	// println(b)
	// c := 30
	// d := &c
	// println(c)
	// println(d)
	nums := [5] int {1,2,3,6,7}
	fmt.Println(nums)
	var nums_pointer [5] *int

	for val := 0; val < 5; val++{
		nums_pointer[val] = &nums[val]
	}
	fmt.Println(nums_pointer)

}

func main() {
	// message()
	// room_area()
	// positive_intgers()
	// even_odd_num()
	// vowel_consonant()
	// shapes()
	// name_triangle()
	// grade_calculator()
	// is_leap_year()
	// is_palindrome()
	// prime_func()
	// reverse_order()
	// trial()
	pointer_1()
}
