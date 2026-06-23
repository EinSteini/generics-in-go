//go:build ignore

package main

import "fmt"

// Snippet 5: Type Sets mit Tilde-Operator
// Testet: Type Set Constraint, Operator-Nutzung (+), benannte Typen (~)

type Addable interface {
	~int | ~float64 | ~string
}

func Add[T Addable](a, b T) T {
	return a + b
}

type Meter float64

func main() {
	fmt.Println(Add(3, 4))         // 7
	fmt.Println(Add("Go", "Lang")) // GoLang

	var d1, d2 Meter = 10, 20
	fmt.Println(Add(d1, d2)) // 30 (benannter Typ dank ~)
}