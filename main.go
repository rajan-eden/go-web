package main

import "fmt"

const LoginToken string = "jdhkldhasfgkohdfsil"

func main() {
	var username string = "prabha"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type: %T \n", isLoggedin)

	var smallvall uint8 = 255
	fmt.Println(smallvall)
	fmt.Printf("variable is of type: %T \n", smallvall)

	var smallfloat float32 = 255.6743895634826582346586
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n", smallfloat)

	var bigfloat float64 = 255.6743895634826582346586
	fmt.Println(bigfloat)
	fmt.Printf("variable is of type: %T \n", bigfloat)

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T \n", anothervariable)

	// implicit type
	var website = "prabha.me"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)

	// no var style
	numberofuser := 300000.0
	fmt.Println(numberofuser)
	fmt.Printf("variable is of type: %T \n", numberofuser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
