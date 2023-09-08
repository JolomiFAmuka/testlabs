package main

import "fmt"

func main() {

    //define the nested map. This one has a string value
	windowsteam := map[string]map[string]string{
		"Gene": map[string]string{
			"fullname": "Eugene Moreau",
			"Address":  "No 8 wilbur drive",
		},
		"Kevin": map[string]string{
			"fullname": "Eugene Moreau",
			"Address":  "No 19 st boniface",
		},
		"Ian": map[string]string{
			"fullname": "Ian Kaplonski",
			"Address":  "No 22 portage crescent",
		},
		"Rolando": map[string]string{
			"fullname": "Rolando Martinez",
			"Address":  "No 25 east river lane",
		},
	}
    //run a for loop to get full access of output
	if temp, staff := windowsteam["Ian"]; staff {
		fmt.Println(temp["fullname"], "\n", temp["Address"])
	}
}
