# Vergleich generischer Typparameter von Go und TypeScript

## 1. Einleitung

Generische Programmierung ermöglicht es, Algorithmen und Datenstrukturen typunabhängig zu formulieren, ohne dabei auf Typsicherheit zu verzichten. Statt für jeden Typ eine eigene Implementierung zu schreiben, können allgemeine und wiederverwendbare Lösungen geschaffen werden.

Dieses Projekt vergleicht, wie Go und TypeScript dieses Konzept umsetzen, wo sich diese Umsetzungen unterscheiden, und ob eine KI-gestützte Übersetzung zwischen beiden Sprachen semantisch bijektiv ist.

## 2. Überblick: Generics in Go und TypeScript

Im Zentrum der generischen Programmierung steht das Konzept des **generischen Typparameters**. Hierbei handelt es sich um einen Platzhalter, typischerweise `T`, der anstelle eines Datentyps verwendet wird, bis ein konkreter Typ bei der Instanziierung gebunden wird:

```go
func Print[T any](value T)  // T ist ungebunden
Print[int](42)              // T wird zu int gebunden
```

Während generische Typparameter als Feature in Go erst mit Version 1.18 (2022) und in TypeScript mit Version 2.0 (2016) eingeführt wurden, sind mehrere eingebaute Primitive beider Sprachen bereits generisch definiert, u.a.:

| Sprache | Primitive |
| --- | --- |
| Go | `[]T` (Slice), `[N]T` (Array), `map[K]V`, `chan T`, `*T` (Pointer) |
| TypeScript | `Array<T>`, `Map<K, V>`, `Set<T>`, `Promise<T>`, `Record<K, V>` |

Das Konzept des generischen Typparameters ist damit in beiden Sprachen schon lange vertraut.

Um eigene generische Typen und Funktionen zu definieren, teilen beide Sprachen zwei weitere Kernkonzepte:

- **Type Constraints**: Schränken ein, welche Typen als Argumente für `T` zulässig sind.
- **Typinferenz**: Typargumente müssen meist nicht explizit angegeben werden, sondern der Compiler kann diese aus den Funktionsargumenten ableiten.

Die folgenden Abschnitte gehen weiter auf die genaue Syntax in den beiden Sprachen ein.

### Go

Der generische Typparameter `T` wird in eckigen Klammern `[]` angegeben:

```go
type Stack[T any] struct{ items []T }

func Print[T any](value T) { ... }
```

Die **Type Constraints** werden in Go als Interface angegeben, direkt hinter dem Typparameter `T`:

```go
type Ordered interface {
  Integer|Float|~string
}

func Min[T Ordered](a, b T) T { ... }
```

Bei der Instanziierung wird ein konkreter Datentyp als **Typargument** übergeben. Bei Funktionsaufrufen kann es durch Typinferenz auch weggelassen werden:

```go
var intStack Stack[int]        // Typdefinition: T muss angegeben werden

var a, b float64
m := Min[float64](a, b)        // Funktionsaufruf: explizites Typargument
m := Min(a, b)                 // Funktionsaufruf: T wird zu float64 inferiert
```

**Quellen:**

- [An Introduction To Generics](https://go.dev/blog/intro-generics)

### TypeScript

Statt eckiger Klammern verwendet TypeScript spitze Klammern `<>` zur Deklaration von Typparameter `T`:

```typescript
class Stack<T> { items: T[] = [] }

function print<T>(value: T): void { ... }
```

Type Constraints werden in TypeScript mit dem Schlüsselwort `extends` hinter dem Typparameter `T` angegeben:

```typescript
type Ordered = number | string

function min<T extends Ordered>(a: T, b: T): T { ... }
```

Wie auch in Go kann das Typargument explizit übergeben oder bei Funktionsaufrufen durch Typinferenz weggelassen werden:

```typescript
const intStack = new Stack<number>(); // Typdefinition: T muss angegeben werden

let a: number, b: number;
min<number>(a, b); // Funktionsaufruf: explizites Typargument
min(a, b);         // Funktionsaufruf: T wird zu number inferiert
```

TypeScript unterstützt zudem Standardwerte für Typparameter, z.B.x `type Stack<T = string>`. Dadurch wird das Typargument optional: Wird keines angegeben, nimmt `T` den Standardtyp `string` an.

**Quellen:**

- [TypeScript Handbook – Generics](https://www.typescriptlang.org/docs/handbook/2/generics.html)

### Syntaktische Positionen eines generischen Typparameters

Die folgende Tabelle fasst zusammen, an welchen Stellen im Programmcode ein generischer Typparameter `T` vorkommen kann und wie Go und TypeScript das jeweils ausdrücken:

| Position im Code             | Go                                           | TypeScript                                       |
| ---------------------------- | -------------------------------------------- | ------------------------------------------------ |
| Funktionsdefinition          | `func Min[T Ordered](a, b T) T`              | `function min<T extends Ordered>(a: T, b: T): T` |
| Arrow Function               | —                                            | `const id = <T>(x: T): T => x`                   |
| Typdefinition (Struct/Class) | `type Stack[T any] struct{ items []T }`      | `class Stack<T> { items: T[] = [] }`             |
| Type Alias (ab Go 1.24)      | `type MyStack[T any] = Stack[T]`             | `type MyStack<T> = Stack<T>`                     |
| Interface-Definition         | `type Container[T any] interface{ Get() T }` | `interface Container<T> { get(): T }`            |
| Typparameter im Constraint   | `func Clone[S ~[]E, E any](s S) S`           | `function get<T, K extends keyof T>(o: T, k: K)` |
| Methode (T vom Typ definiert)          | `func (s *Stack[T]) Push(v T)`               | `push(v: T): void`                               |
| Methode (eigener Typparameter)          | nicht möglich                            | `map<U>(fn: (x: T) => U): U`                     |

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
