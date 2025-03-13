package cmd

import (
	"fmt"
	"os"
)

func Execute() error {
	if len(os.Args) < 2 {
		fmt.Println("expected 'export' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "export":
		return executeExport(os.Args[2:])
	default:
		fmt.Println("expected 'export' subcommand")
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}
}

// reorderArgs puts flags before positional arguments
func reorderArgs(args []string) []string {
	flags := []string{}
	flagValues := []string{}
	positional := []string{}

	for i := 0; i < len(args); i++ {
		if len(args[i]) > 0 && args[i][0] == '-' {
			flags = append(flags, args[i])
			// If this is not the last argument and the next argument doesn't start with a dash
			// then it's a flag value
			if i+1 < len(args) && (len(args[i+1]) == 0 || args[i+1][0] != '-') {
				flagValues = append(flagValues, args[i+1])
				i++ // Skip the next argument as we've already processed it
			}
		} else {
			positional = append(positional, args[i])
		}
	}

	// Combine flags and their values, followed by positional arguments
	result := []string{}
	for i := 0; i < len(flags); i++ {
		result = append(result, flags[i])
		if i < len(flagValues) {
			result = append(result, flagValues[i])
		}
	}
	result = append(result, positional...)

	return result
}
