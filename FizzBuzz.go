package main //  this is know as a package declaration, and very Go program must start with it

import "fmt" //  this is how we include code from other package

func main()  {
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 && i % 5 == 0 {
      // divisible by 3 and 5
      fmt.Println(i, "FizzBuzz")
    } else if i % 3 == 0 {
      // divisible by 3
      fmt.Println(i, "Fizz")
    } else if i % 5 == 0 {
      // divisible by 5
      fmt.Println(i, "Buzz")
    } else {
      // Print all the numbers
      fmt.Println(i)
    }
  }
}
