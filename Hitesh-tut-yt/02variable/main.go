package main

import "fmt"

const LoginToken string = "dfdfdfdfd"; // public variable
const a int = 10;

//  a := 10 // not allowed in global space

func main() {
	// fmt.Printf("varibales")
	var username string = "hitesh";
	fmt.Println(username);
	fmt.Printf("variable of type %T \n", username);

	var isLoggedIn bool = true;
	fmt.Println(isLoggedIn);
	fmt.Printf("variable of type %T \n", isLoggedIn);

	var tempint int;
	var tempStr string;
	fmt.Println(tempint,tempStr);


	// without var
	numberOfUsers := 320004500;
	fmt.Println(numberOfUsers, LoginToken,a);


}