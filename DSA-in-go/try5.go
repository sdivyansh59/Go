package main

import "fmt"

func seats(A string) int {
	var count int
	c1, c2 := 0, 0

	i, j := 0, len(A)-1

	for A[i] == '.' {
		i++
	}
	for A[j] == '.' {
		j--
	}

	fmt.Println(i,j)

	for i < j {
		// handle i
		if c1 <c2 {
			if A[i] == '.' {
			count += c1 
		} else {
			c1++
		}
		i++


		}else{
			// handle j
		if A[j] == '.' {
			count += c2
		} else {
			c2++

		}

		j--
		}
		

		

		// if c1 <c2 {
		// 	i++
		// }else{
		// 	j--
		// }
		// i++
		// j--

	}

	fmt.Println(count)

	return count

}
