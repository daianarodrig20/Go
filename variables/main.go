package main

import "fmt"

var numero int
var texto string
// Por defecto bool inicializa en false si quisiera forzarlo a que inicie en tru lo defino como var status bool = true
var status bool = true

func main() {
    numero1, numero2, numero3, texto, status := 1, 2, 10, "Que locura", false

    fmt.Println(numero1)
    fmt.Println(numero2)
    fmt.Println(numero3)
    fmt.Println(texto)
    fmt.Println(status)

    mostrarStatus()
}

func mostrarStatus() {
    fmt.Println(status)
}
