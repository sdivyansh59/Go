package main

import (
	"chapterthree/userdefinedatatype"
	// "fmt"
	"time"
)


func main() {
	
	// p := &userdefinedatatype.Person {
	// 	FirstName : "Divyansh ",
	// 	LastName : "Singh",
	// 	Dob : time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
	// 	Email : "sdivyansh59@gmail.com",
	// 	Location : "Gorakhpur",
	// }

	// p.PrintName()
	// p.PrintDetails()
	// fmt.Println("****************")
	// p.ChangeLocation(                     "mau")
	// p.PrintDetails()

	alex := userdefinedatatype.Admin{
		Person: userdefinedatatype.Person{
			FirstName: "Alex",
			LastName: "John",
			Dob: time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			Email: "|alex@emil.com",
			Location: "New York",
		},
		Roles: []string{"Manage Team", "Manage Tasks"},
	}

	
	shiju := userdefinedatatype.Member {
		Person: userdefinedatatype.Person {
			FirstName: "Shiju",
			LastName: "Varghese",
			Dob: time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			Email: "shiju|@gmail.com",
			Location: "kochiu",
		},
		Skills : []string{"Go", "Docker", "Kubernates"},
	}

	chris := userdefinedatatype.Member {
		Person: userdefinedatatype.Person {
			FirstName: "Chris",
			LastName: "Martin",
			Dob: time.Date(1989, time.February, 27, 0, 0, 0, 0, time.UTC),
			Email: "martin@gmail.com",
			Location: "benglore",
		},
		Skills : []string{"Go", "Docker", "GCP","AWS"},
	}

	team := userdefinedatatype.Team{
		Name: "Go Developers",
		Desription: "Golang COE",
		User: []userdefinedatatype.User {alex, chris, shiju},
	}

	team.GetTeamDetils()
	

	// call method ofr alex
	// alex.PrintName()
	// alex.PrintDetails()

	// // call method for shiju
	// shiju.PrintName()
	// shiju.PrintDetails()

	// users := []userdefinedatatype.User{alex, shiju}

	// for _, v := range users{
	// 	v.PrintName()
	// 	v.PrintDetails()
	// }


}

