// sucesión de Fibonacci con Golang
// Buscamos mostrar los primeros numeros de la sirie
// 0,1,1,2,3,5,8.13,21,34

package main

import (
        "fmt"
)

var (
        firstNumber         int = 1
        lastNumber          int = 0
        auxiliarNumber      int
)

func main() {


        for i := 1; i <= 20; i++ {
                fmt.Println(lastNumber)
                auxiliarNumber = firstNumber
                firstNumber = firstNumber + lastNumber
                lastNumber = auxiliarNumber

        }
        fmt.Println(auxiliarNumber)


}