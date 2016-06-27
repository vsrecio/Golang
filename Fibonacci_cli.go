// sucesión de Fibonacci con Golang
// Buscamos mostrar los primeros numeros de la sucesión
// 0,1,1,2,3,5,8.13,21,34

package main

import (
        "fmt"
)

var (
        firstNumber         int = 1     // El valor inicial de la susesion
        lastNumber          int = 0     // Valor anterior al valor inicial
        auxiliarNumber      int
)

func main() {

        for i := 1; i <= 20; i++ {
                fmt.Println(lastNumber)                         // Mostramos el numero anterior al numero inicial
                auxiliarNumber = firstNumber                    // Guardamos el valor del numero inicial
                firstNumber = firstNumber + lastNumber
                lastNumber = auxiliarNumber                     // Usamos el valor anteriormente guardado
        }
        fmt.Println(auxiliarNumber)

}