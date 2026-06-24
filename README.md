# Vergleich generischer Typparameter von Go und TypeScript

## Zusammenfassung

### Inhalt

Diese Arbeit untersucht die Unterschiede und Äquivalenzen generischer Typparameter in Go und TypeScript.
Zunächst werden dazu die Grundlagen generischer Programmierung in beiden Sprachen verglichen: Syntax der Typparameter (`[T any]` in Go vs. `<T>` in TypeScript), Constraints (`interface` mit Type Sets vs. `extends`), Typinferenz und die Positionen, an denen generische Typparameter im Code auftreten können (Funktionen, Structs/Klassen, Interfaces, Type Aliases, Methoden). 
Dabei werden sowohl Gemeinsamkeiten als auch fundamentale Unterschiede herausgearbeitet. Zu zweiteren gehört, dass TypeScript Methoden mit eigenen Typparametern erlaubt, Go hingegen nicht, oder dass Go mit Type Sets und dem Tilde-Operator (`~`) ein mächtigeres Constraint-System besitzt.

### Vorgehen

Um die Äquivalenz der beiden Generics-Implementierungen zu prüfen, wird eine Untersuchung der Bijektivität durchgeführt: Fünf Go-Testbeispiele werden mithilfe des LLM „Gemini Flash 2.5" nach TypeScript übersetzt und das Ergebnis anschließend zurück nach Go übersetzt. Die Testfälle decken gezielt verschiedene Aspekte ab:

1. **Generische Primitive** – Einfache generische Funktion mit Slice (`[]T` ↔ `T[]`)
2. **Type Sets** – Interface-Constraint mit Tilde-Operator und Operator-Garantie (`~int | ~float64 | ~string`)
3. **Pointer-Semantik** – Generischer Struct mit Call-by-Value vs. Call-by-Reference (`Pair[T]` vs. `*Pair[T]`)
4. **Struct-Embedding** – Nicht-generisches Struct eingebettet in generisches Struct (Komposition vs. Vererbung)
5. **Methoden mit eigenen Typparametern** – Go-Workaround mit freistehender Funktion (`Map[T, U any]`)

Für jeden Testfall wird vorab eine Vermutung über mögliche Übersetzungsprobleme formuliert und anschließend mit dem tatsächlichen Ergebnis verglichen.

### Ergebnisse

Die Untersuchung zeigt ein differenziertes Bild:

- **Generische Typparameter selbst sind robust übersetzbar:** Typparameter-Deklaration, Multi-Parameter-Generics, generische Structs/Klassen und Funktionsparameter mit generischen Typen werden in allen fünf Tests korrekt hin- und zurückübersetzt. Die Abbildung `[T any]` ↔ `<T>`, `[]T` ↔ `T[]` und `func(T) U` ↔ `(value: T) => U` ist verlustfrei.

- **Type Sets und Operator-Constraints brechen die Bijektivität:** Go's `~int | ~float64 | ~string` wird zu TypeScript's `number | string` vereinfacht. Der Tilde-Operator und die nominale Typunterscheidung gehen verloren. TypeScript kann keine Operator-Unterstützung in Constraints ausdrücken, weshalb `a + b` einen Kompilierfehler erzeugt. Bei der Rückübersetzung halluziniert das LLM zusätzliche Typen.

- **Pointer-Semantik wird durch Kontextwissen kompensiert:** Obwohl TypeScript keine Zeiger kennt, rekonstruiert das LLM die Value-/Pointer-Unterscheidung korrekt, vermutlich begünstigt durch explizite Funktionsnamen.

- **Struct-Embedding wird über Vererbung abgebildet:** Das LLM wählt `extends` statt Komposition, wodurch der direkte Feldzugriff erhalten bleibt und die Rückübersetzung das Embedding korrekt rekonstruiert.

- **Workarounds bleiben erhalten:** Das LLM behält Go's freistehende Funktion `Map[T, U any]` bei, obwohl TypeScript eine Methode `map<U>(...)` ermöglichen würde.

