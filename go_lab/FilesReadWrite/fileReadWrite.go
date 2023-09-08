package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// create a file

	file, err := os.Create("sample.txt") //use the os package
    //creare methos returns error
    //ALWAYS catch the error
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hi, my name is Jolomi and this file was created using GOLANG")
	file.Close()

    //read the file with the ioutil package
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)
	fmt.Println(s1)
}
