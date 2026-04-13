export {};

// WITHOUT generics: separate types, duplicated code

// IntStack only works for number.
class IntStack {
  private items: number[] = [];

  push(v: number): void {
    this.items.push(v);
  }
  len(): number {
    return this.items.length;
  }
  pop(): [number, boolean] {
    if (this.items.length === 0) return [0, false];
    return [this.items.pop()!, true];
  }
}

// StringStack is an identical copy — just with string instead of number.
class StringStack {
  private items: string[] = [];

  push(v: string): void {
    this.items.push(v);
  }
  len(): number {
    return this.items.length;
  }
  pop(): [string, boolean] {
    if (this.items.length === 0) return ["", false];
    return [this.items.pop()!, true];
  }
}

// AnyStack avoids the duplication but loses type safety:
// compiler cannot catch push(42) on a "string stack" — only a runtime error can.
class AnyStack {
  private items: any[] = [];

  push(v: any): void {
    this.items.push(v);
  }
  len(): number {
    return this.items.length;
  }
  pop(): [any, boolean] {
    if (this.items.length === 0) return [undefined, false];
    return [this.items.pop(), true];
  }
}

// WITH GENERICS: one implementation, full type safety

class Stack<T> {
  private items: T[] = [];

  push(v: T): void {
    this.items.push(v);
  }
  len(): number {
    return this.items.length;
  }
  pop(): [T | undefined, boolean] {
    if (this.items.length === 0) return [undefined, false];
    return [this.items.pop()!, true];
  }
}

function main() {
  console.log("=== Without generics ===");

  const intStack = new IntStack();
  intStack.push(1);
  intStack.push(2);
  intStack.push(3);
  const [intVal, intOk] = intStack.pop();
  if (intOk) console.log("IntStack popped:", intVal);

  const stringStack = new StringStack();
  stringStack.push("go");
  stringStack.push("generics");
  const [strVal, strOk] = stringStack.pop();
  if (strOk) console.log("StringStack popped:", strVal);

  // AnyStack: both pushes compile fine, error only surfaces at runtime.
  const anyStack = new AnyStack();
  anyStack.push(42);
  anyStack.push("oops");
  const [anyVal, anyOk] = anyStack.pop();
  if (anyOk) {
    const s = anyVal as string; // runtime cast, throws if wrong
    console.log("AnyStack top was a string:", s);

    // This would throw at runtime:
    // const n = anyVal as number;
  }

  console.log("\n=== With generics ===");

  const genericIntStack = new Stack<number>();
  genericIntStack.push(10);
  genericIntStack.push(20);
  genericIntStack.push(30);

  // This would not compile: Argument of type 'string' is not assignable to parameter of type 'number'
  // genericIntStack.push("hello");

  const [gIntVal, gIntOk] = genericIntStack.pop();
  if (gIntOk) console.log("Stack<number> popped:", gIntVal);

  const genericStringStack = new Stack<string>();
  genericStringStack.push("go");
  genericStringStack.push("generics");

  // This would not compile: Argument of type 'number' is not assignable to parameter of type 'string'
  // genericStringStack.push(123);
  
  const [gStrVal, gStrOk] = genericStringStack.pop();
  if (gStrOk) console.log("Stack<string> popped:", gStrVal);
}

main();
