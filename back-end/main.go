package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Operacion struct {
	Op1 int    `json:"op1"`
	Op2 int    `json:"op2"`
	Op  string `json:"op"`
	Res int    `json:"res"`
}

var historial []Operacion

func main() {
	fmt.Println("Hola back-end")
	router := mux.NewRouter()

	router.HandleFunc("/calcular", calculadoraHandler).Methods("POST")
	router.HandleFunc("/inicio", inicioHand).Methods("GET")
	router.HandleFunc("/historial", obtenerHistorial).Methods("GET") // Agregar ruta para historial

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("../front-end/public")))
	handler := cors.Default().Handler(router)

	http.Handle("/", router)
	http.ListenAndServe(":8000", handler)

	//op := Operacion{Op1: 100, Op2: 15, Op: "+"}
	//calculadora(&op)
	//fmt.Println(op.Res)
}

func inicioHand(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola desde el back")
}
func calculadoraHandler(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Hola mundo")
	fmt.Println("hola estoy aqui")
	var op Operacion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&op)
	fmt.Println(decoder)
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
	historial = append(historial, op)
}
func obtenerHistorial(w http.ResponseWriter, r *http.Request) {
	historialJSON, err := json.Marshal(historial)
	if err != nil {
		http.Error(w, "Error al convertir el historial a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(historialJSON)
	if err != nil {
		http.Error(w, "Error al escribir el historial JSON en la respuesta", http.StatusInternalServerError)
		return
	}
}
func calculadora(op *Operacion) error {

	switch op.Op {
	case "+":
		op.Res = op.Op1 + op.Op2
	case "-":
		op.Res = op.Op1 - op.Op2
	case "/":
		if op.Op2 == 0 {
			return errors.New("division por cero")
		} else {
			op.Res = op.Op1 / op.Op2
		}
	case "*":
		op.Res = op.Op1 * op.Op2
	default:
		return errors.New("operador no válido")
	}
	return nil
}
func validarRango(numero int) bool {
	return numero >= 0 && numero <= 99
}
