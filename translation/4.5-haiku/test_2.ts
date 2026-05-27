class Stack<T> {
	private items: T[] = [];

	push(v: T): void {
		this.items.push(v);
	}

	pop(): [T | undefined, boolean] {
		if (this.items.length === 0) {
			return [undefined, false];
		}
		const top = this.items[this.items.length - 1];
		this.items = this.items.slice(0, this.items.length - 1);
		return [top, true];
	}

	len(): number {
		return this.items.length;
	}
}

// main
const s = new Stack<number>();
s.push(10);
s.push(20);
s.push(30);

console.log("Len:", s.len());
const [v, ok] = s.pop();
if (ok) {
	console.log("Popped:", v);
}