export {};

// Go limitation: Generic methods are not allowed to introduce type bounds

// Summable is a constraint requiring an add method.
interface Summable {
  add(other: Summable): Summable;
}

// Stack is a generic stack with no constraint on T.
class Stack<T> {
  private items: T[] = [];

  push(v: T): void {
    this.items.push(v);
  }

  // ERROR: cannot narrow T to Summable only for this method (same as Go).
  // Type 'T' is not assignable to type 'T'. Two different types with this name exist, but they are unrelated. 'T' could be instantiated with an arbitrary type which could be unrelated to 'T'.
  //
  // sumAll<T extends Summable>(): T {
  //   let acc = this.items[0];  // this.items is still the class T[], not the method T[]
  //   for (let i = 1; i < this.items.length; i++) {
  //     acc = acc.add(this.items[i]) as T;  // error: Type 'T' is not assignable to type 'T'
  //   }
  //   return acc;
  // }

  // WORKAROUND: use a typed 'this' parameter to constrain the method to only be callable on Stack<T> where T satisfies Summable.
  sumAll(this: Stack<T & Summable>): T & Summable {
    if (this.items.length === 0) throw new Error("empty stack");
    let acc = this.items[0];
    for (let i = 1; i < this.items.length; i++) {
      acc = acc.add(this.items[i]) as T & Summable;
    }
    return acc;
  }
}

class Money {
  constructor(public cents: number) {}

  add(other: Summable): Summable {
    return new Money(this.cents + (other as Money).cents);
  }

  toString(): string {
    return `$${Math.floor(this.cents / 100)}.${String(this.cents % 100).padStart(2, "0")}`;
  }
}

function main() {
  const s = new Stack<Money>();
  s.push(new Money(350));
  s.push(new Money(1275));
  s.push(new Money(99));

  // s.sumAll() works, because Money satisfies the Summable constraint required by the 'this' parameter of sumAll.
  const total = s.sumAll();
  console.log(total.toString()); // Output:  $17.24
}

main();
