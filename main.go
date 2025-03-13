package main

import (
	"fmt"
	"os"

	"github.com/vila89/gh-org-secrets/cmd"
)

func main() {
	fmt.Println("Starting gh-org-secrets")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
