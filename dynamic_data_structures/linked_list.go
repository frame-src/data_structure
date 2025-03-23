package main 
import "fmt"

type Node struct {
	value int
	next *Node
}

func (ll *Node) Print(){
	for ll != nil {
		fmt.Println(ll.value)
		ll = ll.next
	}
}

func (ll *Node) Len() int {
	i := 0
	for ll != nil {
		ll = ll.next
		i++
	}
	return i
}



func main () {
	ll := &Node{value:0}
	head := ll

	for j := 1; j<10; j++ {
		ll.next = &Node{value:j}
		ll = ll.next
	}

	fmt.Printf("Last Node:\n")
	ll.Print()
	fmt.Printf("Head:\n")
	head.Print()
	head.Print()
	fmt.Printf("List lenght: %d \n", head.Len())
}