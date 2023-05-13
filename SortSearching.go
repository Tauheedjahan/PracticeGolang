package main

import (
    "fmt"
     "sort"
   
)

func main (){
		
        i:=[]int{1,2,3,6,8,12,33,45,100}
        j:=sort.Search(len(i), func(f int) bool{return i[f]>=100})
       	sort.Ints(i)
        fmt.Println("The value is 11 greater then", j)
        
   }
	