package main

import (
	"fmt"
)

type Operacion struct {
	Op1 int
	Op2 int
	Op  string
	Res int
}

func main() {
	fmt.Println("Hola back-end")

	op := Operacion{Op1: 100, Op2: 15, Op: "+"}
	calculadora(&op)
	fmt.Println(op.Res)
}

func calculadora(op *Operacion) {

	if !validarRango(op.Op1) || !validarRango(op.Op2) {
		fmt.Println("Los operandos son numeros de mas de 2 digitos")
		return
	}

	switch op.Op {
	case "+":
		op.Res = op.Op1 + op.Op2
	case "-":
		op.Res = op.Op1 - op.Op2
	case "/":
		op.Res = op.Op1 / op.Op2
	case "*":
		op.Res = op.Op1 * op.Op2
	default:
		fmt.Println("fin switch")
	}
}
func validarRango(numero int) bool {
	return numero >= 0 && numero <= 99
}
