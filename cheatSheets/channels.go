// ///////////////////////////////
// Intro to Channels
// Go Playground: https://play.golang.org/p/Uc7iiqVeZLL
// ///////////////////////////////

package main

import (
	"fmt"
	"strings"
	"time"
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

	/////////////////////////////////
	// Unbuffered Channels
	// Go Playground: https://play.golang.org/p/_44csjQDJvM
	/////////////////////////////////

	c1 := make(chan int) //unbuffered channel

	// Launching a goroutine
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1) // calling the anonymous func and passing c1 as argument

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	// we sleep for a second to give time to the goroutine to finish
	time.Sleep(time.Second)

	// After running the program we notice that the sender (the func goroutine) blocks on the channel
	// until the receiver (the main goroutine) receives the data from the channel.
	//** EXPECTED OUTPUT: **//
	// main goroutine sleeps for 2 seconds
	// func goroutine starts sending data into the channel
	// main goroutine starts receiving data
	// main goroutine received data: 10
	// func goroutine after sending data into the channel

	/////////////////////////////////
	// Buffered Channels
	// Go Playground: https://play.golang.org/p/1wwkXh4dcs3
	/////////////////////////////////

	// Declaring a buffered channel.
	c1 := make(chan int, 3)

	fmt.Println("Channel's capacity:", cap(c1)) // => 3

	// spawning a new goroutine
	go func(c chan int) {
		// sending 5 values into the channel
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		// closing the buffered channel.
		close(c)

	}(c1) //calling the anonymous func and passing c1 as argument

	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	// receiving data from the channel
	for v := range c1 { // v is the value read from the channel, it's like using v := <- c2
		fmt.Println("main goroutine received value from channel:", v)

	}

	// After running the program  we notice that the goroutines start sending data
	// into the channel BEFORE the main goroutine had a chance
	// to receive data from the channel.

	// The sender of this buffered channel will block only when there is no empty slot in the channel, in this
	// case after 3 writing attempts because the channel has a capacity of 3.
	// The receiver will block on the channel when it's empty.

	// A receive operation on a closed channel will proceed without blocking
	// and yield the zero-value for the type that is sent through the channel.
	fmt.Println(<-c1) // => 0

	// Sending a value into a closed channel will panic.
	// c1 <- 10 // => panic: send on closed channel

}
