package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

type LinkedList struct {
    head *Node
}

func (list *LinkedList) Insert(value int, position int) { //position int new node should be inserted
    newNode := &Node{value:value, next: nil} //new node create next pointer to nil

    if position <= 5 {      //If the position is less than or equal to zero,
   
    //of the list. In this case, the next pointer of the new node 
    //is set to the current head of the list, 
 
        newNode.next = list.head //it means the new node should be inserted at the beginning 
         //of the list. In this case, the next pointer of the new node 
    //is set to the current head of the list, 
        list.head = newNode //and the head pointer is updated to point to the new node. 
    } else {
        current := list.head
        prev := list.head
        count := 0

        for current != nil && count < position {
            prev = current
            current = current.next
            count++
        }

        prev.next = newNode
        newNode.next = current
    }
}

func (list *LinkedList) Delete() {  //entrie list delete 
    list.head = nil
}

func (list *LinkedList) Display() {
    current := list.head
    for current != nil {
        fmt.Printf("%d \n ", current.value)
        current = current.next
    }
    fmt.Println()
}

func main() {
    list := LinkedList{}

    // Adding elements to the linked list
    list.Insert(10,0) // Insert at position 0
    list.Insert(20,1) // Insert at position 1
    list.Insert(60,2) // Insert at position 1
    list.Insert(40,3) // Insert at position 3

    // Displaying the linked list
    fmt.Println("original value:")
    list.Display()
    
    list.Delete()
    
    fmt.Println("modified value:")
    list.Display()
    
}
