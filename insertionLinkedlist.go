package main

import "fmt"

type Node struct {
    number int
    next  *Node
}

type LinkedList struct {
    head *Node
}

func (list *LinkedList) Insert(number int, pos int) { 
    newNode := &Node{number:number, next: nil} 

    if pos <= 5 {      
    newNode.next = list.head 
    list.head = newNode 
    } else {
        c := list.head
        p := list.head
        cou := 0

        for c != nil && cou < pos {
            p = c
            c = c.next
            cou++
        }

        p.next = newNode
        newNode.next = c
    }
}

func (list *LinkedList) Display() {
    c := list.head
    for c != nil {
        fmt.Printf("%d \n ", c.number)
        c = c.next
    }
    fmt.Println()
}

func main() {
    list := LinkedList{}

    // Adding elements to the linked list
    list.Insert(10, 0) 
    list.Insert(20, 1) 
    list.Insert(60, 2) 
    list.Insert(40, 3) 

    // Displaying the linked list
    list.Display()
}
