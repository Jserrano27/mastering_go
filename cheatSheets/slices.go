/////////////////////////////////
// Slices in Go
// Go Playground: https://play.golang.org/p/e5Ijl_7BgTr
/////////////////////////////////

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// declaring a string slice, by default is initialized with nil or uninitialized
	var cities []string

	fmt.Println("cities is equal to nil: ", cities == nil) // -> cities is equal to nil:  true
	fmt.Printf("cities: %#v\n", cities)                    // -> cities: []string(nil)

	// we can not assign elements to nil slice:
	// cities[0] = "Paris" // -> runtime error

	// declaring a slice using a slice literal
	numbers := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers)         // => [2 3 4 5]

	// creating a slice using the make() built-in function
	// creating a slice with 2 int elements initialized with zero.
	nums := make([]int, 2) // the same as []int{0, 0}.
	fmt.Println(nums)      // => [0 0]

	// declaring a slice using a slice literal
	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	// getting an element from the slice
	x := numbers[0]
	fmt.Println("x is", x) // => x is 2

	// modifying an element of the slice
	numbers[1] = 200
	fmt.Printf("%#v\n", numbers) // => []int{2, 200, 4, 5}

	// iterating over a slice
	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	//iterating over a slice (C/C++, Java Style)
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers[i])

	}

	// slices with the same element type can be assigned to each other
	var n []int
	n = numbers
	_ = n

	//** COMPARING SLICES **//
	// a Go slice can only be compared to nil

	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	// we can not compare slices using the equal (=) operator
	// fmt.Println(nn == mm) //error -> slice can only be compared to nil

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	//********************//
	// SLICE EXPRESSIONS //
	//******************//

	// arrays, slices and strings are sliceable
	// slicing doesn't modify the array or the slice, it returns a new one

	// declaring an [5]int array
	a := [5]int{1, 2, 3, 4, 5}

	// a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
	// this selects a range of elements which includes the element at index start, but excludes the element at index stop.

	// slicing an array returns a slice, not an array
	b := a[1:3]                                 // 1 is called start (included), 3 is called stop (excluded)
	fmt.Printf("Type: %T , Value: %#v\n", b, b) // => Type: []int , Value: []int{2, 3}

	// declaring a slice
	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]   //start included, stop excluded
	fmt.Println(s2) //[2 3]

	//for convenience, any of the indexis may be omitted.
	// a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
	fmt.Println(s1[2:])       // => [3 4 5 6], same as s1[2:len(s1)]
	fmt.Println(s1[:3])       // => [1 2 3], same as s1[0:3]
	fmt.Println(s1[:])        // => [1 2 3 4 5 6], same with s1[0:len(s1)]
	fmt.Println(s1[:len(s1)]) // => => [1 2 3 4 5 6], returns the entire slice
	// fmt.Println(s1[:45])   //panic: runtime error: slice bounds out of range

	s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
	fmt.Println(s1)          // -> [1 2 3 4 100]

	s1 = append(s1[:4], 200) // overwrites the last element
	fmt.Println(s1)          // -> [1 2 3 4 200]

	/////////////////////////////////
	// Slice's Backing Array
	// Go Playground: https://play.golang.org/p/UKc96Oq20IN
	/////////////////////////////////

	// Go implements a slice as data structure called Slice Header.
	// Slice Header contains 3 fields:
	// - the address of the backing array (pointer).
	// - the length of the slice.  The built-in function len() returns it.
	// - the capacity of the slice. The size of the backing array after the slice first element. cap() built-in function returns it.

	// A nil slice doesn't have backing array, so all the fields in the slice header are equal to zero.

	// a slice expression doesn't create a new backing array. The original and the returned slice are connected!
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] //s3, s4 share the same backing array with s1

	s3[1] = 600     // modifying the backing array so s1, s3 and s4 are in fact modified!!
	fmt.Println(s1) // -> [10 600 30 40 50]
	fmt.Println(s4) // -> [600 30]

	// when a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2                       // modifying the array
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                              // only cars is modified
	fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda]

	// Slices weights less than arrays:
	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("Array's size in bytes: %d \n", unsafe.Sizeof(a)) // 40 BYTES
	fmt.Printf("Slice's size in bytes: %d \n", unsafe.Sizeof(s)) // 24 BYTES
}
