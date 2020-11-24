/////////////////////////////////
// Functions in Go
// Go Playground: https://play.golang.org/p/lFz6eoYWPFa
/////////////////////////////////

package main

import (
	"fmt"
	"math"
	"strings"
)

// defining a function with no parameters
func f1() {
	fmt.Println("This is f1() function")
}

// defining a function with 2 parameters, a and b
func f2(a int, b int) {
	//a and b are local to the function
	fmt.Println("Sum:", a+b)
}

// defining a function using shorthand parameters notation
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// defining a function that have one parameter of type float64 and returns a value of type float64
func f4(a float64) float64 {
	return math.Pow(a, a)
	//any statements below the return statement are never executed

}

// defining a function that have two parameters of type int and returns two values of type int
func f5(a, b int) (int, int) {
	return a * b, a + b
}

// defining a function that have one parameter of type int and returns a "named parameter"
func sum(a, b int) (s int) {
	fmt.Println("s:", s) // -> s is a variable with the zero value inside the function
	s = a + b

	// it automatically return s
	return // This is known as a "naked" return.
}

/////////////////////////////////
// Variadic Functions
// Go Playground: https://play.golang.org/p/ANNpW2SgpKw
/////////////////////////////////

// Variadic functions are functions that take a variable number of arguments.
// Ellipsis prefix (three-dots) in front of the parameter type makes a function variadic.
// The function may be called with zero or more arguments for that parameter.
// If the function takes parameters of different types, only the last parameter of a function can be variadic.

// creating a variadic function
func f1(a ...int) {
	fmt.Printf("%T\n", a) // => []int, slice of int
	fmt.Printf("%#v\n", a)
}

// variadicfunction that modifies one of the arguments passed.
func f2(a ...int) {
	a[0] = 50
}

// creating a variadic function that calculates and returns the sum and product of its arguments
func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString
}

// An anonymous function is a function which doesn’t contain any name and is declared inline using a function literal.
// Anonymous functions can be used closures.

// function that takes an int as an argument and returns another function that returns an int
func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {

	// calling variadic functions
	// a variadic function can be invoked with zero or more arguments
	f1(1, 2, 3, 4)

	f1() // a is: []int(nil)

	// an example of a well-known variadic function is append() built-in function.
	// it appends items to an existing slice and returns back the same slice.
	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	s, p := sumAndProduct(2., 5., 10.)
	fmt.Println(s, p) // -> 17 100

	info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // => Age: 35, Full Name:Wolfgang Amadeus Mozart

	// a defer statement defers or postpones the execution of a function until the surrounding function returns.

	// by deferring foo() it will execute it just before exiting the surrounding function which is main()
	defer foo()
	bar()

	// declaring an anonymous functions
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!") // calling the anonymous function

	// calling the increment function. It returns an anonymous function
	a := increment(10)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
