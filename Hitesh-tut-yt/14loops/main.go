package main

import "fmt"

func main() {
	fmt.Println("loops and goto/ continue statement")

	days := []string{"Monday", "Tuesday", "Wednesday", "Sunday"}

	// Type -> 1 , i is index 
	for i := range days {
		fmt.Println(days[i])
	}

	// Type -> 2 , 
	for idx, val := range days {
		fmt.Printf("index %v  val is %v \n", idx , val)
	}
	// if u want to escape something use _

	// Type -> 3 , not use much 
	for day := 0 ; day < len(days) ; day++ {
		fmt.Println(days[day])
	}

	rougueVal := 0

	for rougueVal <7 {
		rougueVal++
		if  rougueVal == 4 {
			continue
		}
		if rougueVal == 6 {
			goto lco
		}

		fmt.Println(rougueVal);
		//++rougueVal Doubt why this is not working ?
		

		// if  rougueVal == 3 {
		// 	break;
		// }
		
	}


	// gotto
	lco:
		fmt.Println("Jumpng at learncodeonline.in")
}