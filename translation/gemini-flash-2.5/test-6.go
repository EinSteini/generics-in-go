package main

import (
	"fmt"
	"golang.org/x/exp/constraints" // Benötigt Go 1.18+ für Generics-Constraints
)

// Ptr[T] repräsentiert eine generische, zeigerähnliche Struktur,
// die einen Wert des Typs T enthält.
// Dies entspricht dem TypeScript-Objekt { value: T }.
type Ptr[T any] struct {
	Value T
}

// Swap tauscht die Werte, die in zwei Ptr-Strukturen enthalten sind.
// Die Funktion nimmt Zeiger auf Ptr-Strukturen entgegen,
// um deren Inhalte direkt zu modifizieren, was dem Verhalten
// von Objektreferenzen in TypeScript entspricht.
func Swap[T any](a, b *Ptr[T]) {
	// Original bleibt unverändert
	// const temp = a.value;
	temp := a.Value
	// a.value = b.value;
	a.Value = b.Value
	// b.value = temp;
	b.Value = temp
}

// Double verdoppelt einen numerischen Wert.
// Der Typ-Parameter T muss ein numerischer Typ sein (Ganzzahl oder Gleitkommazahl).
// 'constraints.Signed | constraints.Unsigned | constraints.Float' deckt alle
// Standard-Go-Zahlentypen ab, auf denen Multiplikationen sinnvoll sind,
// ähnlich wie 'extends number' in TypeScript.
func Double[T constraints.Signed | constraints.Unsigned | constraints.Float](v T) T {
	// Original bleibt unverändert
	return v * 2 // Original bleibt unverändert
}

func main() {
	// let x: Ptr<number> = { value: 10 };
	// In Go verwenden wir einen Zeiger auf die Ptr-Struktur,
	// um die Modifikation durch die Swap-Funktion zu ermöglichen,
	// was dem Objektreferenzverhalten von TypeScript entspricht.
	x := &Ptr[int]{Value: 10}
	// let y: Ptr<number> = { value: 20 };
	y := &Ptr[int]{Value: 20}

	// Swap(x, y);
	Swap(x, y)

	// console.log(x.value, y.value); // Erwartet: 20 10
	// Go dereferenziert x und y automatisch für den Feldzugriff (x.Value ist gültig für *Ptr[int]).
	fmt.Println(x.Value, y.Value) // Erwartet: 20 10

	// let n = 5;
	n := 5

	// console.log(Double(n), n); // Erwartet: 10 5
	fmt.Println(Double(n), n) // Erwartet: 10 5
}