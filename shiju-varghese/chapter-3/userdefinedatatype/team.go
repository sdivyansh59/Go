package userdefinedatatype

import "fmt"

type Team struct {
	Name, Desription string
	User             []User
}

func (t Team) GetTeamDetils() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Desription)
	fmt.Println("Details of team members:")
	for _,v := range t.User {
		v.PrintName()
		v.PrintDetails()	
	}
}