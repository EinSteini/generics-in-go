// Go: type Dog struct{ Name string }
// Im TypeScript wird eine Go-Struktur in der Regel als Klasse dargestellt.
class Dog {
    // Go: Name string
    // Die öffentliche Eigenschaft 'Name' wird direkt im Konstruktor deklariert.
    constructor(public Name: string) {}

    // Go: func (d Dog) String() string { return d.Name }
    // Die Go-Methode 'String()' wird in TypeScript/JavaScript typischerweise als 'toString()' implementiert,
    // um die standardmäßige String-Repräsentation eines Objekts zu liefern.
    toString(): string {
        return this.Name;
    }
}

// Go: func Describe(v any) string
// 'any' in Go wird hier als 'unknown' übersetzt, was typisch für TypeScript ist,
// wenn der Typ zur Laufzeit bestimmt wird. Dies erfordert explizite Typ-Guards.
function Describe(v: unknown): string {
    // Go: switch val := v.(type) {
    // In TypeScript werden Typ-Switches durch 'if-else if'-Ketten mit Typ-Guards realisiert.

    // Go: case bool:
    if (typeof v === 'boolean') {
        // TypeScript ist hier direkt äquivalent zur Go-Logik.
        if (v) {
            return "YES";
        }
        return "NO";
    }

    // Go: case int:
    // Go's 'int' entspricht in TypeScript 'number'. Um die Semantik von 'int' (ganze Zahl) zu bewahren,
    // wird 'Number.isInteger' verwendet.
    if (typeof v === 'number' && Number.isInteger(v)) {
        // Go: return fmt.Sprintf("#%d", val)
        // String-Formatierung erfolgt idiomatisch mit Template-Literalen.
        return `#${v}`;
    }

    // Go: case Dog:
    // Für benutzerdefinierte Typen wie 'Dog' verwenden wir 'instanceof' als Typ-Guard.
    if (v instanceof Dog) {
        // Go: return "Dog: " + val.Name
        // String-Verkettung oder Template-Literale sind hier möglich.
        return `Dog: ${v.Name}`;
    }

    // Go: default:
    // Go: return fmt.Sprint(v)
    // Die Funktion 'String(v)' in TypeScript ist eine gute Entsprechung zu Go's 'fmt.Sprint(v)'.
    // Sie versucht, einen beliebigen Wert in eine Zeichenkette zu konvertieren.
    // Für Objekte ruft sie 'v.toString()' auf, falls vorhanden.
    // Für primitive Typen (String, Number, Boolean, Null, Undefined) gibt sie deren String-Repräsentation zurück.
    return String(v);
}

// Go: func main() { ... }
// Eine Hauptfunktion, die den Code ausführt.
function main() {
    // Go: fmt.Println(Describe(true))             // YES
    console.log(Describe(true));

    // Go: fmt.Println(Describe(42))               // #42
    console.log(Describe(42));

    // Go: fmt.Println(Describe(Dog{Name: "Rex"})) // Dog: Rex
    // Instanziierung der 'Dog'-Klasse.
    console.log(Describe(new Dog("Rex")));

    // Go: fmt.Println(Describe("hello"))          // hello
    console.log(Describe("hello"));
}

// Aufruf der Hauptfunktion, um das Programm zu starten.
main();