/////////////////////////////////
// Creating, Opening, Closing, Renaming, Moving, and Removing files in Go
// Go Playground: https://play.golang.org/p/Sz_LfNS9GKU
/////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//** Use valid paths according to your OS. **//

	// CREATING A FILE

	// os.Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.
	// it returns a file descriptor which is a pointer to os.File and an error value.
	newFile, err := os.Create("a.txt")

	// error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors

	}

	// TRUNCATING A FILE
	err = os.Truncate("a.txt", 0) //0 means completely empty the file.

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// CLOSING THE FILE
	newFile.Close()

	// OPEN AND CLOSE AN EXISTING FILE
	file, err := os.Open("a.txt") // open in read-only mode

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//OPENING a FILE WITH MORE OPTIONS
	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
	// We can Use opening attributes individually or combined
	// using an OR between them
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// GETTING FILE INFO
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")

	p := fmt.Println
	p("File Name:", fileInfo.Name())        // => File Name: a.txt
	p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
	p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
	p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
	p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----

	// CHECKING IF FILE EXISTS
	fileInfo, err = os.Stat("b.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	// RENAMING AND MOVING A FILE
	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// REMOVING A FILE
	err = os.Remove("aa.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}

	/////////////////////////////////
	// Writing Bytes to Files
	// Go Playground: https://play.golang.org/p/Zc3KDG7kYvt
	/////////////////////////////////

	// opening the file in write-only mode if the file exists and then it truncates the file.
	// if the file doesn't exist it creates the file with 0644 permissions
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// WRITING BYTES TO FILE

	byteSlice := []byte("I learn Golang! 传")   // converting a string to a bytes slice
	bytesWritten, err := file.Write(byteSlice) // writing bytes to file.
	// It returns the no. of bytes written and an error value
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten) // => 2019/10/21 16:26:16 Bytes written: 19

	// WRITING BYTES TO FILE USING ioutil.WriteFile()

	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("Go Programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
