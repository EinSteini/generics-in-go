package main

import (
	"fmt" // Benötigt für das Drucken auf die Konsole und die String-Formatierung
)

// Container ist eine generische Struktur, die ein Slice von Elementen vom Typ T enthält.
// Entspricht der TypeScript-Klasse 'class Container<T>'.
type Container[T any] struct {
	Items []T // Öffentliches Feld, entspricht 'public items: T[]'
}

// NewContainer erstellt einen neuen Container mit den gegebenen Elementen.
// Dies dient als Äquivalent zum Konstruktor der TypeScript-Klasse.
// Das Standard-leere Array aus dem TypeScript-Konstruktor wird durch
// die Nil-Slice-Initialisierung von Go behandelt, falls keine Elemente bereitgestellt werden.
func NewContainer[T any](items []T) Container[T] {
	if items == nil {
		// Wenn keine Elemente bereitgestellt werden, initialisiere mit einem leeren Slice
		// für Konsistenz, obwohl ein Nil-Slice auch funktionieren würde.
		return Container[T]{Items: []T{}}
	}
	return Container[T]{Items: items}
}

// Map wendet eine Funktion 'f' auf jedes Element im Eingabe-Container 'c' an
// und gibt einen neuen Container mit den Ergebnissen zurück.
// Entspricht der TypeScript-Funktion 'function Map<T, U>(c: Container<T>, f: (value: T) => U): Container<U>'.
func Map[T, U any](c Container[T], f func(value T) U) Container[U] {
	// Erstelle einen neuen Slice zum Speichern der gemappten Ergebnisse, mit der gleichen Länge wie die Eingabe.
	// Entspricht 'const resultItems: U[] = new Array<U>(c.items.length);'
	resultItems := make([]U, len(c.Items))

	// Erstelle einen neuen Container für die Ergebnisse, initialisiert mit dem resultItems-Slice.
	// Entspricht 'const result = new Container<U>(resultItems);'
	result := NewContainer(resultItems) // Verwende NewContainer für Konsistenz

	// Iteriere über die Elemente im Eingabe-Container und wende die Mapping-Funktion an.
	// Entspricht der TypeScript-For-Schleife.
	for i := 0; i < len(c.Items); i++ {
		result.Items[i] = f(c.Items[i])
	}

	return result
}

// mainGoEquivalent demonstriert die Verwendung von Container und Map in Go.
// Entspricht der TypeScript-Funktion 'function mainGoEquivalent(): void'.
func mainGoEquivalent() {
	// Erstelle einen neuen Container von Integern.
	// Entspricht 'const nums = new Container<number>([1, 2, 3]);'
	nums := NewContainer([]int{1, 2, 3})
	fmt.Println("Original numbers container:", nums.Items) // Ausgabe: [1 2 3]

	// Mappe die Zahlen zu Strings, indem "!" an jede angehängt wird.
	// Entspricht 'const strs = Map(nums, (n: number) => `${n}!`);'
	strs := Map(nums, func(n int) string {
		return fmt.Sprintf("%d!", n)
	})
	fmt.Println("Mapped strings container:", strs.Items) // Erwartete Ausgabe: [1! 2! 3!]
}

// main ist der Einstiegspunkt für das Go-Programm.
func main() {
	mainGoEquivalent()
}