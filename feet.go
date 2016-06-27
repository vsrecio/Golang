package main

import "fmt"


func main()  {
  fmt.Print("Introduce la cantidad en pies: ")
  var feet float64
  fmt.Scanf("%f", &feet)

  meters := feet * 0.3048

  fmt.Println(meters)
}
