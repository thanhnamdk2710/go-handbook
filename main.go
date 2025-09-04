package main

import (
	"fmt"
)

// Function to greet a person
func greet(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// Function to check if a number is even
func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	// Demonstrate function calls
	message := greet("Gopher")
	fmt.Println(message)

	// Demonstrate for loop and if statements
	fmt.Println("\nChecking numbers from 1 to 5:")
	for i := 1; i <= 5; i++ {
		if isEven(i) {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}

	// Demonstrate slice (dynamic array)
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("\nNumbers:", numbers)

	// Demonstrate map (key-value pairs)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println("\nColors:")
	for name, hex := range colors {
		fmt.Printf("%s: %s\n", name, hex)
	}
}
