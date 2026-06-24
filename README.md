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

| Sprache    | Primitive                                                                    |
| ---------- | ---------------------------------------------------------------------------- |
| Go         | `[]T` (Slice), `[N]T` (Array), `map[K]V`, `chan T`, `*T` (Pointer) |
| TypeScript | `Array<T>`, `Map<K, V>`, `Set<T>`, `Promise<T>`, `Record<K, V>`    |

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
min(a, b); // Funktionsaufruf: T wird zu number inferiert
```

TypeScript unterstützt zudem Standardwerte für Typparameter, z.B.x `type Stack<T = string>`. Dadurch wird das Typargument optional: Wird keines angegeben, nimmt `T` den Standardtyp `string` an.

**Quellen:**

- [TypeScript Handbook – Generics](https://www.typescriptlang.org/docs/handbook/2/generics.html)

### Syntaktische Positionen eines generischen Typparameters

Die folgende Tabelle fasst zusammen, an welchen Stellen im Programmcode ein generischer Typparameter `T` vorkommen kann und wie Go und TypeScript das jeweils ausdrücken:

| Position im Code                     | Go                                             | TypeScript                                         |
| ------------------------------------ | ---------------------------------------------- | -------------------------------------------------- |
| Funktionsdefinition                  | `func Print[T any](v T)`                     | `function print<T>(v: T): void`                  |
| Funktionsdefinition (mit Constraint) | `func Min[T Ordered](a, b T) T`              | `function min<T extends Ordered>(a: T, b: T): T` |
| Arrow Function                       | —                                             | `const id = <T>(x: T): T => x`                   |
| Typdefinition (Struct/Class)         | `type Stack[T any] struct{ items []T }`      | `class Stack<T> { items: T[] = [] }`             |
| Type Alias (ab Go 1.24)              | `type MyStack[T any] = Stack[T]`             | `type MyStack<T> = Stack<T>`                     |
| Interface-Definition                 | `type Container[T any] interface{ Get() T }` | `interface Container<T> { get(): T }`            |
| T im Constraint von S                | `func Clone[S ~[]T, T any](s S) S`           | `function get<T, S extends keyof T>(o: T, k: S)` |
| Methode (T vom Typ definiert)        | `func (s *Stack[T]) Push(v T)`               | `push(v: T): void`                               |
| Methode (eigener Typparameter)       | nicht möglich                                 | `map<U>(fn: (x: T) => U): U`                     |
| Default Type Parameter               | —                                             | `function wrap<T = string>(v: T): T`             |

## 3. Untersuchung der Bijektivität der Übersetzung: Go ↔ TypeScript

In Abschnitt 2 wurde gezeigt, dass Go und TypeScript generische Typparameter syntaktisch unterschiedlich verwenden, inhaltlich aber ähnliche Konzepte abbilden. Das wirft die Frage auf, ob eine Übersetzung zwischen beiden Sprachen bijektiv ist, d.h. ob eine Hin- und Rückübersetzung das ursprüngliche Ergebnis reproduzieren kann.

Ein semantisch und strukturell identisches Ergebnis deutet dabei auf eine hohe Äquivalez in der Umsetzung der generischen Programmierung beider Sprachen hin. Abweichungen hingegen offenbaren grundlegende Unterschiede in den Typsystemen.

In diesem Abschnitt werden daher mehrere Testbeispiele dargestellt, um die Bijektivität der Übersetzung gestützt durch das LLM "Gemini Flash 2.5" zu untersuchen.

### Testbeispiele

Für die folgenden Testbeispiele wird jeweils zuerst der Go-Quellcode gezeigt und eine Vermutung formuliert, welche Probleme bei der Übersetzung in TypeScript-Quellcode auftreten könnten.
Anschließend wird das LLM mit dem unten definierten Prompt zur Übersetzung und Rückübersetzung aufgefordert.

Zwischen jeder Anfrage wird der Kontext zurückgesetzt, damit das Modell sich nicht an den vergangenen Nachrichten orientieren kann.
Desweiteren werden Kommentare entfernt, welche auf die ursprüngliche Version hinweisen könnten.

---

#### Generische Funktionsdefinition in Kombination mit generischen Primitiven

```go
func Repeat[T any](v T, n int) []T {
	result := make([]T, n)
	for i := range result {
		result[i] = v
	}
	return result
}

