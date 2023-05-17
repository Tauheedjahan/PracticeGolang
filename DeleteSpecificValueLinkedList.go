package main


import "fmt"

type Node struct{
      
      value int
      next *Node
}
type LinkedList struct{
      head *Node
}

func (list *LinkedList) Insert(value int){
  newNode := &Node{value:value}
  
  if list.head ==nil{
    list.head=newNode
  } else {
    
    current :=list.head
    for current.next !=nil{
      current=current.next
    }
    current.next=newNode
  }
}

func (list*LinkedList) Delete (value int){
    if list.head ==nil{
      return
    }
    if list.head.value == value{
      list.head=list.head.next
      return
    }
    current:=list.head
    prev:=list.head
    
    
    for current!=nil && current.value!= value{
      prev = current
      current=current.next
      
    }
    
    if current==nil{
      return
    }
    
    prev.next=current.next
}

func (list*LinkedList) Display(){
      current:=list.head
      for current!=nil{
        fmt.Printf("%d\n",current.value)
        current=current.next
      }

    fmt.Println()
}
 func main(){
   list:= LinkedList{}
   
   list.Insert(10)
   list.Insert(20)
   list.Insert(30)
   list.Insert(80)
   
   fmt.Println ("original linked list:")
   list.Display()
   
   list.Delete(20)
   list.Delete(80)
   
   fmt.Println("modified value:")
   list.Display()
}









































































