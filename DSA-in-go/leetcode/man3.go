package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("hello")
	arr := []int{4, 42, 40, 26, 46}
	fmt.Println(solve(arr, 20))

}

func solve(arr []int, target int) int {
	n := len(arr)
	suffixSum := make([]int, n)
	suffixSum[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + arr[i]
	}

	return find(arr, suffixSum, target)
}

func find(arr []int, suffixSum []int, targetWood int) int {
	n := len(arr)
	min, max := arr[0], arr[n-1]
    wood := math.MaxInt
	heigh := 0

	for min <= max {
		mid := min + (max-min)/2

		idx := lowerMaximum(arr , mid)
		remainWood := mid * (n-idx)
		outputWood := suffixSum[idx] - remainWood

		if outputWood >= targetWood {
			if outputWood < wood {
				wood = outputWood
				heigh = mid
			}
			min = mid+1
		}else {
			max = min -1
		}
	}
	return heigh
}


func lowerMaximum(arr [] int , target int) int {
	i, j := 0 , len(arr)-1
	currMax := math.MaxInt
    idx := 0

	for i<= j {
		mid := i + (j-i)/2

		if(target== arr[mid]){
			return mid
		}else if (target < arr[mid]) {
			if currMax > arr[mid] {
				currMax = arr[mid]
				idx = mid
			}
			
			j = mid -1
		}else {
			i = mid +1
		}

	}

	return idx
}