func main() {
	fmt.Println(Repeat("go", 3)) // [go go go]
	fmt.Println(Repeat(42, 2))   // [42 42]
}
```

**Vermutung:**
Dieses Beispiel sollte problemlos übersetzbar sein.

- `any` entspricht direkt TypeScript's generischem `<T>` ohne Constraint.
- `make([]T, n)` wird zu `new Array<T>(n)` oder einem Array-Literal.

Die Hin- und Rückübersetzung sollte semantisch identisch sein.
Lediglich `fmt.Println` vs. `console.log` und die Slice-Erzeugung unterscheiden sich syntaktisch.

#### Ergebnisse

**Go → TypeScript:**

```typescript
export function Repeat<T>(v: T, n: number): T[] {
  const result: T[] = new Array(n);
  for (let i = 0; i < n; i++) {
    result[i] = v;
  }
  return result;
}

console.log(Repeat("go", 3));
console.log(Repeat(42, 2));
```

**TypeScript → Go (Rückübersetzung):**

```go
func Repeat[T any](v T, n int) []T {
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = v
	}
	return result
}

func main() {
	fmt.Println(Repeat("go", 3))
	fmt.Println(Repeat(42, 2))
}
```

**Beobachtungen:** Wie vermutet ist die Übersetzung nahezu bijektiv:

1. Die Semantik ist vollständig erhalten.
2. Syntaktischer Unterschied: `for i := range result` (Original) → `for i := 0; i < n; i++` (Rückübersetzung). Funktional jedoch identisch.

---

#### Type Sets in Go vs. Type Union in TypeScript

```go
type Addable interface {
	~int | ~float64 | ~string
}

func Add[T Addable](a, b T) T {
	return a + b
}

type Meter float64

func main() {
	fmt.Println(Add(3, 4))         // 7
	fmt.Println(Add("Go", "Lang")) // GoLang

	var d1, d2 Meter = 10, 20
	fmt.Println(Add(d1, d2))       // 30 (benannter Typ dank ~)
}
```

**Vermutung:**
Hier werden mehrere Probleme erwartet:

1. **Operator-Constraint:**
   Go erlaubt `a + b` nur, weil das Type Set garantiert, dass alle enthaltenen Typen `+` unterstützen. Dabei wird `T` zur Compile-Zeit an einen einzigen konkreten Typ gebunden.

   TypeScript hat keine Möglichkeit, Operator-Unterstützung in einem Constraint auszudrücken.
   Das LLM wird vermutlich `T extends number | string` schreiben, aber dies bedeutet nur, dass `T` einer der beiden Typen sein kann. TypeScript muss daher davon ausgehen, dass `a` und `b` unterschiedliche Typen haben könnten, weshalb es `a + b` nicht generisch erlaubt.
   Das LLM wird daher vermutlich eine Type Assertion (`as any`) oder Overloads benötigen.
2. **Tilde-Operator (`~`):**
   Go's `~float64` erlaubt auch benannte Typen wie `Meter` für den Typparameter zu initialisieren. TypeScript hat kein Äquivalent und `type Meter = number` erzeugt keinen eigenständigen Typ, sondern nur ein Alias.
   Der Unterschied zwischen nominaler (`~`) und struktureller Typisierung geht verloren.
3. **Rückübersetzung:** Aus TypeScript zurück nach Go wird der Tilde-Operator wahrscheinlich fehlen, und die `Meter`-Nutzung geht verloren.

#### Ergebnisse

**Go → TypeScript:**

```typescript
type Addable = number | string;

function Add<T extends Addable>(a: T, b: T): T {
  return a + b; // throws error
}

type Meter = number;

function main() {
  console.log(Add(3, 4)); // 7
  console.log(Add("Go", "Lang")); // GoLang

  const d1: Meter = 10;
  const d2: Meter = 20;
  console.log(Add(d1, d2)); // 30
}

main();
```

**TypeScript → Go (Rückübersetzung):**

```go
type Addable interface {
	~int | ~float32 | ~float64 | ~string
}

func Add[T Addable](a, b T) T {
	return a + b
}

type Meter int