**Gesamtfazit:** Die generischen Typparameter und ihre grundlegende Syntax sind zwischen Go und TypeScript in vielen Fällen bijektiv übersetzbar. Die Verluste entstehen dort, wo Go's Constraint-System (Type Sets, `~`, Operator-Garantien) über das hinausgeht, was TypeScript's `extends`-Constraints ausdrücken können. Umgekehrt werden TypeScript's zusätzliche Möglichkeiten (Methoden mit eigenen Typparametern) durch die Übersetzung nicht ausgeschöpft. Die Abbildung ist somit bijektiv für die Typparameter selbst, aber nicht für das umgebende Typsystem.


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

Ein semantisch und strukturell identisches Ergebnis deutet dabei auf eine hohe Äquivalenz in der Umsetzung der generischen Programmierung beider Sprachen hin. Abweichungen hingegen offenbaren grundlegende Unterschiede in den Typsystemen.

In diesem Abschnitt werden daher mehrere Testbeispiele dargestellt, um die Bijektivität der Übersetzung gestützt durch das LLM "Gemini Flash 2.5" zu untersuchen.

### Testbeispiele

Für die folgenden Testbeispiele wird jeweils zuerst der Go-Quellcode gezeigt und eine Vermutung formuliert, welche Probleme bei der Übersetzung in TypeScript-Quellcode auftreten könnten.
Anschließend wird das LLM mit dem unten definierten Prompt zur Übersetzung und Rückübersetzung aufgefordert.

Zwischen jeder Anfrage wird der Kontext zurückgesetzt, damit das Modell sich nicht an den vergangenen Nachrichten orientieren kann.
Desweiteren werden Kommentare entfernt, welche auf die ursprüngliche Version hinweisen könnten.

---

#### Generische Funktionsdefinition in Kombination mit generischen Primitiven

Dieses Beispiel kombiniert einen generischen Typparameter `[T any]` mit Go's generischem Slice-Primitiv `[]T`. Beide Konzepte haben in TypeScript direkte Gegenstücke (`<T>` und `T[]`), weshalb eine verlustfreie Übersetzung erwartet wird.

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

Dieses Beispiel sollte problemlos übersetzbar sein:

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

**Beobachtungen:**

Wie vermutet ist die Übersetzung bis auf einige syntaktische Unterschiede (for-Schleife) übersetzbar.
Insbesondere die generischen Komponenten konnten vollständig wiederhergestellt werden.

---

#### Type Sets in Go vs. Type Union in TypeScript

Dieses Beispiel nutzt ein Interface-Constraint mit Type Set (`~int | ~float64 | ~string`) und den Tilde-Operator für benannte Typen. Dies sind Konzepte, die auf Go's Operator-Constraints und nominales Typsystem setzen und in TypeScript kein direktes Äquivalent haben.

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

Hier können mehrere Probleme erwartet werden:

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

1. Die generische Type Constraint wurde wieder reproduziert, in TypeScript wurde sie wie erwartet mit einem `extends` modelliert.
2. Der TypeScript-Code kompiliert nicht: `return a + b` mit `T extends number | string` erzeugt einen TypeScript-Fehler, da `a` und `b` unterschiedliche Typen sein könnten und somit der Rückgabetyp `T` nicht garantiert werden kann.
3. Das LLM stellt bei der Rückübersetzung der Type Sets `~` wieder her, obwohl TypeScript dieses Konzept nicht hat.
4. Die Rückübersetzung fügt `~float32` zum Constraint hinzu, das im Original nicht existiert.
5. `Meter` wird zu `int` statt `float64`.

---

#### Call-by-Value & Pointer vs. Call-by-Reference

Dieses Beispiel verwendet einen generischen Struct `Pair[T]` mit Zeiger- und Wert-Übergabe (`*Pair[T]` vs. `Pair[T]`). Dies ist eine Unterscheidung, die in TypeScript nicht existiert, da Objekte stets per Referenz übergeben werden.

```go
type Pair[T any] struct {
	First, Second T
}

func SwapWithCopy[T any](p Pair[T]) Pair[T] {
	p.First, p.Second = p.Second, p.First
	return p
}

func SwapWithRef[T any](p *Pair[T]) {
	p.First, p.Second = p.Second, p.First
}

func main() {
	p := Pair[int]{First: 10, Second: 20}

	q := SwapWithCopy(p)
	fmt.Println(p.First, p.Second) // 10 20 (unverändert!)
	fmt.Println(q.First, q.Second) // 20 10 (neue Kopie)

	SwapWithRef(&p)
	fmt.Println(p.First, p.Second) // 20 10 (mutiert)
}
```

