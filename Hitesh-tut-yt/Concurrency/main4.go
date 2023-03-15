package main

func main() {

}

func candy(A []int) int {
	candy := make([]int, len(A), len(A))
	var sum int

	// dir -->
	candy[0] = 1
	for i := 1; i < len(A); i++ {
		candy[i] = 1
		if A[i] > A[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}

	// <-- dir
	sum += candy[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			if candy[i] <= candy[i+1] {
				candy[i] = candy[i+1]+1
			}
		}
		sum += candy[i]
	}

	return sum
}
