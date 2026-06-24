// type Pair[T any] struct {
//   First, Second T
// }
/**
 * Definiert eine generische Pair-Struktur mit zwei Feldern, First und Second,
 * beide vom Typ T.
 * Entspricht Go's `type Pair[T any] struct { First, Second T }`.
 */
interface Pair<T> {
  First: T;
  Second: T;
}

// func SwapWithCopy[T any](p Pair[T]) Pair[T] {
//   p.First, p.Second = p.Second, p.First
//   return p
// }
/**
 * Tauscht die Felder First und Second eines Paares und gibt ein *neues* Paar
 * mit den getauschten Werten zurück.
 * Das ursprüngliche als Argument übergebene Pair bleibt unverändert,
 * was Go's Pass-by-Value-Verhalten für Strukturen simuliert.
 * @param p Das zu tauschende Paar.
 * @returns Ein neues Paar mit getauschten First- und Second-Werten.
 */
function SwapWithCopy<T>(p: Pair<T>): Pair<T> {
  // In Go werden Strukturen by Value übergeben. Eine Modifikation von 'p'
  // innerhalb der Funktion würde nur eine lokale Kopie modifizieren.
  // Wir simulieren dies, indem wir ein *neues* Objekt mit den getauschten
  // Werten erstellen und zurückgeben, wodurch das ursprüngliche 'p' unberührt bleibt.
  return { First: p.Second, Second: p.First };
}

// func SwapWithRef[T any](p *Pair[T]) {
//   p.First, p.Second = p.Second, p.First
// }
/**
 * Tauscht die Felder First und Second eines Paares direkt (in-place).
 * Diese Funktion mutiert das als Argument übergebene ursprüngliche Pair-Objekt,
 * was Go's Pass-by-Pointer-Verhalten simuliert.
 * @param p Das zu mutierende Pair-Objekt.
 */
function SwapWithRef<T>(p: Pair<T>): void {
  // In TypeScript werden Objekte by Reference übergeben.
  // Das direkte Modifizieren von Eigenschaften von 'p' mutiert das ursprüngliche Objekt.
  const temp = p.First;
  p.First = p.Second;
  p.Second = temp;
}

// func main() {
/**
 * Hauptausführungsblock, der Go's `main` Funktion simuliert.
 */
(function main() {
  // p := Pair[int]{First: 10, Second: 20}
  const p: Pair<number> = { First: 10, Second: 20 };

  // q := SwapWithCopy(p)
  const q = SwapWithCopy(p);

  // fmt.Println(p.First, p.Second) // 10 20 (unverändert!)
  console.log(p.First, p.Second); // Erwartete Ausgabe: 10 20 (unverändert!)

  // fmt.Println(q.First, q.Second) // 20 10 (neue Kopie)
  console.log(q.First, q.Second); // Erwartete Ausgabe: 20 10 (neue Kopie)

  // SwapWithPlace(&p) // Anmerkung: Es wird angenommen, dass SwapWithRef hier gemeint ist, basierend auf der Funktionsdefinition.
  SwapWithRef(p);

  // fmt.Println(p.First, p.Second) // 20 10 (mutiert)
  console.log(p.First, p.Second); // Erwartete Ausgabe: 20 10 (mutiert)
})(); // Sofort aufgerufene Funktionsausdruck (IIFE) zur Simulation der Go-Hauptausführung.