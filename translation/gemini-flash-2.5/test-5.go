package main

import "fmt"

// Addable repräsentiert Typen, die mit dem '+' Operator addiert oder verkettet werden können.
// In TypeScript war dies `number | string`.
// Der `~` Operator (Tilde) ermöglicht, dass benannte Typen (wie z.B. `Meter` unten),
// deren zugrundeliegender Typ `int`, `float32`, `float64` oder `string` ist,
// diese Schnittstellen-Constraint erfüllen.
type Addable interface {
	~int | ~float32 | ~float64 | ~string
}

// Add ist eine generische Funktion, die zwei Werte desselben Addable-Typs addiert.
// Der ursprüngliche TypeScript-Code enthielt den Kommentar "// throws error" neben `return a + b;`.
// In Go ist der `+` Operator für Typen, die durch `Addable` eingeschränkt sind,
// wohldefiniert (für String-Verkettung und numerische Addition). Daher führt diese Zeile
// in Go nicht zu einem Kompilierungsfehler.
func Add[T Addable](a, b T) T {
	return a + b // Ursprünglicher TS-Kommentar: "// throws error". In Go ist dies gültig.
}

// Meter ist ein benannter Typ, basierend auf 'number' in TypeScript.
// Hier verwenden wir 'int' als zugrundeliegenden Typ, da die Beispielwerte (10, 20) Ganzzahlen sind.
// In TypeScript: `type Meter = number;`
type Meter int

func main() {
	// console.log(Add(3, 4));         // 7
	fmt.Println(Add(3, 4))

	// console.log(Add("Go", "Lang")); // GoLang
	fmt.Println(Add("Go", "Lang"))

	// const d1: Meter = 10;
	// const d2: Meter = 20;
	// console.log(Add(d1, d2));       // 30 (benannter Typ dank ~)
	d1 := Meter(10) // Explizite Konvertierung zum benannten Typ Meter
	d2 := Meter(20)
	fmt.Println(Add(d1, d2))
}