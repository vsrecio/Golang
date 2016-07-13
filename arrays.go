package main

import "fmt"

func main()  {
  // Array of int with a length of 10
  arreglo := [10]int{1,2,3,4}
  fmt.Println("the length of array is", len(arreglo)) // we can print the length with len

  // Furthermore, we can make Matrix with our array
  var matriz [5][2]int // Matrix of 5 array with e elements
  matriz[0][1] = 2
  fmt.Println("the matrix is", matriz)
  fmt.Println("We are inide of", matriz[0][1])

  // This is a program that computes the average of a series os scores.
  var x [5]float64
      x[0] = 98
      x[1] = 99
      x[2] = 89
      x[3] = 90
      x[4] = 84

  var total float64 = 0
  for i := 0; i < 5; i++ {
    total += x[i]
  }
  // We convert the type of x that is int to float64
  fmt.Println("the average is", total / float64(len(x)))
}
