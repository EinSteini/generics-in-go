// Go limitation: The use of type sets is strictly limited to built-in primitives
// TS: union constraints achieve the same — here number | string.

type Sortable = number | string;

function sort<T extends Sortable>(vals: T[]): T[] {
  return [...vals].sort((a, b) => (a < b ? -1 : a > b ? 1 : 0));
}

// ERROR in Go: cannot use struct types in a type set
// TS: Works fine, union of object shapes is no problem
type Coordinate =
  | { x: number; y: number }
  | { x: number; y: number; z: number };

function printCoord(c: Coordinate): void {
  if ("z" in c) {
    console.log(`Coord3D(${c.x}, ${c.y}, ${c.z})`);
  } else {
    console.log(`Coord2D(${c.x}, ${c.y})`);
  }
}

// ERROR in Go: type sets with non-basic named types cannot be used for operations
type Point = { x: number; y: number; describe(): string };
type Segment = { start: Point; end: Point; describe(): string };

type Geometry = Point | Segment;

function describeGeometry<T extends Geometry>(g: T): string {
  // g.x;               // error: Property 'x' does not exist on type 'Geometry'. Property 'x' does not exist on type 'Segment'.

  // For fields, we must narrow down the type first:
  if ("start" in g) {
    return `Segment(Point(${g.start.x}, ${g.start.y}) -> Point(${g.end.x}, ${g.end.y}))`;
  }
  return `Point(${g.x}, ${g.y})`;
}

// Shared methods on all union members can be called directly
function describeViaMethod<T extends Geometry>(g: T): string {
  return g.describe();
}

function makePoint(x: number, y: number): Point {
  return { x, y, describe: () => `Point(${x}, ${y})` };
}

function makeSegment(start: Point, end: Point): Segment {
  return {
    start,
    end,
    describe: () => `Segment(${start.describe()} -> ${end.describe()})`,
  };
}

function main() {
  const ints = [30, 10, 20];
  console.log("Sorted ints:", sort(ints));

  const floats = [3.0, 1.5, 2.5];
  console.log("Sorted floats:", sort(floats));

  const words = ["cherry", "apple", "banana"];
  console.log("Sorted strings:", sort(words));

  console.log();
  // Struct union in a type set — ERROR in Go, works in TS
  const coord2d: Coordinate = { x: 1, y: 2 };
  const coord3d: Coordinate = { x: 3, y: 4, z: 5 };
  printCoord(coord2d); // Coord2D(1, 2)
  printCoord(coord3d); // Coord3D(3, 4, 5)

  console.log();

  // For fields on custom types, we need to narrow down the type first
  const p = makePoint(1, 2);
  const p2 = makePoint(3, 4);
  const seg = makeSegment(makePoint(0, 0), makePoint(1, 1));
  const seg2 = makeSegment(makePoint(2, 2), makePoint(4, 5));
  console.log(describeGeometry(p)); // Point(1, 2)
  console.log(describeGeometry(p2)); // Point(3, 4)
  console.log(describeGeometry(seg)); // Segment(Point(0, 0) -> Point(1, 1))
  console.log(describeGeometry(seg2)); // Segment(Point(2, 2) -> Point(4, 5))

  console.log();

  // Shared methods on all union members can be called directly without narrowing
  console.log(describeViaMethod(p)); // Point(1, 2)
  console.log(describeViaMethod(seg)); // Segment(Point(0, 0) -> Point(1, 1))

  console.log();
}

main();