**Vermutung:**

Dieses Beispiel ist hauptsächlich problematisch für die Übersetzung, da Go und TypeScript verschiedenes Verhalten bei der Parameterübergabe zeigen:

1. **`SwapCopy` (by-value):** In Go wird das Struct kopiert, `p` bleibt also unverändert und nur die Rückgabe enthält die getauschten Werte. In TypeScript werden Objekte immer by-reference übergeben. Wenn das LLM die Funktion naiv übersetzt, mutiert sie das Original, wodurch die Value-Semantik verloren geht. Es müsste einen expliziten Clone einfügen.
2. **`SwapInPlace` (by-pointer):** In TypeScript ist Mutation über Referenz der Default für Objekte, es sind also kein Zeiger nötig. Das LLM kann hier einfach das Objekt direkt mutieren. Die Frage ist, ob bei der Rückübersetzung der `*`-Parameter rekonstruiert wird.
3. **Gegensätzliche Defaults:** Go arbeitet standardmäßig mit Pass-By-Value, TypeScript mit Pass-By-Reference. Was in Go zwei verschiedene Funktionen mit unterschiedlichem Verhalten sind, könnte in TypeScript zu zwei identisch wirkenden Funktionen werden.
Wenn TypeScript keinen Clone für `SwapCopy` enthält, wird die Rückübersetzung wahrscheinlich auch keinen Value-Parameter (`Pair[T]`) rekonstruieren, sondern einen Pointer (`*Pair[T]`), sodass dann beide Funktionen identisch werden.

#### Ergebnisse:

**Go → TypeScript:**

```typescript
interface Pair<T> {
  First: T;
  Second: T;
}

function SwapWithCopy<T>(p: Pair<T>): Pair<T> {
  return { First: p.Second, Second: p.First };
}

function SwapWithRef<T>(p: Pair<T>): void {
  const temp = p.First;
  p.First = p.Second;
  p.Second = temp;
}

(function main() {
  const p: Pair<number> = { First: 10, Second: 20 };
  const q = SwapWithCopy(p);
  console.log(p.First, p.Second);
  console.log(q.First, q.Second);
  SwapWithRef(p);
  console.log(p.First, p.Second);
})();
```

**TypeScript → Go (Rückübersetzung):**

```go
type Pair[T any] struct {
	First  T
	Second T
}
func SwapWithCopy[T any](p Pair[T]) Pair[T] {
	return Pair[T]{First: p.Second, Second: p.First}
}

func SwapWithRef[T any](p *Pair[T]) {
	temp := p.First
	p.First = p.Second
	p.Second = temp
}

func main() {
	p := Pair[int]{First: 10, Second: 20}
	q := SwapWithCopy(p)
	fmt.Println(p.First, p.Second);
	fmt.Println(q.First, q.Second)
	SwapWithRef(&p)
	fmt.Println(p.First, p.Second)
}
```

**Beobachtungen**:

Insgesamt funktioniert die Übersetzung in diesem Beispiel überraschend gut, rein semantisch ist die Rückübersetzung fast äquivalent zum Original:

1. Das LLM hat das Kernproblem verstanden und `SwapWithCopy` korrekt übersetzt: Statt das Objekt zu mutieren, gibt es ein neues Objekt zurück. Die Value-Semantik von Go wird dadurch in TypeScript bewahrt. Bei der Rückübersetzung erzeugt die Version des LLM allerdings ein weiteres Objekt, was zwar nicht notwendig ist, das Ergebnis allerdings auch nicht verändert.
2. Die Rückübersetzung rekonstruiert sowohl den Value-Parameter als auch den Pointer-Parameter korrekt. Das LLM könnte aus dem Muster und aus den Methodennamen erkannt haben, welche Go-Semantik gemeint ist.

---

#### Klassen (Struct-Embedding)

Dieses Beispiel bettet ein nicht-generisches Struct (`Named`) in ein generisches Struct (`Box[T]`) ein. Dieses Kompositionsmuster ist in TypeScript nicht direkt umsetzbar, sodass bei der Übersetzung zwischen Vererbung und Komposition entschieden werden muss.

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

**Vermutung:**

