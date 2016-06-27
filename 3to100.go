package main

import "fmt"

func main()  {
  for i := 1; i <= 100; i ++ {
    if i % 3 == 0 { // Si i es un numero divisible entre 3
      fmt.Println(i)
    }
  }

}
