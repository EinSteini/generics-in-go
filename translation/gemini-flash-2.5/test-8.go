package main

import (
	"fmt"
	"math"
)

// Dog class
// Entspricht der TypeScript-Klasse Dog.
type Dog struct {
	Name string
}

// NewDog ist eine konstruktorähnliche Funktion für Dog.
// Entspricht dem TypeScript-Konstruktor `constructor(public Name: string)`.
func NewDog(name string) *Dog {
	return &Dog{Name: name}
}

// String implementiert das fmt.Stringer-Interface für Dog.
// Entspricht der TypeScript-Methode `toString(): string`.
func (d *Dog) String() string {
	return d.Name
}

// Describe-Funktion
// Entspricht der TypeScript-Funktion Describe(v: unknown): string.
// Der Typ 'any' (Go 1.18+) ist das Äquivalent zu TypeScript's 'unknown'.
func Describe(v any) string {
	switch val := v.(type) {
	case bool: // Entspricht `typeof v === 'boolean'`
		if val {
			return "YES"
		}
		return "NO"

	// Entspricht `typeof v === 'number' && Number.isInteger(v)`.
	// Go erfordert die Prüfung spezifischer Integer-Typen.
	case int:
		return fmt.Sprintf("#%d", val)
	case int8:
		return fmt.Sprintf("#%d", val)
	case int16:
		return fmt.Sprintf("#%d", val)
	case int32:
		return fmt.Sprintf("#%d", val)
	case int64:
		return fmt.Sprintf("#%d", val)
	case uint:
		return fmt.Sprintf("#%d", val)
	case uint8:
		return fmt.Sprintf("#%d", val)
	case uint16:
		return fmt.Sprintf("#%d", val)
	case uint32:
		return fmt.Sprintf("#%d", val)
	case uint64:
		return fmt.Sprintf("#%d", val)
	case uintptr: // Weniger häufig für Benutzerdaten, aber ein Integer-Typ.
		return fmt.Sprintf("#%d", val)

	// Behandelt Gleitkomma-Typen, wenn sie einen Integer darstellen.
	// TypeScript's `Number.isInteger` funktioniert auch für Gleitkommazahlen wie 42.0.
	case float32:
		// Prüft, ob der float32-Wert keinen Nachkommaanteil hat.
		if val == float32(math.Trunc(float64(val))) {
			return fmt.Sprintf("#%d", int32(val)) // Konvertiert zu int32 für die Formatierung
		}
		// Wenn es keine Integer-Gleitkommazahl ist, fällt es in den Default-Fall.
	case float64:
		// Prüft, ob der float64-Wert keinen Nachkommaanteil hat.
		if val == math.Trunc(val) { // `math.Trunc` gibt den Integer-Teil der Gleitkommazahl zurück.
			return fmt.Sprintf("#%d", int64(val)) // Konvertiert zu int64 für die Formatierung
		}
		// Wenn es keine Integer-Gleitkommazahl ist, fällt es in den Default-Fall.

	case *Dog: // Entspricht `v instanceof Dog`
		return fmt.Sprintf("Dog: %s", val.Name)

	default: // Entspricht `String(v)` für jeden anderen Typ in TypeScript.
		return fmt.Sprint(v)
	}
}

func main() {
	// Entspricht der TypeScript-Funktion `main`.
	fmt.Println(Describe(true))          // Erwartet: YES
	fmt.Println(Describe(42))            // Erwartet: #42
	fmt.Println(Describe(42.0))          // Entspricht Number.isInteger(42.0) in TS, erwartet: #42
	fmt.Println(Describe(4.5))           // Entspricht String(4.5) in TS, erwartet: 4.5
	fmt.Println(Describe(NewDog("Rex"))) // Entspricht new Dog("Rex"), erwartet: Dog: Rex
	fmt.Println(Describe("hello"))       // Erwartet: hello
	fmt.Println(Describe([]int{1, 2}))   // Beispiel für den Default-Fall, erwartet: [1 2]
}