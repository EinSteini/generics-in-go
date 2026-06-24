// type Container[T any] struct {
//   items []T
// }

/**
 * Container ist eine generische Struktur (in TypeScript eine Klasse), die
 * ein Slice (in TypeScript ein Array) von Elementen des Typs T enthält.
 *
 * @template T Der Typ der im Container gespeicherten Elemente.
 */
class Container<T> {
  // items []T
  public items: T[];

  /**
   * Erstellt eine neue Container-Instanz.
   * @param items Ein optionales Array von Elementen, um den Container zu initialisieren.
   *              Standardmäßig ein leeres Array.
   */
  constructor(items: T[] = []) {
    this.items = items;
  }
}

// func Map[T any, U any](c Container[T], f func(T) U) Container[U] {
//   result := Container[U]{items: make([]U, len(c.items))}
//   for i, v := range c.items {
//     result.items[i] = f(v)
//   }
//   return result
// }

/**
 * Map transformiert einen Container vom Typ T in einen Container vom Typ U,
 * indem eine Funktion `f` auf jedes Element angewendet wird.
 *
 * Diese Funktion ist generisch und funktioniert mit beliebigen Typen T und U.
 *
 * @template T Der Typ der Eingabeelemente.
 * @template U Der Typ der Ausgabeelemente nach der Transformation.
 * @param c Der Eingabe-Container vom Typ T.
 * @param f Die Mapping-Funktion, die ein Element vom Typ T entgegennimmt
 *          und ein Element vom Typ U zurückgibt.
 * @returns Einen neuen Container vom Typ U mit den transformierten Elementen.
 */
function Map<T, U>(c: Container<T>, f: (value: T) => U): Container<U> {
  // In Go: result := Container[U]{items: make([]U, len(c.items))}
  // In TypeScript: Wir erstellen zuerst ein Array der benötigten Länge.
  // Beachte: `new Array(length)` erstellt ein Array mit 'leeren' Slots.
  // Die Go-Funktion `make` initialisiert mit Nullwerten, die jedoch
  // im Loop sofort überschrieben werden. Dieser Ansatz ist daher
  // semantisch äquivalent für das Endergebnis.
  const resultItems: U[] = new Array<U>(c.items.length);
  const result = new Container<U>(resultItems);

  // In Go: for i, v := range c.items { result.items[i] = f(v) }
  // In TypeScript: Iteriere durch die Eingabeelemente und wende die Funktion an.
  for (let i = 0; i < c.items.length; i++) {
    result.items[i] = f(c.items[i]);
  }

  return result;
}

// func main() {
//   nums := Container[int]{items: []int{1, 2, 3}}
//
//   strs := Map(nums, func(n int) string {
//     return fmt.Sprintf("%d!", n)
//   })
//   fmt.Println(strs.items) // [1! 2! 3!]
// }

// Beispielanwendung, die den Container und die Map-Funktion demonstriert,
// äquivalent zur `main`-Funktion in Go.
function mainGoEquivalent(): void {
  // Go: nums := Container[int]{items: []int{1, 2, 3}}
  // TypeScript: Instanziierung des Containers mit Zahlen.
  const nums = new Container<number>([1, 2, 3]);
  console.log("Original numbers container:", nums.items); // Ausgabe: [ 1, 2, 3 ]

  // Go: strs := Map(nums, func(n int) string { return fmt.Sprintf("%d!", n) })
  // TypeScript: Aufruf von Map mit einer anonymen Funktion (Pfeilfunktion),
  // um Zahlen in Strings umzuwandeln. fmt.Sprintf("%d!", n) ist äquivalent zu `${n}!` in TypeScript.
  const strs = Map(nums, (n: number) => `${n}!`);

  // Go: fmt.Println(strs.items) // [1! 2! 3!]
  // TypeScript: Ausgabe der Elemente des neuen Containers.
  console.log("Mapped strings container:", strs.items); // Erwartete Ausgabe: [ '1!', '2!', '3!' ]
}

// Führt die äquivalente Go-Main-Funktion aus.
mainGoEquivalent();