package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ---------------- Declaración de Variables : ----------------
	// Con inferencia
	miNumero := 9
	miNombre := "David"
	miDecimal := 9.81

	// Con tipo predefinido
	var otroNumero int
	var otroNombre string
	var otroDecimal float32 = 3.1416

	otroNumero = 2

	fmt.Printf("%v %v %v\n%v %v %v\n", miNumero, miNombre, miDecimal, otroNumero, otroNombre, otroDecimal)

	// --------------------- Tipos básicos : ---------------------
	// Todos los tipos básicos de Go son:
	// bool
	// string
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias de uint8
	// rune // alias de int32
	//      // representa caracter Unicode
	// float32 float64
	// complex64 complex128

	// ----------------- Conversiones de Tipos : -----------------
	// A veces, requerimos convertir tipos de variables. Esto lo
	// hacemos de varias formas, por ejemplo:
	numero := 30
	cadena := strconv.Itoa(numero)
	fmt.Printf("%T - %v\n", cadena, cadena)

	// ------------------- Arreglos y slices : -------------------
	// Declaración de arreglos
	var arreglo [3]int
	arreglo[0] = 1
	arreglo[1] = 2
	arreglo[2] = 3

	otroArreglo := [3]int{3, 2, 1}

	fmt.Println(arreglo)
	fmt.Println(arreglo[1])
	fmt.Println(otroArreglo)

	// Declaración de Slices
	// Sin make():
	var sliceDeInt []int
	_ = append(sliceDeInt, 3)
	fmt.Println(sliceDeInt) // <-- ¿Qué sucederá?

	// Con make():
	sliceDeRune := make([]rune, 3)
	for i := range sliceDeRune {
		sliceDeRune[i] = 'a'
	}
	fmt.Println(sliceDeRune)
}
