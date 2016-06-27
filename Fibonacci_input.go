// sucesión de Fibonacci con Golang, buscamos mostrar los primeros 10 numeros de la sucesión
// 0,1,1,2,3,5,8,13,21,34

package main

import (
    "fmt"
)

var firstNum        int             // capturamos el valor en la linea de comandos
var auxNum          int
var inputStr        string = "Introduce un numero entero positivo: "

func main() {

    fmt.Print(inputStr)             // Mostrar el string que nos pide el valor en el terminal
    fmt.Scanf("%d", &firstNum)      // Capturamos el valor introducido por el usuario y lo enviamos a la variable fiersNum via el puntero
    lastNum := firstNum - 1         // Tomamos el valor anterior restando un munero al anteriormente introducido

    for i := 0; i <= 10; i++ {
        fmt.Println(lastNum)
        auxNum = firstNum
        firstNum += lastNum
        lastNum = auxNum
    }

    fmt.Println(auxNum)

}
