nterface Stringer {
	String(): string;
}

interface Validator {
	IsValid(): boolean;
}

interface Entry extends Stringer, Validator {}

function PrintValid<T extends Entry>(items: T[]): void {
	for (const item of items) {
		if (item.IsValid()) {
			console.log("✓", item.String());
		} else {
			console.log("✗", item.String());
		}
	}
}

// --- Konkreter Typ ---

class Email implements Entry {
	Address: string;

	constructor(address: string) {
		this.Address = address;
	}

	String(): string {
		return this.Address;
	}

	IsValid(): boolean {
		return this.Address.length > 3;
	}
}

const emails: Email[] = [
	new Email("alice@example.com"),
	new Email("ab"),
	new Email("bob@test.org"),
];

PrintValid(emails);