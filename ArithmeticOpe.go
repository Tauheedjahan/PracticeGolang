package main
import "fmt"

func main() {
  x:= 60
  y:= 40
  
  x++
  y--
  
  fmt.Println("addition=",x+y)
  fmt.Println("substraction=",x-y)
  fmt.Println("Multiplication=",x*y)
  fmt.Println("divison=",x/y)
  fmt.Println("Modulus=",x%y)
  fmt.Println("increment=",x)
  fmt.Println("decrement=",y)

}