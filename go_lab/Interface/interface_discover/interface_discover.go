package main

import (
	"fmt"
)

//1. define an interface
type usercreds interface {
	access() string
}

//2. define the struct (i.e. type to call the interface)
type userdetails struct {
	firstname, lastname string
}

//3. define a method for the struct (implement the interface method signature**) 
func (usercreation userdetails) access() string {
	return fmt.Sprintf("%s.%s@bellmts.ca", usercreation.firstname, usercreation.lastname)
}

//4. define a method for usercreds
func getcreds(creds usercreds) string {
	return creds.access()
}

func main() {
	person1 := userdetails{
		firstname: "jolomi",
		lastname:  "amuka",
	}

	person2 := userdetails{
		firstname: "oyin",
		lastname:  "akinboye",
	}

	person3 := userdetails{
		firstname: "scott",
		lastname:  "Gbegbosky",
	}

	fmt.Printf("Access details: %s\n", getcreds(person1))
	fmt.Printf("Access details: %s\n", getcreds(person2))
	fmt.Printf("Access details: %s\n", getcreds(person3))
}
