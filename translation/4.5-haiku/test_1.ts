// Type Constraint für numerische Typen
type Number = number;

// Generische Funktion Min mit Constraint
function Min<T extends Number>(a: T, b: T): T {
  if (a < b) {
    return a;
  }
  return b;
}

// Generische Map-Funktion
function Map<T, U>(items: T[], fn: (item: T) => U): U[] {
  const result: U[] = [];
  for (let i = 0; i < items.length; i++) {
    result[i] = fn(items[i]);
  }
  return result;
}

// Hauptprogramm
function main(): void {
  console.log(Min(3, 7));
  console.log(Min(2.5, 1.8));

  const words: string[] = ["hello", "world"];
  const lengths: number[] = Map(words, (s: string): number => s.length);
  console.log(lengths);
}

main();