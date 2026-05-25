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

**Quellen:**

- [An Introduction To Generics](https://go.dev/blog/intro-generics)

## 3. Kompilierung von Generics

### Go

Am Beispiel einer einfachen generischen `Min`-Funktion soll nun gezeigt werden wie der Compiler mit dem generischen Code umgeht.

```go
type Ordered interface { ~int | ~float64 | ~string }

func Min[T Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

result := Min[int](3, 7)
```

#### Schritt 1: Typprüfung der generischen Funktion

Der Compiler prüft zunächst die generische Funktion `Min` auf Korrektheit. Dafür stellt der Compiler sicher, dass alle Operationen innerhalb der Funktion, hier `a < b`, für jeden Typ zulässig sind, der den Constraint `Ordered` erfüllt. Zudem wird geprüft, dass die Funktion die korrekte Syntax von Generics einhält.

#### Schritt 2: Instanziierung

Die Verwendung der Funktion mit konkreten Datentypen und Werten, z.B. `Min[int](3, 7)`, beginnt die Instanziierung:

1. **Substitution**: Zuerst ersetzt der Compiler alle Typparameter `T` in der Funktion durch ihr spezifisches Typargument, in diesem Fall `int`:

```go
// Konzeptionell erzeugt der Compiler eine spezialisierte Version:
func Min_int(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

2. **Constraint-Prüfung**: Der Compiler überprüft, dass jedes Typargument den gegebenen Constraint erfüllt.
   <br>Bezogen auf das Beispiel prüft der Compiler also, ob `int` im Type Set von `Ordered` enthalten ist. Da das Interface `Ordered` die Typen `~int | ~float64 | ~string` erlaubt und ist `int` als Typargument gültig. Im Gegensatz dazu würde `Min[bool](...)` nicht erlaubt sein und das Programm würde sich nicht kompilieren lassen.

#### Schritt 3: Monomorphisation der GC Shapes

Anstatt tatsächlich für jedes Typargument eine neue Implementierung zu erstellen arbeitet Go mit **GC Shapes**, die die gleiche Instanziierung einer generischen Funktion teilen können.

Die GC Shape eines Typen beschreibt wie dieser Typ aus Sicht des Garbage Collectors aussieht. Dabei teilen sich zwei Typen eine GC Shape, wenn sie entweder denselben zugrunde liegenden Typ haben oder beides Pointer-Typen sind. Für jede GC Shape erzeugt der Go Compiler dann eine eigene Implementierung des generischen Codes. Konkret heißt das:

```text
Min[int]     → GC Shape: int     → eigene Implementierung
Min[float64] → GC Shape: float64 → eigene Implementierung
Min[*Foo]    → GC Shape: *uint8  → geteilte Implementierung
Min[*Bar]    → GC Shape: *uint8  → geteilte Implementierung
```

Für `Min[int]` erzeugt der Compiler also einen eigene Implementierung für die GC Shape `int`.

#### Schritt 4: Dictionary als versteckter Parameter

Wenn sich mehrere Typen dieselbe Implementierung teilen (z.B. `*Foo` und `*Bar`), muss der generierte Code zur Laufzeit wissen, mit welchem konkreten Typ er es zu tun hat. Dafür fügt der Compiler bei jedem Aufruf einer generischen Funktion/Methode ein **Dictionary** hinzu. Dieses Dictionary wird zur Compile-Zeit statisch erzeugt und enthält alle relevanten Informationen über die Typargumente, darunter den instanziierten Typ selbst.

#### Zusammenfassung

Während die Generics es ermöglichen Duplikation im Programmcode zu verringern, wird bei der Kompilierung wieder Duplikation eingeführt, da der Compiler für jede einzigartige GC Shape eine eigene Implementierung des generischen Codes ersetellt.

**Quellen:**

- [Generics Implementation – Dictionaries and Gcshapes](https://github.com/golang/proposal/blob/master/design/generics-implementation-dictionaries-go1.18.md)
- [Generics Implementation - GC Shape Stenciling](https://github.com/golang/proposal/blob/master/design/generics-implementation-gcshape.md)
- [How Generics were implemented in Go 1.18](https://deepsource.com/blog/go-1-18-generics-implementation)
- [Go Generics – A Deep Dive](https://dev.to/leapcell/go-generics-a-deep-dive-1one)
- [Generics in Go — From Basics to Advanced for Senior Developers](https://medium.com/@sogol.hedayatmanesh/generics-in-go-from-basics-to-advanced-for-senior-developers-887790b018d0)

### TypeScript

TypeScript verwendet das Prinzip der **Type Erasure**: Alle Typinformationen, inklusive der generischen Typparameter, existieren ausschließlich zur Compile-Zeit und werden beim Transpilieren nach JavaScript vollständig gelöscht.

Während der Kompilierung können die Typparameter verwendet werden, um den Code auf Korrektheit zu prüfen und Beziehungen zwischen Input und Output von Funktionen herzustellen. Danach werden alle Typinformationen entfernt, sodass regulärer JavaScript-Code übrig bleibt.

```typescript
// TypeScript
function unwrap<T>(value: T): T {
  return value;
}
const x = unwrap<number>(42);

// Generiertes JavaScript
function unwrap(value) {
  return value;
}
const x = unwrap(42);
```

**Quellen:**

- [What Happens When You Use Generics in TypeScript Functions](https://medium.com/@AlexanderObregon/what-happens-when-you-use-generics-in-typescript-functions-df5c23085da0)
- [What is Type Erasure in TypeScript?](https://www.geeksforgeeks.org/typescript/what-is-type-erasure-in-typescript/)

## 4. Anwendungsbereiche von Generics in Go

Generics können in Go an verschiedenen Stellen im Code eingesetzt werden. Dieser Abschnitt zeigt die grundlegenden Anwendungsbereiche mit Fokus auf die jeweilige Syntax.

### Generische Funktionen

Generische Funktionen eignen sich für Algorithmen, die unabhängig von konkreten Typ funktionieren sollen, z.B. Sortieren, Filtern oder Transformieren. Die Typparameter stehen dabei direkt hinter dem Funktionsnamen in eckigen Klammern `[]` und können innerhalb der Funktion wiederverwendet werden.

```go
func Swap[T any](a, b T) (T, T) {
    return b, a
}

// Verwendung
x, y := 1, 2
x, y = Swap(x, y)

a, b := "hello", "world"
a, b = Swap(a, b)
```

### Generische Structs und Methoden

Generische Structs werden für typsichere Datenstrukturen verwendet, die für beliebige Elementtypen funktionieren sollen, z.B. Stack, Queue oder ein typisierter Cache. Die Typparameter stehen hinter dem Typnamen und können im gesamten Struct und seinen Methoden verwendet werden. Methoden eines generischen Structs erhalten den Typparameter direkt über den Receiver:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    if len(s.items) == 0 {
        var zero T
        return zero
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func (s *Stack[T]) Len() int  {
  return len(s.items)
}

// Verwendung
var s Stack[string]
s.Push("hello")
s.Push("world")
fmt.Println(s.Pop()) // "world"
```

Methoden können dabei keine eigenen neuen Typparameter einführen, sie sind auf die Typparameter des Receiver-Typs beschränkt.

### Generische Interfaces

Interfaces können eigene Typparameter tragen und beschreiben damit einen typsicheren Vertrag für beliebige Elementtypen. Das ermöglicht Abstraktionen, bei denen die konkrete Implementierung austauschbar ist, z.B. für unterschiedliche Datenbankbackends:

```go
type Repository[T any] interface {
    FindByID(id int) (T, error)
    Save(entity T) error
    Delete(id int) error
}
```

In Go werden Interfaces implizit implementiert: Ein Struct erfüllt das Interface automatisch, sobald es alle geforderten Methoden für den konkreten Typ bereitstellt:

```go
type UserRepository struct { ... }

func (r UserRepository) FindByID(id int) (User, error) { ... }
func (r UserRepository) Save(u User) error             { ... }
func (r UserRepository) Delete(id int) error           { ... }

var repo Repository[User] = UserRepository{}  // gültig, da alle Methoden vorhanden
```

### Eingebaute Generics

Neben den zuvor beschriebenen Anwendungbereichen von Generics, hat Go auch Generics direkt in die Sprache eingebaut. Beispiele hierfür sind Slices, Maps und Channels, die den Elementtypen als Typparameter verwenden und sich verhalten wie generische Structs:

```go
[]int          // entspricht etwa Slice[int]
map[string]int // entspricht etwa Map[string, int]
chan float64   // entspricht etwa Chan[float64]
```

Dazu kommen eingebaute Funktionen, die typunabhängig auf diesen Typen operieren:

```go
append(s []T, elems ...T) []T   // fügt Elemente an Slice an
len(v T) int                    // Länge von Slice, Map, Channel, String, Array
cap(v T) int                    // Kapazität von Slice, Channel, Array
make(t T, size ...int) T        // alloziert Slice, Map oder Channel
copy(dst, src []T) int          // kopiert Elemente zwischen Slices
delete(m map[K]V, key K)        // entfernt Eintrag aus Map
```

**Quellen:**

- [An Introduction To Generics](https://go.dev/blog/intro-generics)
- [The Go Programming Language Specification – Type parameters](https://go.dev/ref/spec#Type_parameter_declarations)
- [The Go Programming Language Specification – Built-in functions](https://go.dev/ref/spec#Built-in_functions)
- [When To Use Generics](https://go.dev/blog/when-generics)
- [Generics in Go](https://medium.com/@ksandeeptech07/generics-in-go-ffd5100ec642)
