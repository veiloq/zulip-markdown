package main

import (
	"fmt"
	"os"

	"github.com/veiloq/zulip-markdown/zlmd"
)

var version = "dev"

func main() {
	args := os.Args[1:]

	if len(args) > 0 && (args[0] == "-v" || args[0] == "--version") {
		fmt.Printf("ZLMD version %s\n", version)
		return
	}

	fmt.Println("Zulip Markdown (ZLMD) CLI")
	fmt.Println("A tool for working with Zulip-flavored Markdown")

	// Basic example of how to use the library
	if len(args) > 0 {
		input := args[0]
		result, err := zlmd.Process(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing markdown: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
	} else {
		fmt.Println("Usage: zlmd [markdown text]")
		fmt.Println("       zlmd -v | --version")
	}
}
