package linkedlist
import "fmt"

type Node[T any] struct {
	value T
	next *Node
	prev *Node
}

func (ll *Node[T]) Print(){

	for ll != nil {
		fmt.Println(ll.value)
		ll = ll.next
	}
}

func (ll *Node[T]) Len() int {
	i := 0
	
	for ll != nil {
		ll = ll.next
		i++
	}
	return i
}


func Append[T any](ll *Node[T], new *Node[T]) *Node[T] {
	head := ll
	
	for ll.next != nil {
		ll = ll.next
		i++
	}
	ll.next = new
	return head
}


func InsertAfter[T any](pr *Node[T], new *Node[T]) *Node {
	tmp := pr.next
	pr.next = new
	new.next = tmp
}


func LinkedListInsert[T any](head *Node[T], i int, val T) *Node[T]{
	new := &Node{value:val}
	if i == 0{
		new.next = head
		return new
	}
	current := head
	for j:=0; j!= i && current; j++{
		current = current.next
	}
	if !current {
		return nil
	}
	tmp = current.next
	current.next = new
	new.next = tmp
	return head
}


func LinkedListInsertDelete[T any] (head *Node[T], i int){
	if i == 0{
		head = head.next
		return head
	}
	prev := head 
	current:= head

	for j:=0; j!= i && current; j++{
		prev = current
		current = current.next
	}
	if !current { 
		return nil 
	}
	prev.next = current.next
	return head
}

// func test () {
// 	ll := &Node{value:0}
// 	head := ll

// 	for j := 1; j<10; j++ {
// 		ll.next = &Node{value:j}
// 		ll = ll.next
// 	}

// 	fmt.Printf("Last Node:\n")
// 	ll.Print()
// 	fmt.Printf("Head:\n")
// 	head.Print()
// 	head.Print()
// 	fmt.Printf("List lenght: %d \n", head.Len())
// }