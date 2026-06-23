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

Für die folgenden Testbeispiele wird jeweils zuerst der Go-Quellcode gezeigt, dann eine Vermutung formuliert, welche Probleme bei der Übersetzung auftreten könnten.
Anschließend wird das LLM mit dem unten definierten Prompt zur Übersetzung und Rückübersetzung aufgefordert.
Zwischen jeder Anfrage wird der Kontext zurückgesetzt, damit das Modell sich nicht an vergangenen Nachrichten orientieren kann.
Außerdem werden Kommentare entfernt, welche auf die ursprüngliche Version hinweisen.

---

#### Generische Primitive

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
`any` entspricht direkt TypeScript's generischem `<T>` ohne Constraint.
`make([]T, n)` wird zu `new Array<T>(n)` oder einem Array-Literal.
Die Hin- und Rückübersetzung sollte semantisch identisch sein.
Lediglich `fmt.Println` vs. `console.log` und die Slice-Erzeugung unterscheiden sich syntaktisch.

#### Type Sets vs. Type Union

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
Go erlaubt `a + b` nur, weil das Type Set garantiert, dass alle enthaltenen Typen `+` unterstützen.
TypeScript hat keine Möglichkeit, Operator-Unterstützung in einem Constraint auszudrücken.
Das LLM wird vermutlich `T extends number | string` schreiben, aber TypeScript erlaubt `a + b` nicht generisch. 
Es wird wahrscheinlich eine Type Assertion (`as any`) oder Overloads benötigen.
2. **Tilde-Operator (`~`):** 
Go's `~float64` erlaubt auch benannte Typen wie `Meter`. TypeScript hat kein Äquivalent und `type Meter = number` erzeugt keinen eigenständigen Typ, sondern nur ein Alias.
Der Unterschied zwischen nominaler (`~`) und struktureller Typisierung geht verloren.
3. **Rückübersetzung:** Aus TypeScript zurück nach Go wird der Tilde-Operator wahrscheinlich fehlen, und die `Meter`-Nutzung geht verloren.

#### Call-by-Value & Pointer vs. Call-by-Reference

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

**Vermutung:** Struct-Embedding hat kein direktes TypeScript-Äquivalent:
1. **Embedding vs. Vererbung:** Das LLM muss zwischen `class Box<T> extends Named` (Vererbung) und Komposition (`name: Named`) wählen. Beides hat Nachteile – Vererbung ändert die Hierarchie, Komposition verliert den direkten Feldzugriff (`b.Name`).
2. **Konstruktor:** Go's `NewBox`-Funktion wird wahrscheinlich zu einem `constructor` in einer Klasse. Die Rückübersetzung sollte das zu einer `NewBox`-Funktion rekonstruieren.
3. **Direkter Feldzugriff:** In Go kann man `b.Name` schreiben (promoted field). Bei Komposition in TypeScript wäre es `b.named.name`. Das LLM wird vermutlich Vererbung wählen, um diesen Zugriff zu erhalten.

#### Type Switch (if type == bool)

```go
type Dog struct{ Name string }

func (d Dog) String() string { return d.Name }

func Describe(v any) string {
	switch val := v.(type) {
	case bool:
		if val { return "YES" }
		return "NO"
	case int:
		return fmt.Sprintf("#%d", val)
	case Dog:
		return "Dog: " + val.Name
	default:
		return fmt.Sprint(v)
	}
}

func main() {
	fmt.Println(Describe(true))             // YES
	fmt.Println(Describe(42))               // #42
	fmt.Println(Describe(Dog{Name: "Rex"})) // Dog: Rex
	fmt.Println(Describe("hello"))          // hello
}
```

**Vermutung:** Der Type Switch ist ein interessanter Grenzfall:
1. **Primitive Typen (`bool`, `int`):** Für diese kann das LLM `typeof v === "boolean"` verwenden – das funktioniert.
2. **Struct-Typ (`Dog`):** Hier versagt `typeof`, da es nur `"object"` liefert. Das LLM muss `instanceof` (erfordert Klassen) oder ein Discriminator-Feld einführen.
3. **Nominale Unterscheidung:**
In Go sind Structs nominal verschieden. Zwei Structs mit gleichen Feldern sind verschiedene Typen. In TypeScript sind strukturell identische Typen austauschbar; nur `instanceof` unterscheidet sie zur Laufzeit.
4. **Rückübersetzung:** `typeof`-Checks für Primitive werden korrekt rekonstruiert. Aber `instanceof Dog` wird wahrscheinlich nicht zum originalen Struct-Pattern zurückübersetzt, sondern eine Klasse beibehalten.

