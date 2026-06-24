package main

import "fmt" // Benötigt für fmt.Println

// interface Pair<T> {
//   First: T;
//   Second: T;
// }
// Eine generische Struktur, die zwei Werte des gleichen Typs speichert.
type Pair[T any] struct {
	First  T
	Second T
}

// function SwapWithCopy<T>(p: Pair<T>): Pair<T> {
//   return { First: p.Second, Second: p.First };
// }
// Tauscht die Elemente eines Paares und gibt ein *neues* Paar mit den getauschten Werten zurück.
// Das ursprüngliche Paar wird nicht verändert.
func SwapWithCopy[T any](p Pair[T]) Pair[T] {
	return Pair[T]{First: p.Second, Second: p.First}
}

// function SwapWithRef<T>(p: Pair<T>): void {
//   const temp = p.First;
//   p.First = p.Second;
//   p.Second = temp;
// }
// Tauscht die Elemente eines Paares *an Ort und Stelle* durch Modifikation des Originalpaares
// über einen Zeiger.
func SwapWithRef[T any](p *Pair[T]) { // Beachte den Zeiger '*' um die Originalstruktur zu modifizieren
	temp := p.First
	p.First = p.Second
	p.Second = temp
}

// (function main() { ... })();
// Der Hauptausführungsblock in Go.
func main() {
	// const p: Pair<number> = { First: 10, Second: 20 };
	p := Pair[int]{First: 10, Second: 20}

	// const q = SwapWithCopy(p);
	q := SwapWithCopy(p)

	// console.log(p.First, p.Second); // Erwartete Ausgabe: 10 20 (unverändert!)
	fmt.Println(p.First, p.Second) // Erwartete Ausgabe: 10 20 (unverändert!)

	// console.log(q.First, q.Second);
	fmt.Println(q.First, q.Second)

	// SwapWithRef(p);
	// Übergibt die Adresse von p, damit SwapWithRef es modifizieren kann.
	SwapWithRef(&p)

	// console.log(p.First, p.Second);
	fmt.Println(p.First, p.Second)
}
