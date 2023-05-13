package main

import (
    "fmt"
     "sort"
   
)

func main (){
		
        i:=[]int{1,2,3,33,8,100,12,55,45}
        j:=sort.IntsAreSorted(i) 
        fmt.Println("its a sorted value", j)
        
         k:=[]int{1,2,3,4,5,6,7,8,9}
        fmt.Println("It's sorted",sort.IntsAreSorted(k))
   }
	