Die Wahl zwischen Vererbung und Komposition für die Darstellung des Embeddings stellt das LLM vor eine Herausforderung. Beides hat Nachteile: Vererbung ändert die Hierarchie und Komposition verliert den direkten Feldzugriff mit `b.Name`. Bei der Rückübersetzung könnte es zu Problemen kommen, gerade wenn das LLM die Strategie der Komposition wählt.

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


**Beobachtungen:**

Das LLM wählt Vererbung statt Komposition, um den direkten Feldzugriff `b.Name` zu erhalten.
Bei der Rückübersetzung zu Go kommt es dementsprechend nicht zu Problemen, da das LLM das Struct-Embedding auch im generischen Fall aus der Vererbung ableiten kann.

---

#### Methoden mit eigenen Typparametern

In Go ist es nicht möglich, dass Methoden mit generischer Signatur weitere generische Parameter einführen.
In TypeScript ist dieses Verhalten möglich.
Dieses Beispiel untersucht die Übersetzung und Rekonstruktion des dafür in Go genutzten Workarounds mit einer freistehenden Funktion.


```go
type Container[T any] struct {
	items []T
}

func Map[T any, U any](c Container[T], f func(T) U) Container[U] {
	result := Container[U]{items: make([]U, len(c.items))}
	for i, v := range c.items {
		result.items[i] = f(v)
	}
	return result
}

func main() {
	nums := Container[int]{items: []int{1, 2, 3}}

	strs := Map(nums, func(n int) string {
		return fmt.Sprintf("%d!", n)
	})
	fmt.Println(strs.items) // [1! 2! 3!]
}
```

**Vermutung:**

1. Vermutlich wird das LLM versuchen den Workaround zu übersetzen, obwohl TypeScript eine Methode mit eigenen Typparametern unterstützen würde.
2. Da TypeScript die Methode `map<U>` direkt ausdrücken kann, könnte das LLM (falls es korrekt nach TypeScript übersetzt hat) beim Zurückübersetzen versuchen, `func (c Container[T]) Map[U any](...)` zu schreiben, was ungültiger Go-Code wäre.

#### Ergebnisse:

**Go → TypeScript:**

```typescript
class Container<T> {
  public items: T[];
  constructor(items: T[] = []) {
    this.items = items;
  }
}

function Map<T, U>(c: Container<T>, f: (value: T) => U): Container<U> {
  const resultItems: U[] = new Array<U>(c.items.length);
  const result = new Container<U>(resultItems);
	.
  for (let i = 0; i < c.items.length; i++) {
    result.items[i] = f(c.items[i]);
  }

  return result;
}

function mainGoEquivalent(): void {
  const nums = new Container<number>([1, 2, 3]);
  console.log("Original numbers container:", nums.items); // Ausgabe: [ 1, 2, 3 ]

  const strs = Map(nums, (n: number) => `${n}!`);
  console.log("Mapped strings container:", strs.items); // Erwartete Ausgabe: [ '1!', '2!', '3!' ]
}

mainGoEquivalent();
```

**TypeScript → Go (Rückübersetzung):**

```go
type Container[T any] struct {
	Items []T
}

func NewContainer[T any](items []T) Container[T] {
	if items == nil {
		return Container[T]{Items: []T{}}
	}
	return Container[T]{Items: items}
}

func Map[T, U any](c Container[T], f func(value T) U) Container[U] {
	resultItems := make([]U, len(c.Items))
	result := NewContainer(resultItems) // Verwende NewContainer für Konsistenz

	for i := 0; i < len(c.Items); i++ {
		result.Items[i] = f(c.Items[i])
	}

	return result
}

func mainGoEquivalent() {
	nums := NewContainer([]int{1, 2, 3})
	fmt.Println("Original numbers container:", nums.Items) // Ausgabe: [1 2 3]

	strs := Map(nums, func(n int) string {
		return fmt.Sprintf("%d!", n)
	})
	fmt.Println("Mapped strings container:", strs.Items) // Erwartete Ausgabe: [1! 2! 3!]
}

func main() {
	mainGoEquivalent()
}
```

**Beobachtungen:**

In diesem Beispiel wurde die Syntax durch die Übersetzunge teils stark verändert, allerdings bleibt die Übersetzung gerade bei den generischen Komponenten sehr konsistent.

