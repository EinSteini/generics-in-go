# Generics in Go vs TypeScript: Eine vergleichende Fallstudie

## 1. Einleitung

Generische Programmierung ermöglicht es, Algorithmen und Datenstrukturen typunabhängig zu formulieren, ohne dabei auf Typsicherheit zu verzichten. In diesem Projekt geht es darum Generics in Go zu verstehen und evaluieren, mit einem besonderen Fokus auf die Grenzen und Limitationen der Generics in Go. Desweiteren soll anhand konkreter Fallbeispiele ein Vergleich zu Generics in TypeScript ermöglicht werden.

## 2. Überblick: Generics in Go und TypeScript

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

Neben Typparametern und Constraints wurde auch die Typinterferenz eingeführt, um die Verwendung von Generics einfacher zu gestalten. Durch die Typinterferenz müssen die Typargumente meist nicht explizit angegeben werden, sondern der Compiler kaann die Typargumente aus den Funktionsargumenten ableiten:

```go
var a, b, m float64
m = GMin[float64](a, b) // explicit type argument

m = GMin(a, b) // no type argument, still valid
```
