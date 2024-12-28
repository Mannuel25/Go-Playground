package main

import (
	"fmt"
	"strings"
)

func DisplayUpper(x string) {
	fmt.Println("Original text:", x)
	fmt.Println("Revised text:", strings.ToTitle(x))
	fmt.Println("Revised text:", strings.ToUpper(x))
	
}

func rectangleOperations(length, breadth float32) (float32, float32) {
	area := length * breadth
	perimeter := 2 * (length + breadth)
	return area, perimeter
}


// variadic function -> a function with variable parameters
func sumN(numbers ... int){
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("The sum of all numbers is: %v\n", sum)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
	fmt.Println("...", n * factorial(n- 1))
    return n * factorial(n- 1)
}


func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        current := a   
        a, b = b, a+b
        return current
    }
}


func main() {
	// a := "elizabeth"
	// DisplayUpper(a)
	// area, perimeter := rectangleOperations(2.35,9.22)
	// fmt.Println("area:", area)
	// fmt.Println("perimeter:", perimeter)
	// fmt.Println("Hey!")
	// sumN(2,3,3,4,221,5,5)
	// sumN(2,19,929,12234,2,4,24)
	// b := 5
    // fmt.Println(factorial(b))

	// using a function as a value
	sumNums := func() int {
		nums, sum := [2]int{2,9}, 0
		for _, value := range(nums) {
			sum += value
		}
		return sum
	}
	fmt.Println("The sum of both numbers is",sumNums())
	y, z := 10, 100
	add := func() int {	// a closure
		return y + z
	}
	fmt.Println(add())
	fib := fibonacci()
	// print the first 10 fibonacci numbers
    for i := 0; i < 10; i++ {
        fmt.Println(fib())
    }
}

