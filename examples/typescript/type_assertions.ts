export {};

// Go limitation: Go does not allow type assertions or type switches directly on type
// parameter values. You must first convert to `any`, adding boilerplate.

// TS: TS does not allow runtime type checks against interfaces, because
// interfaces are erased at compile time.

// Optional interfaces that items in a pipeline might implement.
interface Validatable {
  validate(): Error | null;
}

interface Loggable {
  logEntry(): string;
}

// A runtime instanceof check on the interface is illegal:
//
//   if (item instanceof Validatable) {   // ERROR: 'Validatable' only refers
//     ...                                // to a type, but is being used as
//   }                                    // a value here.
//
// WORKAROUND: use type-guard functions that perform structural runtime checks.
function isValidatable(v: unknown): v is Validatable {
  return typeof v === "object" && v !== null && "validate" in v;
}

function isLoggable(v: unknown): v is Loggable {
  return typeof v === "object" && v !== null && "logEntry" in v;
}

function process<T>(item: T): void {
  if (isValidatable(item)) {
    const err = item.validate();
    if (err !== null) {
      console.log("Validation failed:", err.message);
      return;
    }
    console.log("Validated OK");
  } else {
    console.log("No validation available");
  }

  // Same limitation for a plain type assertion
  // (item as Loggable).logEntry()        // compiles, but crashes at runtime if item is not Loggable
  //
  // WORKAROUND
  if (isLoggable(item)) {
    console.log("Log:", item.logEntry());
  } else {
    console.log("Not loggable");
  }
}

// --- Test Implementations ---

// Order implements both Validatable and Loggable.
class Order {
  constructor(
    public id: number,
    public total: number,
  ) {}

  validate(): Error | null {
    if (this.total <= 0) {
      return new Error(`order ${this.id} has non-positive total`);
    }
    return null;
  }

  logEntry(): string {
    return `Order#${this.id} total=${this.total.toFixed(2)}`;
  }
}

// RawEvent implements neither interface.
class RawEvent {
  constructor(public name: string) {}
}

function main() {
  console.log("--- Order (implements both) ---");
  process(new Order(1, 49.99));

  console.log();
  console.log("--- RawEvent (implements neither) ---");
  process(new RawEvent("click"));
}

main();
