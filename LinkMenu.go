/*

Create a LinkedList
- Push element at the end of LinkedList
- Get elementby Position in the LinkedList
- Remove last item from the LinkedList
- edit element by position in LinkedList

*/

package main

import "fmt"

type Node struct {
	next  *Node
	value int
}

type List struct {
	head *Node
	tail *Node
}

func (ls *List) Insert(key int) {
	node := &Node{
		value: key,
	}

	if ls.head == nil {
		ls.head = node

	} else {
		ls.tail.next = node
	}
	ls.tail = node
}

func (ls *List) Print() {
	list := ls.head
	for list != nil {
		fmt.Printf("%+v -> ", list.value)
		list = list.next
	}
}

func (ls *List) RemoveLast() {

	list := ls.head
	if list == nil {
		ls.head = nil
		ls.tail = nil
	} else {
		if list.next == nil {
			ls.head = nil
			ls.tail = nil
		} else {
			for list != nil {

				fmt.Printf("%+v -> ", list.value)
				if list.next.next == nil {
					list.next = nil
					ls.tail = list
					break
				}

				list = list.next
			}
		}
	}
}

func (ls *List) GetElementByPos(pos int) {

	list := ls.head
	count := 0
	for list != nil {
		count++
		if count == pos {
			fmt.Printf("Key is %+v\n ", list.value)
			break
		}
		list = list.next
	}
}

func (ls *List) EditElementByPos(pos int, val int) {

	list := ls.head
	count := 0
	for list != nil {
		count++
		if count == pos {
			list.value = val
			break
		}
		list = list.next
	}
}

func main() {
	list := &List{}

here:
	fmt.Println("Project LinkedList\nSelect\n" +
	"1 - Insert element at the end of LinkedList\n" +
	"2 - Get elementby Position in the LinkedList\n" +
	"3 - Remove last item from the LinkedList\n" +
	"4 - Edit element by position in LinkedList\n" +
	"5 - Display LinkedList") 
	fmt.Println("Enter choice : ")

	var selectOption int
	fmt.Scanln(&selectOption)

	switch selectOption {
	case 1:
		fmt.Println("Enter key to insert : ")
		var value int
		fmt.Scanln(&value)
		list.Insert(value)
		fmt.Println("=====================")
		list.Print()
		fmt.Println("\n=====================")
		goto here
	case 2:
		var pos int
		fmt.Println("Enter position to get : ")
		fmt.Scanln(&pos)
		fmt.Println("=====================")
		list.GetElementByPos(pos)
		fmt.Println("\n=====================")
		goto here
	case 3:
		fmt.Println("=====================")
		list.RemoveLast()
		fmt.Println("\n=====================")
		goto here
	case 4:
		var pos int
		var val int
		fmt.Println("Enter position to Update : ")
		fmt.Scanln(&pos)
		fmt.Println("Enter value to Update : ")
		fmt.Scanln(&val)
		fmt.Println("=====================")
		list.EditElementByPos(pos, val)
		fmt.Println("\n=====================")
		goto here
	case 5:
		fmt.Println("=====================")
		list.Print()
		fmt.Println("\n=====================")
		goto here
	default:
		goto here
	}
}