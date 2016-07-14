package main

import "fmt"

func main()  {
	var mySlice []float64 // Here is an example of a slice
	if mySlice == nil {
		fmt.Println("The slice is empty")
	}
	otherSlice := []int{1,2,3}
	if otherSlice == nil {
		fmt.Println("Is empty")
	}else {
		fmt.Println("The length is:", len(otherSlice))
	}
	//
	fmt.Println("=============================================")
	//
	arreglo := [5]int{1,2,3,4,5} // This is a Array
	aSlice  := arreglo[:]        // Converting the Array to Slice
	/***********************************************************************************
	We can use [low : high] expression, for example:
	arreglo[0:] is the same as arreglo[0:len(arreglo)]
	arreglo[:5] is the same as arreglo[0:5]
	arreglo[:]  is the same as arreglo[0:len(arreglo)]
	************************************************************************************/
	fmt.Println(aSlice)
	//
	fmt.Println("=============================================")
	//
	mSlice  := make([]int, 5, 10)
	mSlice   = append(mSlice, 3)

	fmt.Println(mSlice)
	fmt.Println("The length is:", len(mSlice))
	fmt.Println("The capacity is:", cap(mSlice))

	//
	fmt.Println("=============================================")
	//

	zSlice      :=      []int{1,2,3,4,5,6,7}
	copySlice   :=      make([]int, len(zSlice))

	copy(copySlice,zSlice)

	fmt.Println(zSlice)
	fmt.Println(copySlice)
}