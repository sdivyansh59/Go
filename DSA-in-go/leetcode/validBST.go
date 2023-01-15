package main

import "fmt"

func main() {
	// fmt.Println("fdf")
    // A := [] int {7,7,10,10,9,5,2,8}
	//  A := [] int {1,1,2,3,4}
	A := [] int {5,3,1,2,8,7}
	fmt.Println(solve(A))

}


func solve(arr []int )  (int) {
	start , end := 0, len(arr)-1
	ans := validBST(arr, start, end)
    fmt.Println(ans)
    if ans {
		return 1
	}else {
		return 0;
	}
}

func validBST(arr [] int , start int, end int ) bool{
	if( start >= end) {
		return true
	}
	rootVal := arr[start]
	rightStart := end+1
	for i := 0; i <= end ; i++ {
		if arr[i] > rootVal {
			rightStart = i
			break
		}
	} 


	// left sub-tree ele 
	if check ( rootVal, arr, start+1, rightStart-1, "smaller")  == false {
		return false
	}else {
		fmt.Println("check passed ", (start+1) , rightStart-1)
	}
	// right sub-tree ele
	if check ( rootVal,arr, rightStart, end, "greater") == false {
		return false
	}else {
		fmt.Println("check passed for", rightStart,  end)
	}

	// left call
    if validBST(arr, start+1, rightStart-1) == false{
		return false
	}
	// right call
	if validBST(arr, rightStart, end)== false {
		return false
	}
	return true
}


func check ( key int ,arr[]int , start int , end int , condition string ) bool {
	for  i:= start ; i<= end ; i++{
		if condition == "greater"  && arr[i] <= key{
			return false
		}
		if condition == "smaller" && arr[i] >= key  {
			return false
		}
	}
	return true
}