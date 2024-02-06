package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Operacion struct {
	Op1 int    `json:"op1"`
	Op2 int    `json:"op2"`
	Op  string `json:"op"`
	Res int    `json:"res"`
}

func main() {
	fmt.Println("Hola back-end")
	router := mux.NewRouter()

	router.HandleFunc("/calcular", calculadoraHandler).Methods("POST")
	router.HandleFunc("/inicio", inicioHand).Methods("GET")

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("../front-end/public")))

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)

	//op := Operacion{Op1: 100, Op2: 15, Op: "+"}
	//calculadora(&op)
	//fmt.Println(op.Res)
}
func inicioHand(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola desde el back")
}
func calculadoraHandler(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Hola mundo")
	var op Operacion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&op)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
		return
	}
	if !validarRango(op.Op1) || !validarRango(op.Op2) {
		http.Error(w, "Los operandos son números de más de 2 dígitos", http.StatusBadRequest)
		return
	}
	calculadora(&op)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(op)

}
func calculadora(op *Operacion) {

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
