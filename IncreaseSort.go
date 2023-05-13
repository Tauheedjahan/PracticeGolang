package main

import (
    "fmt"
     "sort"
   
)

func main (){
		
        i:=[]int{1,2,3,33,8,100,12,55,45}
        sort.Ints(i)
        fmt.Println("increment ", i)
        
         s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} 
	      sort.Float64s(s)
	      fmt.Println("increament ",s)
	      
	      
	      strs := []string{"c", "a", "b" , "z"}
        sort.Strings(strs)
        fmt.Println("increment", strs)
        
   }
	