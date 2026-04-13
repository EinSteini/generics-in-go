// Go limitation: Methods are not allowed to introduce type parameters

// Container is a generic wrapper holding an array of T.
class Container<T> {
  constructor(public items: T[]) {}

  // In TypeScript it is perfectly valid — methods may have their
  // own type parameters independently of the class type parameter.
  map<U>(f: (v: T) => U): Container<U> {
    return new Container(this.items.map(f));
  }
}

function main() {
  const nums = new Container([1, 2, 3, 4]);

  // Works, methods can have their own type parameters independent of the class's T.
  const strs = nums.map((n) => "*".repeat(n));
  console.log(strs.items); // Output: [ '*', '**', '***', '****' ]
}

main();
