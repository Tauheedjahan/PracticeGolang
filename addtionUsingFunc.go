package main


import "fmt"


func addNumber(number1, number2 int) int {
   return number1 + number2
}
func main() {

  
   var number1, number2, number3 int
   
   
   number1 = 18
   number2 = 9
   
  
   number3 = addNumber(number1, number2)
   
  
   fmt.Println("The addition of ", number1, " and ", number2, " is \n", number3, "\n(adding two integers outside the function)")
}