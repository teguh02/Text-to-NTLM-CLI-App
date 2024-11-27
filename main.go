package main

// Simple App CLI-Based Command to Generate NTHASH
// From the Plaintext Password
// Usage : texttontlm -p 12345678
// Result : 259745CB123A52AA2E693AAACCA2DB52
// If error, it will return "something_error"

import (
	"textToNTLM/helpers"
	"flag"
	"fmt"
)
	
func main() {
	// Get the plaintext password from the command line
	plainText := flag.String("p", "", "Plaintext Password")
	flag.Parse()

	// Generate NTHASH
	result := helpers.NTHASH(*plainText)

	// Print the result
	fmt.Println(result)
}

// Run the command below to generate NTHASH
// go run main.go -p 12345678
