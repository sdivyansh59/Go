package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var nums = [] int{0,0,1,1,1,2,2,3,3,4}
	ans := removeDuplicates(nums)
	fmt.Println(ans)
}


func removeDuplicates(nums []int) int {
	size := len(nums)
    if size ==0 {
		return 0;
	}

    i := 0
    j := 1
	for  j<size {
		if(nums[i] != nums[j]){
			i++
			nums[i] = nums[j]	
		}
		j++
	}
	return i+1  
}
