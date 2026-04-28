# Anwendungsbereiche von Generics in Go

Generics (Typparameter) wurden in Go primär eingeführt, um Code-Duplikation zu vermeiden und Typsicherheit zu gewährleisten, ohne auf das unsichere `interface{}` zurückgreifen zu müssen. 

Die sinnvollen Einsatzgebiete von Generics in Go lassen sich in **fünf architektonische Hauptkategorien** unterteilen:

### 1. Generische Datenstrukturen (Container / Collections)
Vor der Einführung von Generics musste man für jeden Datentyp eine eigene Struktur schreiben (z. B. `IntStack`, `StringStack`) oder Typzusicherungen zur Laufzeit vornehmen. Generics ermöglichen den Bau von wiederverwendbaren Datenstrukturen.
*   **Beispiele:** Sets, Stacks, Queues, verkettete Listen (Linked Lists), Bäume (Trees), Caches.

```go
// Ein Set für jeden vergleichbaren Typ
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

// Anwendungsbeispiel:
// intSet := make(Set[int])
// stringSet := make(Set[string])
```
Ein Problem ist natürlich, dass das ganze nicht mit eigenen Typen funktioniert.

### 2. Algorithmische Code-Deduplizierung (Slice- & Map-Utilities)
Oft muss dieselbe Logik auf Slices oder Maps unterschiedlicher Datentypen angewendet werden. Anstatt für Datentypen jeweils eigene Schleifen zu schreiben, können generische Funktionen diese Arbeit zentralisieren.
- **Beispiele:** `Filter`, `Map`, `Reduce`, `Sort`, `Reverse`, Extrahieren von Keys/Values aus Maps.

```go
// Filtert ein beliebiges Slice basierend auf einer übergebenen Bedingung
func Filter[T any](items []T, condition func(T) bool) []T {
	var result []T
	for _, item := range items {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}
```

### 3. Typsichere API-Wrapper & Datenumschläge (DTOs)
In modernen Anwendungen gibt es oft standardisierte "Briefumschläge" (Envelopes) für HTTP-Antworten oder Datenbankergebnisse. Die Metadaten (wie Status oder Fehlermeldungen) bleiben gleich, aber die eigentlichen Nutzdaten (Payload) variieren.
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

### 4. Nebenläufigkeitsmuster (Concurrency & Channel-Utilities)
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

### 5. Hilfs- & Utility-Funktionen (Boilerplate-Reduzierung)
Generics eignen sich hervorragend für kleine Helferfunktionen, die typische syntaktische Umständlichkeiten ("Boilerplate") in Go auflösen.
*   **Beispiele:** Zeiger (Pointer) aus Literalen generieren, Null-Werte überprüfen (`Coalesce`), grundlegende Mathematik (`Max`/`Min`).

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