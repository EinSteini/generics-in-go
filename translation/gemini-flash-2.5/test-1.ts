/**
 * Repeats a given value `v` `n` times and returns an array containing these values.
 * This is a generic function, meaning it can work with any type `T`.
 *
 * @template T The type of the value to repeat.
 * @param v The value to repeat.
 * @param n The number of times to repeat the value.
 * @returns An array of type `T[]` containing `n` copies of `v`.
 */
export function Repeat<T>(v: T, n: number): T[] {
  // Erstelle ein Array vom Typ T mit der Länge n.
  // In TypeScript erstellt `new Array(n)` ein Array mit 'n' undefinierten Elementen.
  const result: T[] = new Array(n);

  // Iteriere über die Indizes des Arrays.
  for (let i = 0; i < n; i++) {
    // Weise den bereitgestellten Wert 'v' jedem Element zu.
    result[i] = v;
  }
  return result;
}

// Dies ist das Äquivalent der `main`-Funktion in Go.
// In TypeScript werden Skripte typischerweise von oben nach unten ausgeführt.

// Beispielverwendung: Wiederhole den String "go" 3 Mal.
console.log(Repeat("go", 3)); // Erwartete Ausgabe: [ 'go', 'go', 'go' ]
// Beispielverwendung: Wiederhole die Zahl 42 2 Mal.
console.log(Repeat(42, 2));   // Erwartete Ausgabe: [ 42, 42 ]