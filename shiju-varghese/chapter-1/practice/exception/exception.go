package exception

import "fmt"


func DoPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recover with: ", e)
		}
	}()
	panic("Just paniking for the sake of demo ")
	fmt.Println("This will never be called.")
}

func CheckDeferSe3quence() {
	defer func() {
		fmt.Println("First Defer Function")
	}()

	defer func() {
		fmt.Println("Second Defer Function")
	}()

	defer func() {
		fmt.Println("Third Defer Function")
	}()
	

}