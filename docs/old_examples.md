
#### Snippet 1: Generische Funktion mit Type-Set-Constraint (`translation/test_1_generic_function.go`)

```go
type Number interface {
	~int | ~float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Map[T any, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		result[i] = fn(v)
	}
	return result
}
```

**Fragen:**
- Wird `~int | ~float64` korrekt als TypeScript Union-Type abgebildet und zurückübersetzt?
- Was passiert mit `make(..., size)`?

##### Claude 4.5 Haiku (gekürzt, Original in translation/4.5-haiku)
**TypeScript:**
```ts
type Number = number;

function Min<T extends Number>(a: T, b: T): T {
  if (a < b) {
    return a;
  }
  return b;
}

function Map<T, U>(items: T[], fn: (item: T) => U): U[] {
  const result: U[] = [];
  for (let i = 0; i < items.length; i++) {
    result[i] = fn(items[i]);
  }
  return result;
}
```
- kein union Type für Number, stattdessen ts-Typ `number`

**Go:**
```go
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

func Min[T Number](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func Map[T, U any](items []T, fn func(item T) U) []U {
	result := make([]U, len(items))
	for i := 0; i < len(items); i++ {
		result[i] = fn(items[i])
	}
	return result
}
```

**Beobachtungen:**
- Number type deutlich umfangreicher
- Go `make(..., size)` wieder rekonstruiert
- for schleife über index anstatt for-each

#### Snippet 2: Generischer Struct mit Methoden (`translation/test_2_generic_struct.go`)

```go
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
```

**Fragen:**
- Wird `var zero T` nach der Rückübersetzung wieder rekonstruiert?
- Wird der Pointer-Receiver `*Stack[T]` wiederhergestellt?
- Bleibt das Tuple-Return `(T, bool)` erhalten?

---

##### Claude 4.5 Haiku (gekürzt, Original in tranlation/4.5-haiku)
**TypeScript:**
```ts
class Stack<T> {
	private items: T[] = [];

	push(v: T): void {
		this.items.push(v);
	}

	pop(): [T | undefined, boolean] {
		if (this.items.length === 0) {
			return [undefined, false];
		}
		const top = this.items[this.items.length - 1];
		this.items = this.items.slice(0, this.items.length - 1);
		return [top, true];
	}

	len(): number {
		return this.items.length;
	}
}
```
**Go:**
```go
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
```
**Beobachtungen:**
- LLM schafft Rückübersetzung erfolgreich zu semantisch identischem Code


#### Snippet 3: Interface-Embedding als Constraint (`translation/test_3_interface_constraint.go`)
```go
type Stringer interface {
	String() string
}

type Validator interface {
	IsValid() bool
}

type Entry interface {
	Stringer
	Validator
}

func PrintValid[T Entry](items []T) {
	for _, item := range items {
		if item.IsValid() {
			fmt.Println("✓", item.String())
		} else {
			fmt.Println("✗", item.String())
		}
	}
}

type Email struct {
	Address string
}

func (e Email) String() string { return e.Address }
func (e Email) IsValid() bool  { return len(e.Address) > 3 }
```

**Fragen:**
- Wird Interface-Embedding (`Entry` enthält `Stringer` + `Validator`) korrekt durch `extends` abgebildet und zurückübersetzt?
- Werden Value-Receiver bei der Rückübersetzung beibehalten oder zu Pointer-Receivern?
- Wird die implizite Interface-Implementierung (Go) nach dem Umweg über `implements` (TS) wieder implizit?

##### Claude 4.5 Haiku (gekürzt, Original in translation/4.5-haiku)
**TypeScript:**
```ts
interface Stringer {
	String(): string;
}

interface Validator {
	IsValid(): boolean;
}

interface Entry extends Stringer, Validator {}

function PrintValid<T extends Entry>(items: T[]): void {
	for (const item of items) {
		if (item.IsValid()) {
			console.log("✓", item.String());
		} else {
			console.log("✗", item.String());
		}
	}
}

class Email implements Entry {
	Address: string;

	constructor(address: string) {
		this.Address = address;
	}

	String(): string {
		return this.Address;
	}

	IsValid(): boolean {
		return this.Address.length > 3;
	}
}
```
**Go:**
```go
type Stringer interface {
	String() string
}

type Validator interface {
	IsValid() bool
}

type Entry interface {
	Stringer
	Validator
}

func PrintValid[T Entry](items []T) {
	for _, item := range items {
		if item.IsValid() {
			fmt.Println("✓", item.String())
		} else {
			fmt.Println("✗", item.String())
		}
	}
}

type Email struct {
	Address string
}

func NewEmail(address string) *Email {
	return &Email{
		Address: address,
	}
}

func (e *Email) String() string {
	return e.Address
}

func (e *Email) IsValid() bool {
	return len(e.Address) > 3
}
```
**Beobachtungen:**
- Entry wurde korrekt rekonstruiert
- Wir haben Pointer-Receiver anstatt von Value-Receivern
- Constraint by PrintValid rekonstruiert
