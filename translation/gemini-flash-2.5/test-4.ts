// Named struct
// In Go, a struct can be embedded. In TypeScript, we can model this with class inheritance
// when the embedded struct primarily contributes fields that the embedding struct then "owns".
// This approach maintains the same public API (e.g., `b.Name`) as in Go.
class Named {
  Name: string;

  // The constructor initializes the Name property.
  // In Go, `Named` fields are initialized when the embedding struct is created.
  // Here, we provide a constructor to allow for this initialization.
  constructor(name: string = "") {
    this.Name = name;
  }
}

// Box is a generic struct that embeds Named and holds a value of type T.
// type Box[T any] struct {
//   Named
//   Value T
// }
class Box<T> extends Named {
  Value: T;

  // The constructor for Box takes both name and value.
  // It calls `super(name)` to initialize the inherited `Name` property
  // and then initializes its own `Value` property.
  constructor(name: string, value: T) {
    super(name); // Pass the name to the base class constructor
    this.Value = value;
  }
}

// NewBox creates a new Box with a given name and value.
// func NewBox[T any](name string, v T) Box[T] {
//   return Box[T]{Named: Named{Name: name}, Value: v}
// }
function NewBox<T>(name: string, v: T): Box<T> {
  // In Go, the `Named` embedded struct is explicitly initialized within the composite literal.
  // In TypeScript, with class inheritance, the `Box` constructor handles
  // both its own fields and the inherited fields by calling `super()`.
  return new Box<T>(name, v);
}

// main function to demonstrate usage
// func main() {
//   b := NewBox("answer", 42)
//   fmt.Println(b.Name, b.Value) // answer 42
// }
function main() {
  const b = NewBox("answer", 42);
  console.log(b.Name, b.Value); // Outputs: answer 42
}

// Call main to execute the example
main();