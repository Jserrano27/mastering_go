/////////////////////////////////
// String is Go
// Go Playground: https://play.golang.org/p/-o07MQbIsDv
/////////////////////////////////

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// Strings are defined between double quotes "..."
	// Strings in Go are UTF-8 encoded by default
	// A string is in fact a slice of bytes in Go

	// declaring a string
	s1 := "Hi there  Go!"

	// printing a string
	fmt.Printf("%s\n", s1) // => Hi there  Go!
	fmt.Printf("%q\n", s1) // => "Hi there  Go!"

	// using double-quotes inside double quotes
	fmt.Println("He say: \"Hello!\"")

	// double quotes inside backticks (backquote)
	fmt.Println(`He say: "Hello!"`)

	// a string literal inclosed in backticks is called a raw string and it is interpreted literally.
	// backslashes or \n  have no special meaning
	s2 := `Hi there Go!`
	fmt.Println(s2)

	// declaring a multiline string
	fmt.Println("Price: 100 \nBrand: Nike")

	//the same with:
	fmt.Println(`
Price: 100
Brand: Nike`)

	// using backslashes inside a string:
	fmt.Println(`C:\Users\Andrei`)
	fmt.Println("C:\\Users\\Andrei")

	// concatenating strings (+)
	// Go creates a new string because strings are immutable in Go (this is not efficient).
	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!") // -> I love Go Programming!

	// getting an element (byte) of a string:
	fmt.Println("Element at index zero:", s3[0]) // => 73 (ascii code for I)
	//  a string is in fact a slice of bytes in Go

	// strings are immutable and can't be changed
	// s3[5] = 'x' // => error: Cannon assign to s3[5].

	/////////////////////////////////
	// Strings, Runes, Bytes and Unicode Code Points
	// Go Playground: https://play.golang.org/p/pttCqLAAvKA
	/////////////////////////////////

	// characters or rune literals are expressed in Go by enclosing them in single quotes
	// declaring a variable of type rune (alias to int32)
	var1 := 'a'
	fmt.Printf("Type: %T, Value:%d \n", var1, var1) // => Type: int32, Value:97
	// value is 97 (the code point for a)

	// declaring a string value that contains non-ascii characters
	str := "ţară" // ţară means country in Romanian
	// 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println(len(str)) // -> 6,  4 runes in the string but the length is 6

	// returning the number of runes in the string
	m := utf8.RuneCountInString(str)
	fmt.Println(m) // => 4

	// by using indexes we get the byte at that positioin, not rune.
	fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163

	// decoding a string byte by byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // -> Å£arÄ
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune manually:
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:]) // it returns the rune in string in variable r
		//and its size in bytes in variable size

		// printing out each rune
		fmt.Printf("%c", r) // -> ţară
		i += size           // incrementing i by the size of the rune in bytes
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune automatically:
	for i, r := range str { //the first value returned by range is the index of the byte in string where rune starts
		fmt.Printf("%d -> %c", i, r) // => ţară
	}

	 // Slicing a string is efficient because it reuses the same backing array
    // Slicing returns bytes not runes
 
    s1 := "abcdefghijkl"
    fmt.Println(s1[2:5]) // -> cde, bytes from 2 (included) to 5 (excluded)
 
    s2 := "中文维基是世界上"
    fmt.Println(s2[0:2]) // -> � - the unicode representation of bytes from index 0 and 1.
 
    // returning a slice of runes
    // 1st step: converting string to rune slice
    rs := []rune(s2)
    fmt.Printf("%T\n", rs) // => []int32
 
    // 2st step: slicing the rune slice
    fmt.Println(string(rs[0:3])) // => 中文维
}
}