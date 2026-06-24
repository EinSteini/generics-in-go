
type Addable = number | string;

/**
 * Adds two values of the same type `T`.
 * `T` must be a type whose underlying type is `number` or `string`.
 * @template T - The type of the values, constrained to `Addable`.
 * @param {T} a - The first value.
 * @param {T} b - The second value.
 * @returns {T} The result of the addition/concatenation.
 */
function Add<T extends Addable>(a: T, b: T): T {
  // TypeScript's '+' operator handles both numeric addition and string concatenation.
  // The type inference for `T` and the return type will correctly match the input types.
  return a + b;
}

type Meter = number;

function main() {
  console.log(Add(3, 4));         // 7
  console.log(Add("Go", "Lang")); // GoLang

  const d1: Meter = 10;
  const d2: Meter = 20;
  console.log(Add(d1, d2));       // 30 (benannter Typ dank ~)
}

// Call the main function to execute the example
main();