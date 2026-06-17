
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

- [What Happens When You Use Generics in TypeScript Functions](https://medium.com/@AlexanderObregon/what-happens-when-you-use-generics-in-typescript-functions-df5c23085da0))
- [What is Type Erasure in TypeScript?](https://www.geeksforgeeks.org/typescript/what-is-type-erasure-in-typescript/)
