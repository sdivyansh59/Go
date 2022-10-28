package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang ; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 37}

	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are %+v \n", hitesh)

	// methods concepts
	hitesh.GetStatus()
	hitesh.NewMail()

	// Ques -> now print hitest email and see it is changed or not?
	fmt.Println(hitesh.Email)
	// Ans -> not changed, becz it is pass by value ( not by ref)



}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int	
}
// Note: oneAge is not  exporable becz it is not Capital letter


func (u User) GetStatus (){
	fmt.Println("Is  user active ", u.Status)
}

func ( u User) NewMail (){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

/* Note

1 . Function : normal func
2. method : func which is use for Struct ( syntax is also little bit changed)
3. pass by val vs pass by ref ( using pointer)
*/