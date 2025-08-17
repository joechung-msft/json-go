// cmd/cli/main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	shared "github.com/joechung-msft/json-go/internal/shared"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	result, err := safeParse(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pretty, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if _, err := fmt.Fprintln(os.Stdout, string(pretty)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func safeParse(input string) (any, error) {
	var panicErr error
	defer func() {
		if r := recover(); r != nil {
			panicErr = fmt.Errorf("%v", r)
		}
	}()
	result := shared.Parse(input)
	if panicErr != nil {
		return nil, panicErr
	}
	return result, nil
}