### Ergebnisse (Gemini Flash 2.5)

#### Test 4: Generische Primitive

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
1. Wie vermutet ist die Übersetzung nahezu bijektiv. Die Semantik ist vollständig erhalten.
2. Syntaktischer Unterschied: `for i := range result` (Original) → `for i := 0; i < n; i++` (Rückübersetzung). Funktional identisch.

---

#### Type Sets vs. Type Union

**Go → TypeScript:**

```typescript
type Addable = number | string;

function Add<T extends Addable>(a: T, b: T): T {
  return a + b; // throws error
}

type Meter = number;

function main() {
  console.log(Add(3, 4));         // 7
  console.log(Add("Go", "Lang")); // GoLang

  const d1: Meter = 10;
  const d2: Meter = 20;
  console.log(Add(d1, d2));       // 30
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
1. **TypeScript-Code kompiliert nicht:** `return a + b` mit `T extends number | string` erzeugt einen TypeScript-Fehler.
2. **Tilde in Rückübersetzung vorhanden:** Überraschenderweise fügt das LLM bei der Rückübersetzung `~` hinzu (`~int | ~float32 | ~float64 | ~string`), obwohl TypeScript dieses Konzept nicht hat. Das LLM hat hier offenbar Go-Idiome angewandt.
3. **`Meter` wird zu `int` statt `float64`:** Im Original ist `type Meter float64`, die Rückübersetzung erzeugt `type Meter int`. TypeScript's `number` ist nicht eindeutig auf `int` oder `float64` abbildbar.
4. **Zusätzlicher Typ `float32`:** Die Rückübersetzung fügt `~float32` zum Constraint hinzu, das im Original nicht existiert.
5. **Initialisierung geändert:** `var d1, d2 Meter = 10, 20` (Original) → `d1 := Meter(10)` (Rückübersetzung). Funktional äquivalent, aber syntaktisch verschieden.

---

#### Call-by-Value & Pointer

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
1. **Vererbung gewählt:** Wie vermutet wählt das LLM `extends Named` statt Komposition, um den direkten Feldzugriff `b.Name` zu erhalten.
2. **Rückübersetzung nahezu identisch:** Das Embedding wird korrekt rekonstruiert. `Named` wird wieder eingebettet, `NewBox` bleibt eine Funktion. Dies ist das einzige Beispiel mit fast perfekter Bijektivität.

---

#### Type Switch

**Go → TypeScript:**

```typescript
class Dog {
    constructor(public Name: string) {}
    toString(): string {
        return this.Name;
    }
}

function Describe(v: unknown): string {
    if (typeof v === 'boolean') {
        if (v) { return "YES"; }
        return "NO";
    }
    if (typeof v === 'number' && Number.isInteger(v)) {
        return `#${v}`;
    }
    if (v instanceof Dog) {
        return `Dog: ${v.Name}`;
    }
    return String(v);
}

function main() {
    console.log(Describe(true));
    console.log(Describe(42));
    console.log(Describe(new Dog("Rex")));
    console.log(Describe("hello"));
}

main();
```

**TypeScript → Go (Rückübersetzung):**

```go
type Dog struct {
	Name string
}

func NewDog(name string) *Dog {
	return &Dog{Name: name}
}

func (d *Dog) String() string {
	return d.Name
}

func Describe(v any) string {
	switch val := v.(type) {
	case bool:
		if val { return "YES" }
		return "NO"
	case int:
		return fmt.Sprintf("#%d", val)
	case int8:
		return fmt.Sprintf("#%d", val)
	case int16:
		return fmt.Sprintf("#%d", val)
	// ... (alle Integer-Typen)
	case float64:
		if val == math.Trunc(val) {
			return fmt.Sprintf("#%d", int64(val))
		}
	case *Dog:
		return fmt.Sprintf("Dog: %s", val.Name)
	default:
		return fmt.Sprint(v)
	}
}

