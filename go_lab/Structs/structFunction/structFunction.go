package main

import "fmt"

// struct method/function example

type authenticationInfo struct {
	username string
	password string
}

//this method below calls the struct above
func (auth authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("AUTHORIZATION: BASIC %s:%s", auth.username, auth.password)
}

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}