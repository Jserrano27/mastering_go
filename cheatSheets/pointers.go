/////////////////////////////////
// Pointers in Go
// Go Playground: https://play.golang.org/p/hkXQnr--H17
/////////////////////////////////

package main

import "fmt"

func main() {

	// the & (ampersand) operator also known as address of operator returns the memory address of a variable.
	name := "Andrei"
	fmt.Println(&name) // -> 0xc0000101e0

	//** DECLARING AND INITIALIZING POINTERS **//

	var x int = 2
	// the expression &x means the address of x and creates a pointer to an integer variable,
	// ptr is of type *int, which is pronounced "pointer to int".
	ptr := &x
	fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, ptr, &ptr) // -> p is of type *int with value 0xc000014140.

	// declaring a pointer without initializing it
	// its zero value is nil
	var ptr1 *float64
	_ = ptr1

	// creating a pointer using new() built-in function.
	p := new(int) // it creates a pointer called p that is a pointer to an int type

	x = 100
	p = &x // initializing p

	fmt.Printf("p is of type %T with value %v\n", p, p) // => p is of type *int with value 0xc000014140
	fmt.Printf("address of x is %p\n", &x)              // => address of x is 0xc000016120

	//** THE DEREFERENCING OPERATOR **//

	// * in front of a pointer is called the dereferencing operator
	*p = 90 //equivalent to x = 90 because *p is x
	// x and *p is the same thing.

	fmt.Println(*p)                  // => 90
	fmt.Println("*p == x:", *p == x) // => *p == x: true

	fmt.Println("Value of x:", *p) // => Value of x: 90 , equivalent to fmt.Println(x)

	*p = 10        // If I write *p = 10, this is equivalent to x = 10
	*p = *p / 2    //dividing x through the pointer
	fmt.Println(x) // -> 5

	// In a nutshell:
	// &value => pointer -> if you have a value you turn it into an address or pointer by using the ampersand operator
	// *pointer => value  -> and if you have pointer you turn it into value value by using the star operator


	/////////////////////////////////
	// Passing Values and Pointers to Functions
	// Go Playground: https://play.golang.org/p/4dAWL-iWp4I
	/////////////////////////////////

	// Go is a pass by value language, and not by reference (no exception here)
	// declaring a function that takes an int, a float, a string and a bool value.
	// the function works on copy so the changes are not seen outside (pass by value)
	func changeValues(quantity int, price float64, name string, sold bool) {
		quantity = 3
		price = 500.5
		name = "Mobile Phone"
		sold = false
	}
	
	// declaring a function that takes in a pointer to int, a pointer to float, a pointer to string and a pointer to bool.
	// the function makes a copy of each pointer but they point to the same address as the originals
	func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
		//changing the values the pointers point to is seen outside the function
		*quantity = 3
		*price = 500.5
		*name = "Mobile Phone"
		*sold = false
	}
	
	// declaring struct type
	type Product struct {
		price       float64
		productName string
	}
	
	// declaring a function that takes in a struct value and modifies it
	func changeProduct(p Product) {
		p.price = 300
		p.productName = "Bicycle"
		// the changes are not seen to the outside world
	}
	
	// declaring a function that takes in a pointer to struct value and modifies the value
	func changeProductByPointer(p *Product) {
		(*p).price = 300
		p.productName = "Bicycle"
		// the changes are seen to the outside world
	
	}
	
	// declaring a function that takes in a slice
	func changeSlice(s []int) {
		for i := range s {
			s[i]++
		}
		// the changes are seen to the outside world
	}
	
	// declaring a function that takes in a map
	func changeMap(m map[string]int) {
		m["a"] = 10
		m["b"] = 20
		m["x"] = 30
		// the changes are seen to the outside world
	}

	// declaring some variables
    quantity, price, name, sold := 5, 300.2, "Laptop", true
    fmt.Println("BEFORE calling changeValues():", quantity, price, name, sold)
    // => BEFORE calling changeValues(): 5 300.2 Laptop true
 
    // invoking the function has no effect on the variables.
    // the function works on and modifies copies, not originals.
    changeValues(quantity, price, name, sold)
    fmt.Println("AFTER calling changeValues():", quantity, price, name, sold)
    // => AFTER calling changeValues(): 5 300.2 Laptop true
 
    // the function modifies the values.
    changeValuesByPointer(&quantity, &price, &name, &sold)
    fmt.Println("AFTER calling changeValuesByPointer():", quantity, price, name, sold)
    // => AFTER calling changeValuesByPointer(): 3 500.5 Mobile Phone false
 
    // declaring a struct value
    present := Product{
        price:       100,
        productName: "Watch",
    }
 
    // invoking the function has no effect on the struct value.
    // the function works on and modifies a copy, not the original.
    changeProduct(present)
    fmt.Println(present) // => {100 Watch}
 
    // the function modifies the struct value.
    changeProductByPointer(&present)
    fmt.Println("AFTER calling changeProductByPointer:", present)
    // => AFTER calling changeProductByPointer: {300 Bicycle}
 
    // declaring a slice
    prices := []int{10, 20, 30}
 
    // When a function changes a slice or a map the actual data is changed.
    changeSlice(prices)
    fmt.Println("prices slice after calling changeSlice():", prices)
    // => prices slice after calling changeSlice(): [11 21 31]
 
    // declaring a map
    myMap := map[string]int{"a": 1, "b": 2}
    // When a function changes a slice or a map the actual data is changed.
 
    changeMap(myMap)
    fmt.Println("myMap after calling changeMap():", myMap)
    // => myMap after calling changeMap(): map[a:10 b:20 x:30
 
    // slices and maps are not meant to be used with pointers.

}