func main() {
	fmt.Println(Add(3, 4))
	fmt.Println(Add("Go", "Lang"))
	d1 := Meter(10)
	d2 := Meter(20)
	fmt.Println(Add(d1, d2))
}
```

**Beobachtungen:**

1. **TypeScript-Code kompiliert nicht:** `return a + b` mit `T extends number | string` erzeugt einen TypeScript-Fehler, da `a` und `b` unterschiedliche Typen sein könnten und somit der Rückgabetyp `T`nicht garantiert werden kann.
2. **Tilde in Rückübersetzung vorhanden:** Überraschenderweise fügt das LLM bei der Rückübersetzung `~` hinzu (`~int | ~float32 | ~float64 | ~string`), obwohl TypeScript dieses Konzept nicht hat. Das LLM hat hier offenbar Go-Idiome angewandt.
3. **`Meter` wird zu `int` statt `float64`:** Im Original ist `type Meter float64`, die Rückübersetzung erzeugt `type Meter int`. TypeScript's `number` ist nicht eindeutig auf `int` oder `float64` abbildbar.
4. **Zusätzlicher Typ `float32`:** Die Rückübersetzung fügt `~float32` zum Constraint hinzu, das im Original nicht existiert.
5. **Initialisierung geändert:** `var d1, d2 Meter = 10, 20` (Original) → `d1 := Meter(10)` (Rückübersetzung). Funktional äquivalent, aber syntaktisch verschieden.

---

#### Call-by-Value & Pointer vs. Call-by-Reference

TODO Call-by-Reference kommt hier nicht so wirklich drin vor oder?

```go
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func Double[T ~int | ~float64](v T) T {
	return v * 2 // Original bleibt unverändert
}

func main() {
	x, y := 10, 20
	Swap(&x, &y)
	fmt.Println(x, y) // 20 10

	n := 5
	fmt.Println(Double(n), n) // 10 5
}
```

**Vermutung:** Dieses Beispiel ist problematisch für die Übersetzung:

1. **Zeiger-Parameter (`*T`):**
   TypeScript hat keine Zeiger.
   Das LLM muss entweder Wrapper-Objekte `{ value: T }` einführen oder die Signatur grundlegend ändern (z.B. Rückgabe eines Tupels statt In-Place-Mutation).
2. **Semantischer Unterschied:**
   In Go mutiert `Swap` die Originalvariablen über Zeiger.
   In TypeScript gibt es keine Möglichkeit, primitive Werte (`number`) by-reference zu übergeben.
   Die Übersetzung wird die Semantik entweder verändern oder einen Wrapper einführen, der im Original nicht existiert.
3. **Rückübersetzung:**
   Wenn das LLM einen Wrapper `{value: T}` in TypeScript erzeugt, wird die Rückübersetzung wahrscheinlich keinen Zeiger-Code rekonstruieren, sondern den Wrapper als Struct beibehalten.
4. **Double:** Die Value-Semantik (`n` bleibt 5) ist in TypeScript automatisch gegeben, da primitive Typen ohnehin by-value übergeben werden. Hier sollte die Übersetzung korrekt sein.

#### Ergebnisse

**Go → TypeScript:**

```typescript
type Ptr<T> = { value: T };

function Swap<T>(a: Ptr<T>, b: Ptr<T>): void {
  const temp = a.value;
  a.value = b.value;
  b.value = temp;
}

function Double<T extends number>(v: T): T {
  return (v * 2) as T;
}

function main() {
  let x: Ptr<number> = { value: 10 };
  let y: Ptr<number> = { value: 20 };
  Swap(x, y);
  console.log(x.value, y.value); // 20 10

  let n = 5;
  console.log(Double(n), n); // 10 5
}

main();
```

**TypeScript → Go (Rückübersetzung):**

```go
type Ptr[T any] struct {
	Value T
}

func Swap[T any](a, b *Ptr[T]) {
	temp := a.Value
	a.Value = b.Value
	b.Value = temp
}

func Double[T constraints.Signed | constraints.Unsigned | constraints.Float](v T) T {
	return v * 2
}

func main() {
	x := &Ptr[int]{Value: 10}
	y := &Ptr[int]{Value: 20}
	Swap(x, y)
	fmt.Println(x.Value, y.Value) // 20 10
	n := 5
	fmt.Println(Double(n), n) // 10 5
}
```

**Beobachtungen:**

1. **Wrapper-Objekt bestätigt:** Wie vermutet führt das LLM einen `Ptr<T>`-Typ mit `value`-Feld ein, da TypeScript keine Zeiger hat.
2. **Semantik verändert:** Im Original ist `Swap[T any](a, b *T)` – direkte Zeiger auf Werte. Die Rückübersetzung erzeugt `Swap[T any](a, b *Ptr[T])` – Zeiger auf Wrapper-Structs. Der Aufrufer muss nun `Ptr`-Objekte erstellen statt einfach `&x` zu schreiben.
3. **Nicht bijektiv:** Statt `*T` entsteht `*Ptr[T]` – ein strukturell anderes Programm.
4. **Externe Abhängigkeit:** Die Rückübersetzung importiert `golang.org/x/exp/constraints` – eine externe Bibliothek, die im Original nicht benötigt wird. Der Inline-Constraint `~int | ~float64` wird durch `constraints.Signed | constraints.Unsigned | constraints.Float` ersetzt, was deutlich mehr Typen abdeckt.
5. **Double korrekt:** Die Value-Semantik wird korrekt beibehalten.

---

#### Klassen (Struct-Embedding)

```go
type Named struct {
	Name string
}

