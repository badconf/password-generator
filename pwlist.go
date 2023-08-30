package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// GenerateRandomString generates a random string based on the given character set and length
func GenerateRandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[randomIndex.Int64()]
	}
	return string(b)
}

func main() {
	// Initialize command line arguments
	nPtr := flag.Int("n", 10, "Length of the random strings")
	qPtr := flag.Int("q", 10, "Quantity of random strings to generate")
	lowerCase := flag.Bool("a", false, "Include lowercase letters")
	upperCase := flag.Bool("A", false, "Include uppercase letters")
	numbers := flag.Bool("1", false, "Include numbers")
	special := flag.Bool("spe", false, "Include special characters")

	// Parse command line arguments
	flag.Parse()

	// Define character sets
	lowerCaseSet := "abcdefghijklmnopqrstuvwxyz"
	upperCaseSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet := "0123456789"
	specialSet := "!@#$%^&*()-=_+[]{}|;:',.<>?/"

	// Build the final character set based on the flags
	var finalSet strings.Builder

	if *lowerCase {
		finalSet.WriteString(lowerCaseSet)
	}
	if *upperCase {
		finalSet.WriteString(upperCaseSet)
	}
	if *numbers {
		finalSet.WriteString(numberSet)
	}
	if *special {
		finalSet.WriteString(specialSet)
	}

	charset := finalSet.String()

	// Validate that the character set is not empty
	if len(charset) == 0 {
		fmt.Println("Error: No character sets selected. Please select at least one character set.")
		return
	}

	// Generate the random strings
	var result strings.Builder
	for i := 0; i < *qPtr; i++ {
		randomString := GenerateRandomString(*nPtr, charset)
		result.WriteString(randomString + "\n")
	}

	// Write to a text file
	file, err := os.Create("random_strings.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(result.String())
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully generated random strings and saved to random_strings.txt")
}
