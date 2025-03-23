package main 
import (
	"fmt"
	"time"
)

func linear_scan(array []int, target int ) int {
	i:= 0
	for i < len(array) {
		if array[i] == target {
			return i
		}
		i++
	}
	return -1
}


func binary_search(srt_arr []int, target int) int {
	i_min := 0
	i_max := len(srt_arr) - 1

	for i_min <= i_max	{
		i_mid := (i_max + i_min) / 2  // It is an integer because 2 integer division returns int.
		
		if srt_arr[i_mid] == target {
			return i_mid
		} else if srt_arr[i_mid] > target {
			i_max = i_mid - 1 
		} else {
			i_min = i_mid + 1
		}
	}
	return -1
}


func main () {
	array := make([]int, 10000001)
	for i := 0; i <= 10000000; i++ {
        array[i] = i
    }
	target := 100005

	start := time.Now()
	fmt.Printf("Result linear_scan O(n): %d\n",linear_scan(array, target))
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %d\n", elapsed)
	start = time.Now()
	fmt.Printf("Result binary_search O(logn): %d,\n", binary_search(array, target))
	elapsed = time.Since(start)
	fmt.Printf("Elapsed: %d\n", elapsed)
}