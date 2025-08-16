// cmd/cli/main.go
package main

import (
	"bufio"
	"fmt"
	"os"

	shared "github.com/joechung-msft/json-go/internal/shared"
)

func main() {
	fmt.Println("Enter JSON to parse:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		result := shared.Parse(input)
		fmt.Println("\nParse result:")
		fmt.Printf("%#v\n", result)
	}
}
