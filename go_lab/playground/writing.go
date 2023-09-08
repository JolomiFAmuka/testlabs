package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Kindly Provide a brief Description of GolinuxCloud company:  once done typing, kindly tap enter button to display your content.")
	// using standard input in this case its the terminal
	// once done
	userData, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		_ = fmt.Errorf(" %v error", err)
		return
	}
	// Display user's data stdout
	results := strings.TrimSuffix(userData, "\n")
	fmt.Printf("Results you Provided: \n %v", results)

}
