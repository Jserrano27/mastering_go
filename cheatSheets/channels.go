// ///////////////////////////////
// Intro to Channels
// Go Playground: https://play.golang.org/p/Uc7iiqVeZLL
// ///////////////////////////////

package main

import (
	"fmt"
	"strings"
)

// declaring a function that computes the factorial of n
func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	// sending the factorial value into the channel
	c <- f
}

func main() {
	// Declaring a channel of type `chan int`
	var c1 chan int
	fmt.Println(c1) // => nil (its zero value is nil)

	// Initializing the channel
	c1 = make(chan int)
	fmt.Println(c1) // => 0xc000078060 (the channel stores an address)

	// Declaring and initializing a channel at the same time
	c2 := make(chan int)
	_ = c2

	// Declaring and initilizing a RECEIVE-ONLY channel
	c3 := make(<-chan string)

	// Declaring and initilizing a SEND-ONLY channel
	c4 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", c1, c3, c4) // => chan int, <-chan string, chan<- string

	// ** The "arrow" indicates the direction of data flow!! **//

	// Sending a value into the channel
	c1 <- 10

	// Receiving a value from the channel
	num := <-c1
	_ = num

	// Waiting for a value to be sent into the channel and print out that value
	fmt.Println(<-c1)

	// Closing a channel
	close(c1)

	// declaring and initializing a channel of type `chan int`
	ch := make(chan int)

	// defer closing the channel
	defer close(ch)

	// launching a goroutine
	go factorial(5, ch)

	// main() is waiting for a value to come from the channel
	// this is called a `blocking call`

	f := <-ch // receiving the value from the channel in a new variable
	fmt.Println("5 factorial =", f)

	// Spawning 20 goroutines that calculate the factorial
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Printf("Factorial of %d: %d\n", i, f)
	}

	fmt.Println(strings.Repeat("#", 10))

	// Spawning another 10 goroutines this time as anonymous functions
	for i := 5; i < 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}

			// sending the value f into the channel
			c <- f
		}(i, ch)
		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}

}
