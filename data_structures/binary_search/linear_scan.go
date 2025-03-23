package main 
import (
	"fmt"
	"time"
)

func linearScan(array []int, target int ) int {
	i:= 0
	for i < len(array) {
		if array[i] == target {
			return i
		}
		i++
	}
	return -1
}


func binarySearch(srtArr []int, target int) int {
	iMin := 0
	iMax := len(srtArr) - 1

	for iMin <= iMax	{
		iMid := (iMax + iMin) / 2  // It is an integer because 2 integer division returns int.
		
		if srtArr[iMid] == target {
			return iMid
		} else if srtArr[iMid] > target {
			iMax = iMid - 1 
		} else {
			iMin = iMid + 1
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
	fmt.Printf("Result linear_scan O(n): %d\n",linearScan(array, target))
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %d\n", elapsed)
	start = time.Now()
	fmt.Printf("Result binary_search O(logn): %d,\n", binarySearch(array, target))
	elapsed = time.Since(start)
	fmt.Printf("Elapsed: %d\n", elapsed)
}