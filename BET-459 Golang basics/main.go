package main

import (
	"fmt"
)

// checkIfInt function is used to check if the input is an integer
func checkIfInt() int {
	var x int
	fmt.Println("Enter a new value :")
    for true{
		
		_, err := fmt.Scan(&x)
		if err == nil {
			break
		}
		fmt.Println("Enter a valid int: ",err)
	}
    return x

}
// getMin function is used to get the minimum value of a two given variables as arguments
func getMin(a,b int) int {
	var min int
	if a>b {
		min=b
	}else{
		min=a
	}
	return min
}

// gcf function is used to calculate the greatest common factor of a two given variables as arguments
func gcf(x,y int) int {
	var result int
	min:=getMin(x,y)
	for i := 1; i < min + 1; i++ {
		if (x%i) == 0 && (y%i) == 0{ 
			result = i
		}
	}
    return result
}
func main() {
    var x,y int
	
    
    x = checkIfInt()
	
	y = checkIfInt()
    
	
    fmt.Printf("the greatest common factor of %d and %d is %d\n",y,x,gcf(x,y))
}