type Box[T any] struct {
	Named
	Value T
}

func NewBox[T any](name string, v T) Box[T] {
	return Box[T]{Named: Named{Name: name}, Value: v}
}

func main() {
	b := NewBox("answer", 42)
	fmt.Println(b.Name, b.Value) // answer 42
}
```

TODO: Das hier drunter noch besser als richtige Vermutungen fomulieren

**Vermutung:** Struct-Embedding hat kein direktes TypeScript-Äquivalent:

1. **Embedding vs. Vererbung:** Das LLM muss zwischen `class Box<T> extends Named` (Vererbung) und Komposition (`name: Named`) wählen. Beides hat Nachteile: Vererbung ändert die Hierarchie, Komposition verliert den direkten Feldzugriff (`b.Name`).
2. **Konstruktor:** Go's `NewBox`-Funktion wird wahrscheinlich zu einem `constructor` in einer TypeScript Klasse. Die Rückübersetzung sollte das zu einer `NewBox`-Funktion rekonstruieren.
3. **Direkter Feldzugriff:** In Go kann man `b.Name` schreiben (promoted field). Bei Komposition in TypeScript wäre es `b.named.name`. Das LLM wird vermutlich Vererbung wählen, um diesen Zugriff zu erhalten.

#### Ergebnisse

**Go → TypeScript:**

```typescript
class Named {
  Name: string;
  constructor(name: string = "") {
    this.Name = name;
  }
}

class Box<T> extends Named {
  Value: T;
  constructor(name: string, value: T) {
    super(name);
    this.Value = value;
  }
}

function NewBox<T>(name: string, v: T): Box<T> {
  return new Box<T>(name, v);
}

function main() {
  const b = NewBox("answer", 42);
  console.log(b.Name, b.Value); // answer 42
}

main();
```

**TypeScript → Go (Rückübersetzung):**

```go
type Named struct {
	Name string
}

type Box[T any] struct {
	Named
	Value T
}

func NewBox[T any](name string, value T) Box[T] {
	return Box[T]{
		Named: Named{Name: name},
		Value: value,
	}
}

func main() {
	b := NewBox("answer", 42)
	fmt.Println(b.Name, b.Value) // answer 42
}
```

TODO stand beim ersten nicht auch "nahezu" bijektiv? ist das nicht quasi das gleiche wie "fast perfekt"?
**Beobachtungen:** Dies ist das einzige Beispiel mit fast perfekter Bijektivität:

1. **Vererbung gewählt:** Wie vermutet wählt das LLM `extends Named` statt Komposition, um den direkten Feldzugriff `b.Name` zu erhalten.
2. **Rückübersetzung nahezu identisch:** Das Embedding wird korrekt rekonstruiert. `Named` wird wieder eingebettet, `NewBox` bleibt eine Funktion.

---

#### Methoden mit eigenen Typparametern

```go
type Container[T any] struct {
	items []T
}

// Nicht möglich in Go: Methoden dürfen keine eigenen Typparameter einführen
// func (c Container[T]) Map[U any](f func(T) U) Container[U] { ... }

// Workaround: Top-Level-Funktion mit zwei Typparametern
func Map[T any, U any](c Container[T], f func(T) U) Container[U] {
	result := Container[U]{items: make([]U, len(c.items))}
	for i, v := range c.items {
		result.items[i] = f(v)
	}
	return result
}

