//go:build ignore

package main

import "fmt"

// Snippet 8: Type Switch mit primitiven und komplexen Typen
// Testet: any.(type) Pattern, typeof versagt bei Structs

type Dog struct{ Name string }

func (d Dog) String() string { return d.Name }

func Describe(v any) string {
	switch val := v.(type) {
	case bool:
		if val {
			return "YES"
		}
		return "NO"
	case int:
		return fmt.Sprintf("#%d", val)
	case Dog:
		return "Dog: " + val.Name
	default:
		return fmt.Sprint(v)
	}
}

func main() {
	fmt.Println(Describe(true))             // YES
	fmt.Println(Describe(42))               // #42
	fmt.Println(Describe(Dog{Name: "Rex"})) // Dog: Rex
	fmt.Println(Describe("hello"))          // hello
}