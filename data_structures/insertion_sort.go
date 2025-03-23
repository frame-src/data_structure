package main 
import "fmt"

func insertionSort(array []int ) {
	var i int = 1
	var ln = len(array)

	for  i < ln {
		current := array[i]
		j := i - 1
		for j >= 0 && array[j] > current {
			array[j + 1] = array[j]
			j = j - 1
		}
		array[j + 1] = current
		i = i + 1
	}
}

func main() {
    A := []int{5, 2, 9, 1, 5, 6}
    insertionSort(A)
    fmt.Println("Sorted array:", A)
}

