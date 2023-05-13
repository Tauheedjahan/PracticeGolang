package main
import "fmt"

func main() {
  arr :=[]int{1,2,3,5,6,7}
  myslice := arr[2:5]  //change the length of slice
  
  mslice := make([]int, 5, 10)
  
  moslice :=[]int{10,20,30} 
  moslice[2] = 50           
  
  appslice :=append(arr,moslice...)
  
  newnumber := arr[:len(arr)-2]
  copynumber := make([]int,len(newnumber))
  copy(copynumber,newnumber)
  
	
	fmt.Printf("array slice value = %v \n", arr)
	fmt.Printf("array slice value = %v \n", myslice)
	fmt.Printf("length change value = %d \n", len(myslice)) //length change
	fmt.Printf("array length = %d \n", len(arr))
	fmt.Printf("array capacity = %d \n", cap(arr))
	
	fmt.Printf("array list = %v\n", mslice)
	fmt.Printf("array length = %d \n", len(mslice))
	fmt.Printf("array capacity = %d \n", cap(mslice))
	
	fmt.Println("array number change:=",moslice[2])  //change number the 2 index position
	 
	fmt.Printf("myslice3=%v\n", appslice)
	
	fmt.Printf("copynummber array = %v\n", copynumber)
  fmt.Printf("copynummber array = %v\n", len(copynumber))
  fmt.Printf("copynummber array = %v\n", cap(copynumber))
	
	
	
	
}