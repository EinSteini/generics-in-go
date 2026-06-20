# Vergleich generischer Typparameter von Go und TypeScript

## 1. Einleitung

Generische Programmierung ermöglicht es, Algorithmen und Datenstrukturen typunabhängig zu formulieren, ohne dabei auf Typsicherheit zu verzichten. In diesem Projekt geht es darum Generics in Go zu verstehen und evaluieren, mit einem besonderen Fokus auf die Grenzen und Limitationen der Generics in Go. Desweiteren soll anhand konkreter Fallbeispiele ein Vergleich zu Generics in TypeScript ermöglicht werden.

## 2. Überblick: Generics in Go und TypeScript
TODO:
 - Vergleiche generische Typparameter in Go versus Typescript
 - Am Anfang, welche generischen Primitive gibt es
 - Wo im Programmcode kommen generische Typparameter vor? Dazu kleine Tabelle. Kanonische Beispiele
 
### Go

In Go wurden Generics mit der Version 1.18 eingeführt, um den Code typunabhängiger und für mehrere Typen wiederverwendbar zu machen, während Code-Duplikationen vermieden und Typsicherheit gewährleistet werden. Typische Anwendungsfälle für Generics sind klassische Datenstrukturen wie ein Stack, eine List oder Funktionen für das Sortieren oder Mapping von Elementen.

Die Generics basieren auf **Typparameteren**, die für Funktionen oder Typen definiert werden können. Diese Typparameter stellen Platzhalter für konkrete Typen dar und erlauben die Wiederverwendung des Codes. Die Typparameter werden dabei in eckigen Klammern `[]` angegeben und anstelle eines konkreten Datentyps wird typischerweise der Parameter `T` verwendet:

```go
type Stack[T any] struct{ items []T }

func Print[T any](value T) { ... }
```

Beim späteren Verwenden der Funktionen und Typen kann ein konkreter Datentyp dann als Typargument, wiederrum in eckigen Klammern, übergeben werden, um die Funktion (bzw. den Typen) zu instanziieren:

```go
var intStack Stack[int]
```

Ein weiteres Kernelement der Generics in Go sind **Type Constraints**, die festlegen, welche Typen als Argumente für generische Typparameter zulässig sind. Diese Constraints werden über Interfaces definert. Somit beschreiben Interfaces seit der Einführund der Generics nicht nur die benötigten Methoden für einen Typen, sondern sie können auch eine Menge erlaubter Typen angeben, das **Type Set**.

Beispiel:

```go
type Ordered interface {
  Integer|Float|~string
}

func MinNamed[T Ordered](x, y T) T { ... }

func MinLiteral[S ~[]E, E any]
```

Neben Typparametern und Constraints wurde auch die Typinterferenz eingeführt, um die Verwendung von Generics einfacher zu gestalten. Durch die Typinterferenz müssen die Typargumente meist nicht explizit angegeben werden, sondern der Compiler kann die Typargumente aus den Funktionsargumenten ableiten:

```go
var a, b, m float64
m = GMin[float64](a, b) // explicit type argument

m = GMin(a, b) // no type argument, still valid
```

**Quellen:**

- [An Introduction To Generics](https://go.dev/blog/intro-generics)

### TypeScript

TODO

### Vergleich

TODO Tabelle

## 3. Untersuchung der Bijektivität der Übersetzung: Go ↔ TypeScript

### Motivation

Eine interessante Frage beim Vergleich zweier Typsysteme ist, ob eine Hin- und Rückübersetzung des Codes das ursprüngliche Ergebnis reproduziert.

Wenn das Ergebnis semantisch und strukturell identisch zum Original ist, deutet das auf eine hohe Äquivalenz in der Umsetzung der Generics beider Sprachen hin. Weicht das Ergebnis ab, offenbart das Unterschiede in den Typsystemen.

### Testbeispiele

TODO: 
- evtl. manuell übersetzen und vergleichen
- Immer vermuten, was passiert und mögliche Probleme vorher aufzeigen

#### Generische Primitive

#### Type Sets vs. Type Union

#### Call by value & Pointer vs call by reference

#### Klassen

#### if type == bool

### Prompt für den Test

#### Schritt 1: Go → TypeScript

> **Prompt (Go → TS):**
>
> Übersetze den folgenden Go-Code nach TypeScript. Behalte die gleiche Struktur, Semantik und Kommentare bei. Verwende idiomatisches TypeScript, das dem Go-Code so nahe wie möglich kommt. Behalte generische Typparameter, Constraints und die gleiche öffentliche API bei.
>
> ```
> [Go-Code]
> ```
>
> Anforderungen:
> - Übersetze Go-Interfaces in TypeScript-Interfaces
> - Übersetze Go-Structs mit Methoden in TypeScript-Klassen
> - Übersetze Go Type Constraints in TypeScript `extends`-Constraints
> - Übersetze `(T, bool)` Return-Typen in `[T, boolean]` Tupel
> - Übersetze Go-Konstruktorfunktionen in TypeScript-Konstruktoren
> - Behalte alle Kommentare bei

#### Schritt 2: TypeScript → Go

> **Prompt (TS → Go):**
>
> Übersetze den folgenden TypeScript-Code zurück nach Go. Behalte die gleiche Struktur, Semantik und Kommentare bei. Verwende idiomatisches Go mit Generics (Go 1.18+), das dem TypeScript-Code so nahe wie möglich kommt.
>
> ```
> [TypeScript-Code]
> ```
>
> Anforderungen:
> - Übersetze TypeScript-Interfaces in Go-Interfaces
> - Übersetze TypeScript-Klassen in Go-Structs mit Methoden (Pointer-Receiver)
> - Übersetze TypeScript `extends`-Constraints in Go Interface-Constraints
> - Übersetze `[T, boolean]` Tupel in Go Multiple Return Values `(T, bool)`
> - Übersetze TypeScript-Konstruktoren in Go-Konstruktorfunktionen (`NewXxx`)
> - Behalte alle Kommentare bei
> - Verwende Groß-/Kleinschreibung für exportierte/nicht-exportierte Symbole

### Verwandte Arbeiten
TODO

### Fazit

TODO: Wann geht's schief? Wann klappt's?
Generell Allgeimeiner: In welchem Fall sind die Abbildungen bijektiv?

Die Übersetzung zwischen Go und TypeScript Generics ist nicht bijektiv im strengen Sinne. Während die generischen Konzepte (Typparameter, Constraints, generische Funktionen und Typen) eine gute Korrespondenz aufweisen, führen die fundamentalen Unterschiede der Sprachen (nominale vs. strukturelle Typisierung, Pointer-Semantik) zu unvermeidbaren Abweichungen beim Roundtrip. Ein LLM veruscht, die semantische Intention zu bewahren, aber der resultierende Code wird warscheinlich syntaktisch vom Original abweichen.
