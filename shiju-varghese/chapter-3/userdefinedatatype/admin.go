package userdefinedatatype

import "fmt"

type Admin struct {
	Person
	Roles []string
}

// over ride print method

func (a Admin) PrintDetails() {
	// call person printdetails
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}