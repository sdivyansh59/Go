package swapchar

import "fmt"

func SwapCharactersInSentence(input string) string {

	output := ""
	for _, v := range input {
		// fmt.Printf("%d  %c \n",i, v)
		if v >= 65 && v < 65+26 {
			output += string(v + 32)
		} else if v >= 97 && v < 97+26 {
			output += string(v - 32)
		} else {
			output += string(v)
		}

	}
	fmt.Printf("input : %s  and output : %s \n", input, output)
	return output
}