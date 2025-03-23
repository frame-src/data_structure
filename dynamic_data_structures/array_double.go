package main

import "fmt"


func arrayDouble[T any](old []T) []T {
	lenght := len(old)
	new := make([]T, lenght*2)
	for j:= 0; j<lenght; j++ {
		new[j] = old[j]
	}
	return new
}


func main() {
    arr := make([]int, 20, 20) // type, len, capacity
	for j:= 0; j < len(arr); j++ {
		arr[j] = j
	}
	fmt.Println(arr)
	arr = append(arr,1,2,3,4)
	fmt.Printf("Len old array: %d\n", len(arr))
	fmt.Println("Capacity:", cap(arr))
	new := arrayDouble(arr)
	fmt.Println(new)
	fmt.Printf("Len new array = %d\n", len(new))
	fmt.Println("Capacity:", cap(arr))
}