1. Das LLM übersetzt `Map` als freistehende generische Funktion und es versucht nicht, eine Methode `container.map<U>(...)` daraus zu machen, obwohl TypeScript dies unterstützen würde. Die Struktur des Originals bleibt erhalten.
2. Da TypeScript einen Constructor hat, erzeugt die Rückübersetzung eine `NewContainer`-Funktion, die im Original nicht existiert. Das Original initialisiert den Container direkt als Struct-Literal.
3. `Map[T, U any]` mit zwei Typparametern und der Funktionsparameter `f func(T) U` werden in beiden Richtungen korrekt übersetzt.

---

### Zusammenfassung der Ergebnisse

| Testfall                   | Bijektiv? | Generics-Übersetzung                                                                                          |
| -------------------------- | --------- | ------------------------------------------------------------------------------------------------------------- |
| Generische Primitive       | ✅ Ja      | `[T any]` ↔ `<T>` und `[]T` ↔ `T[]` verlustfrei übersetzt                                               |
| Type Sets                  | ❌ Nein    | Type Set  verliert `~` und Operator-Constraint; TS kompiliert nicht |
| Pointer/Call-by-value      | 🟡 Fast    | `Pair[T]` und `*Pair[T]` korrekt als `Pair<T>` mit unterschiedlicher Semantik übersetzt                    |
| Struct-Embedding           | ✅ Ja      | `Box[T]` mit Embedding → `Box<T> extends Named`; Generics vollständig rekonstruiert                         |
| Methoden mit Typparametern | 🟡 Fast    | `Map[T, U any]` mit `func(T) U` in beiden Richtungen korrekt übersetzt; Workaround bleibt erhalten          |

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

Die Übersetzung zwischen Go und TypeScript Generics ist **nicht bijektiv** im strengen Sinne. Die Experimente zeigen jedoch ein differenziertes Bild:

**Generische Typparameter selbst sind robust übersetzbar:**

- `[T any]` ↔ `<T>`, `[T, U any]` ↔ `<T, U>` und generische Structs wie `Pair[T]` ↔ `Pair<T>` werden in allen fünf Tests korrekt hin- und zurückübersetzt.
- Generische Primitive (`[]T` ↔ `T[]`) und Funktionsparameter (`func(T) U` ↔ `(value: T) => U`) werden verlustfrei abgebildet.
- Auch komplexere Muster wie Multi-Parameter-Generics (`Map[T, U any]`) und generische Konstruktorfunktionen (`NewBox[T any]`) bleiben strukturell erhalten.

**Probleme entstehen bei den Constraints und dem Typsystem, nicht bei den Typparametern:**

- **Type Sets mit Operatoren** (Test 2): Das Interface-Constraint `~int | ~float64 | ~string` wird zu `number | string` vereinfacht. Der Tilde-Operator (`~`) und die damit verbundene nominale Typunterscheidung gehen verloren. Zudem kann TypeScript keine Operator-Unterstützung in Constraints ausdrücken, weshalb `a + b` einen Kompilierfehler erzeugt. Bei der Rückübersetzung halluziniert das LLM zusätzliche Typen (`~float32`) und ändert den Basistyp von `Meter`.
- **Zeiger-Semantik** (Test 3): Obwohl `*Pair[T]` kein TypeScript-Äquivalent hat, erkennt das LLM das Muster und rekonstruiert die Pointer-/Value-Unterscheidung korrekt. Dies gelingt jedoch vermutlich nur durch die sprechenden Funktionsnamen (`SwapWithCopy` vs. `SwapWithRef`).
- **Workaround-Erhaltung** (Test 5): Go's Einschränkung, dass Methoden keine eigenen Typparameter haben dürfen, führt zum Workaround mit freistehenden Funktionen. Das LLM erhält diesen Workaround bei der Übersetzung, obwohl TypeScript eine Methode `map<U>(...)` ermöglichen würde.

Die generischen Typparameter und ihre grundlegende Syntax sind zwischen Go und TypeScript nahezu bijektiv übersetzbar. Die Verluste entstehen dort, wo Go's Constraint-System (Type Sets, `~`, Operator-Garantien) über das hinausgeht, was TypeScript's `extends`-Constraints ausdrücken können. Außerdem zeigt sich, dass durch Limitationen von Go die Möglichkeiten von TypeScript teilweise nicht vollständig ausgeschöpft werden. Die Abbildung ist somit bijektiv für die Typparameter selbst, aber nicht für das umgebende Typsystem.
