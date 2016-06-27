package main

import "fmt"

func main()  {
  // var x [25]int
  // x[10] = 100
  // fmt.Println(x)

// Leccion de los slide
  slide1 := []int{1,2,3}
  slide2 := append(slide1, 4, 5)
  fmt.Println(slide2)
}
