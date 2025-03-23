package main
import "fmt"

func string_equal(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}
	var n int = len(first)
	var i int = 0
	for i < n && first[i] == second[i]{
		i++
	}
	return i == n
}

func main () {
	first := "HELLO world"
	second := "HELLO world"
	res := string_equal(first, second)
	fmt.Println("Result:",res)
	fmt.Println("audi" == "bmw")
}