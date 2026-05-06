# Anwendungsbereiche von Generics in Go

Generics (Typparameter) wurden in Go primär eingeführt, um Code-Duplikation zu vermeiden und Typsicherheit zu gewährleisten, ohne auf das unsichere `interface{}` zurückgreifen zu müssen. 

Die sinnvollen Einsatzgebiete von Generics in Go lassen sich in **fünf architektonische Hauptkategorien** unterteilen:

### 1. Generische Datenstrukturen (Container / Collections)
Vor der Einführung von Generics musste man für jeden Datentyp eine eigene Struktur schreiben (z. B. `IntStack`, `StringStack`) oder Typzusicherungen zur Laufzeit vornehmen. Generics ermöglichen den Bau von wiederverwendbaren Datenstrukturen.
*   **Beispiele:** Sets, Stacks, Queues, verkettete Listen (Linked Lists), Bäume (Trees), Caches.

```go
// Ein generischer Binärbaum  [1]
type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] { /* ... */ }

var stringTree Tree[string]
```

### 2. Code Deduplizierung bei Algorithmen
Oft muss dieselbe Logik auf Slices oder Maps unterschiedlicher Datentypen angewendet werden. Anstatt für Datentypen jeweils eigene Schleifen zu schreiben, können generische Funktionen diese Arbeit zentralisieren.
- **Beispiele:** `Filter`, `Map`, `Reduce`, `Sort`, `Reverse`, Extrahieren von Keys/Values aus Maps, Summierung.

Das offizielle Go-Tutorial [2] zeigt das Problem anhand zweier nahezu identischer Funktionen:

```go
// OHNE Generics: Duplizierter Code für jeden Map-Typ
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
```

Mit Generics lässt sich das in einer einzigen Funktion zusammenfassen:

```go
// MIT Generics: Eine Funktion für alle numerischen Typen [2]
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
```

### 3. Typsichere API-Wrapper & DTOs
In modernen Anwendungen gibt es oft standardisierte Envelopes für HTTP-Antworten oder Datenbankergebnisse. Die Metadaten (wie Status oder Fehlermeldungen) bleiben gleich, aber die eigentliche Payload variiert.
*   **Beispiele:** HTTP JSON-Responses, Pagination-Ergebnisse, Datenbank-Resultate (`Result[T]`).

```go
// Ein standardisiertes API-Antwortformat
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    T      `json:"data,omitempty"` // Die Nutzdaten sind generisch
}

// Anwendungsbeispiel:
// userRes := APIResponse[User]{Success: true, Data: User{Name: "Alice"}}
```

### 4. Nebenläufigkeit
Kanäle (`channels`) sind in Go strikt typisiert. Wenn man allgemeine Muster für Nebenläufigkeit schreiben möchte (z.B. mehrere Kanäle zu einem zusammenführen), geht das dank Generics nun typsicher für jede Art von Channel.
*   **Beispiele:** Fan-In (Zusammenführen von Channels), Fan-Out, Worker-Pools.

```go
// Führt zwei Kanäle eines beliebigen Typs zu einem einzigen Kanal zusammen (Fan-In)
func Merge[T any](ch1, ch2 <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		for v := range ch1 { out <- v }
	}()
	go func() {
		for v := range ch2 { out <- v }
	}()
	return out
}
```

### 5. Hilfs- & Utility-Funktionen
Generics eignen sich hervorragend für kleine Helferfunktionen, die typische syntaktische Umständlichkeiten ("Boilerplate") in Go auflösen.
*   **Beispiele:** Zeiger (Pointer) aus Literalen generieren, Null-Werte überprüfen (`Coalesce`), grundlegende Mathematik (`Max`/`Min`).

Das `GMin`-Beispiel aus dem offiziellen Go Blog [1] zeigt eine generische Min-Funktion:

```go
import "golang.org/x/exp/constraints"

// Generisches Minimum – funktioniert für alle geordneten Typen (aus [1])
func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Aufruf:
// GMin[int](2, 3)    → 2
// GMin[float64](2.5, 1.1) → 1.1
// Dank Type Inference auch: GMin(2, 3)
```

Weitere nützliche Hilfsfunktionen:

```go
// Erstellt sicher einen Pointer aus einem literalen Wert (z.B. Ptr("Text"))
func Ptr[T any](v T) *T {
	return &v
}

// Gibt den ersten Wert zurück, wenn er kein Null-Wert ist, ansonsten den Fallback
func Coalesce[T comparable](val, fallback T) T {
	var zero T
	if val != zero {
		return val
	}
	return fallback
}
```

---

## Quellen

1. **Go Blog – "An Introduction to Generics"** (Robert Griesemer, Ian Lance Taylor, 22. März 2022)  
   https://go.dev/blog/intro-generics  
   *Enthält: `GMin[T constraints.Ordered]`, `Tree[T]`, Type Sets, Instantiation*

2. **Go Tutorial – "Getting started with generics"**  
   https://go.dev/doc/tutorial/generics  
   *Enthält: `SumIntsOrFloats`, Type Constraints mit Union-Types, das `Number` Interface*

3. **Go Blog – "When To Use Generics"** (Ian Lance Taylor, 12. April 2022)  
   https://go.dev/blog/when-generics  
   *Empfehlungen: Generics bei Containern, allgemeinen Datenstrukturen und Funktionen auf Slices/Maps*

4. **Go Type Parameters Proposal** (Design-Dokument)  
   https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md  
   *Das vollständige Design-Dokument mit Begründungen und Einschränkungen*

5. **Go Language Specification – Type Parameter Declarations**  
   https://go.dev/ref/spec#Type_parameter_declarations

6. **samber/lo** – Generische Utility-Library für Go (inspiriert von Lodash)  
   https://github.com/samber/lo  
   *Praxisbeispiel für `Filter`, `Map`, `Reduce` und weitere generische Hilfsfunktionen*