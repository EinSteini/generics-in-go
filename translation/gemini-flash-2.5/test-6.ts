/**
 * Typ `Ptr<T>` ahmt das Zeigerverhalten von Go für veränderliche Werte nach.
 * Für primitive Typen wie `number` oder `string` müssen diese in `{ value: T }`
 * eingepackt werden, um eine In-Place-Modifikation zu ermöglichen.
 * Bei Objekten können Sie den Objektwert direkt an `value` zuweisen,
 * wenn dessen Eigenschaften mutiert werden sollen, oder das gesamte Objekt
 * als Wert in `Ptr<T>` verpacken, um es zu tauschen.
 */
type Ptr<T> = { value: T };

/**
 * Tauscht die Werte, auf die 'a' und 'b' verweisen.
 *
 * In Go funktioniert dies für jeden Typ T, indem Zeiger (`*T`) übergeben werden.
 * In TypeScript müssen für primitive Typen (number, string, boolean),
 * um den gleichen In-Place-Mutationseffekt zu erzielen, diese in ein Objekt
 * (`Ptr<T>`) eingepackt werden.
 * Bei nicht-primitiven Objekten (Arrays, benutzerdefinierte Objekte)
 * würden 'a' und 'b' bereits per Referenz übergeben.
 * Wenn das Ziel ist, die Objekte selbst zu tauschen, ist `Ptr<T>` erforderlich.
 *
 * Beispielverwendung mit Primitiven:
 * let x = { value: 10 };
 * let y = { value: 20 };
 * Swap(x, y);
 * console.log(x.value, y.value); // 20 10
 *
 * Beispielverwendung mit Objekten:
 * let obj1 = { value: { name: "Alice" } };
 * let obj2 = { value: { name: "Bob" } };
 * Swap(obj1, obj2);
 * console.log(obj1.value, obj2.value); // { name: "Bob" } { name: "Alice" }
 */
function Swap<T>(a: Ptr<T>, b: Ptr<T>): void {
  const temp = a.value;
  a.value = b.value;
  b.value = temp;
}

/**
 * Verdoppelt den Eingabewert.
 *
 * In Go ist der Typparameter T auf zugrunde liegende Integer- (`~int`)
 * oder Float64-Typen (`~float64`) beschränkt.
 * In TypeScript deckt `number` sowohl Integer- als auch Gleitkommazahlen ab.
 * Die `extends number`-Constraint stellt sicher, dass Multiplikation eine
 * gültige Operation für Typ T ist. Der Rückgabetyp wird zu T gecastet,
 * um die ursprüngliche Typspezifität zu erhalten, insbesondere wenn T
 * ein numerischer Literal-Typ ist.
 *
 * @param v Der zu verdoppelnde numerische Wert.
 * @returns Der verdoppelte Wert, wobei der ursprüngliche Typ `T` beibehalten wird.
 */
function Double<T extends number>(v: T): T {
  return (v * 2) as T; // Original bleibt unverändert
}

// Emulation der Go-`main`-Funktion
function main() {
  // Go: x, y := 10, 20
  // Um Swap mit Primitiven zu verwenden, müssen wir sie in Ptr-Objekte einpacken.
  let x: Ptr<number> = { value: 10 };
  let y: Ptr<number> = { value: 20 };

  // Go: Swap(&x, &y)
  Swap(x, y);

  // Go: fmt.Println(x, y) // 20 10
  console.log(x.value, y.value); // Erwartet: 20 10

  // Go: n := 5
  let n = 5;

  // Go: fmt.Println(Double(n), n) // 10 5
  console.log(Double(n), n); // Erwartet: 10 5
}

// Ruft die main-Funktion auf, wenn das Skript ausgeführt wird
main();