func main() {
	fmt.Println(Describe(true))
	fmt.Println(Describe(42))
	fmt.Println(Describe(NewDog("Rex")))
	fmt.Println(Describe("hello"))
}
```

**Beobachtungen:**
1. **`typeof` für Primitive:** Wie vermutet verwendet das LLM `typeof v === "boolean"` und `typeof v === "number"`.
2. **`instanceof` für Struct:** Wie vermutet wird `instanceof Dog` für den Struct-Typ verwendet.
3. **`case *Dog` statt `case Dog`:** Die Rückübersetzung erzeugt `case *Dog` (Pointer) und `NewDog` gibt `*Dog` zurück. Das Original verwendet `case Dog` (Value-Typ) und erzeugt Dogs als `Dog{Name: "Rex"}`. Dies ändert die Aufruf-Semantik.
4. **Explosion der Integer-Cases:** TypeScript hat nur `number`, Go hat `int`, `int8`, `int16`, ..., `float32`, `float64`. Das LLM erzeugt separate Cases für jeden Integer-Typ – eine massive Aufblähung gegenüber dem Original (`case int`).
5. **`Number.isInteger`-Logik:** Die Hinübersetzung fügt `Number.isInteger(v)` hinzu, was in der Rückübersetzung zu einer komplexen Float-Prüfung mit `math.Trunc` führt – Logik, die im Original gar nicht existiert.
6. **`String()` → `toString()`:** Go's `String()`-Methode wird zu TypeScript's `toString()` – idiomatisch korrekt, aber nicht bijektiv.
7. **`unknown` statt `any`:** Das LLM wählt `unknown` statt `any` als TypeScript-Typ für Go's `any`, was strenger ist und Type Guards erfordert.

---

### Zusammenfassung der Ergebnisse

| Testfall | Bijektiv? | Hauptproblem |
|----------|-----------|--------------|
| Generische Primitive | ✅ Ja | Nur syntaktische Unterschiede (Loop-Stil) |
| Type Sets | ❌ Nein | TS-Code kompiliert nicht (`a + b`), `Meter` wird zu `int` statt `float64` |
| Pointer/Call-by-value | ❌ Nein | Wrapper-Struct `Ptr[T]` statt Zeiger, externe Abhängigkeit eingefügt |
| Struct-Embedding | ✅ Fast | Nahezu identisch, nur Default-Parameter-Unterschied |
| Type Switch | ❌ Nein | `*Dog` statt `Dog`, Integer-Explosion, `Number.isInteger`-Logik hinzugefügt |

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
- Einfache generische Funktionen mit `any`-Constraint sind nahezu perfekt übersetzbar (Test 4).
- Struct-Embedding ↔ Klassen-Vererbung funktioniert gut, solange keine Pointer-/Value-Semantik involviert ist (Test 7).
- Die grundlegende Struktur (Typparameter, Funktionssignaturen, generische Container) wird in allen Fällen korrekt übertragen.

**Wann geht's schief?**
- **Type Sets mit Operatoren**: TypeScript kann Operator-Constraints nicht ausdrücken. Der generierte Code kompiliert teilweise nicht, und der Tilde-Operator (`~`) sowie benannte Typen gehen bei der Hinübersetzung verloren.
- **Zeiger-Semantik**: Da TypeScript keine Zeiger hat, entstehen Wrapper-Structs, die bei der Rückübersetzung nicht zu einfachen Zeigern rekonstruiert werden. Die Aufruf-Semantik ändert sich grundlegend.
- **Type Switches mit Struct-Typen**: Die Unterschiede zwischen Go's nominalem Typsystem und TypeScript's strukturellem Typsystem werden sichtbar. `typeof` funktioniert nur für Primitive; für Structs braucht es `instanceof`, was zu `*Dog` statt `Dog` führt. Zusätzlich erzeugt die Rückübersetzung eine Explosion an Integer-Cases, da TypeScript's `number` nicht eindeutig auf einen Go-Typ abbildbar ist.

**Allgemein:** Die Abbildung ist bijektiv, wenn der Code ausschließlich Konzepte nutzt, die in beiden Sprachen strukturelle Äquivalente haben: generische Funktionen ohne Operator-Nutzung, einfache Constraints, und Datenstrukturen ohne Zeiger-Semantik. Sobald sprachspezifische Features involviert sind (Type Sets mit `~`, Zeiger, nominale Typunterscheidung), ist die Rückübersetzung nicht mehr äquivalent zum Original.
