package main
// go build -buildmode=plugin -o main.so func-plugin.go
import "fmt"

// Add two integers
func Add(a int, b int) int {
	fmt.Printf("\nAdding a=%d and b=%d\n", a, b)
	return a + b
}

// Sub two integers
func Sub(a int, b int) int {
	fmt.Printf("\nSubtracting b= %d from a=%d\n", b, a)
	return a - b
}