func main() {
	nums := Container[int]{items: []int{1, 2, 3}}

	// Nicht möglich in Go:
	// nums.Map(func(n int) string { return fmt.Sprintf("%d!", n) })

	strs := Map(nums, func(n int) string {
		return fmt.Sprintf("%d!", n)
	})
	fmt.Println(strs.items) // [1! 2! 3!]
}
```

**Vermutung:**

1. **Workaround-Übersetzung:** Vermutlich wird das LLM versuchen den Workaround zu übersetzen, obwohl TypeScript eine Methode mit eigenen Typparametern unterstützen würde.
2. **Rückübersetzung:** Da TypeScript die Methode `map<U>` direkt ausdrücken kann, wird das LLM beim Zurückübersetzen wahrscheinlich versuchen, `func (c Container[T]) Map[U any](...)` zu schreiben. Der Workaround als Top-Level-Funktion wird vermutlich nicht rekonstruiert.
3. **Typinferenz:** Bei der Go→TS-Übersetzung sollte Typinferenz für `U` in beiden Richtungen funktionieren.

---

### Zusammenfassung der Ergebnisse

| Testfall                   | Bijektiv? | Hauptproblem                                                                      |
| -------------------------- | --------- | --------------------------------------------------------------------------------- |
| Generische Primitive       | ✅ Ja     | Nur syntaktische Unterschiede (Loop-Stil)                                         |
| Type Sets                  | ❌ Nein   | TS-Code kompiliert nicht (`a + b`), `Meter` wird zu `int` statt `float64` |
| Pointer/Call-by-value      | ❌ Nein   | Wrapper-Struct `Ptr[T]` statt Zeiger, externe Abhängigkeit eingefügt          |
| Struct-Embedding           | ✅ Fast   | Nahezu identisch, nur Default-Parameter-Unterschied                               |
| Methoden mit Typparametern | ❓ Offen  | Vermutung: Rückübersetzung erzeugt ungültigen Go-Code                          |

### Prompt für den Test

#### Schritt 1: Go → TypeScript

> **Prompt (Go → TS):**
>
> Übersetze den folgenden Go-Code nach TypeScript. Behalte die gleiche Struktur, Semantik und Kommentare bei. Verwende idiomatisches TypeScript, das dem Go-Code so nahe wie möglich kommt. Behalte generische Typparameter, Constraints und die gleiche öffentliche API bei.
>
> ```
> [Go-Code]
> ```

#### Schritt 2: TypeScript → Go

> **Prompt (TS → Go):**
>
> Übersetze den folgenden TypeScript-Code zurück nach Go. Behalte die gleiche Struktur, Semantik und Kommentare bei. Verwende idiomatisches Go mit Generics (Go 1.18+), das dem TypeScript-Code so nahe wie möglich kommt.
>
> ```
> [TypeScript-Code]
> ```

### Fazit

Die Übersetzung zwischen Go und TypeScript Generics ist **nicht bijektiv** im strengen Sinne. Die Experimente zeigen ein klares Muster:

**Wann klappt's?**
TODO die nummern hier passen nicht zu den oben beschriebenen beispielen (also da sind keine 7 stück und das erste gehört zu test 4)

- Einfache generische Funktionen mit `any`-Constraint sind nahezu perfekt übersetzbar (Test 4).
- Struct-Embedding ↔ Klassen-Vererbung funktioniert gut, solange keine Pointer-/Value-Semantik involviert ist (Test 7).
- Die grundlegende Struktur (Typparameter, Funktionssignaturen, generische Container) wird in allen Fällen korrekt übertragen.

**Wann geht's schief?**

- **Type Sets mit Operatoren**: TypeScript kann Operator-Constraints nicht ausdrücken. Der generierte Code kompiliert teilweise nicht, und der Tilde-Operator (`~`) sowie benannte Typen gehen bei der Hinübersetzung verloren.
- **Zeiger-Semantik**: Da TypeScript keine Zeiger hat, entstehen Wrapper-Structs, die bei der Rückübersetzung nicht zu einfachen Zeigern rekonstruiert werden. Die Aufruf-Semantik ändert sich grundlegend.
- **Type Switches mit Struct-Typen**: Die Unterschiede zwischen Go's nominalem Typsystem und TypeScript's strukturellem Typsystem werden sichtbar. `typeof` funktioniert nur für Primitive; für Structs braucht es `instanceof`, was zu `*Dog` statt `Dog` führt. Zusätzlich erzeugt die Rückübersetzung eine Explosion an Integer-Cases, da TypeScript's `number` nicht eindeutig auf einen Go-Typ abbildbar ist.

**Allgemein:** Die Abbildung ist bijektiv, wenn der Code ausschließlich Konzepte nutzt, die in beiden Sprachen strukturelle Äquivalente haben: generische Funktionen ohne Operator-Nutzung, einfache Constraints, und Datenstrukturen ohne Zeiger-Semantik. Sobald sprachspezifische Features involviert sind (Type Sets mit `~`, Zeiger, nominale Typunterscheidung), ist die Rückübersetzung nicht mehr äquivalent zum Original.
