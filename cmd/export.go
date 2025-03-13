package cmd

import (
	"flag"
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
	apiClient "github.com/vila89/gh-org-secrets/internal/api"
	"github.com/vila89/gh-org-secrets/internal/utils"
)

func executeExport(args []string) error {
	// Reorder arguments to ensure flags come before positional arguments
	args = reorderArgs(args)

	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	debugShort := exportCmd.Bool("d", false, "To debug logging")
	debugLong := exportCmd.Bool("debug", false, "To debug logging") // Add long form

	// Define both short and long form of the output file flag
	var outputFile string
	exportCmd.StringVar(&outputFile, "f", "", "Path and Name of CSV file to export secrets to (required)")
	exportCmd.StringVar(&outputFile, "output", "", "Path and Name of CSV file to export secrets to (required)")

	hostname := exportCmd.String("hostname", "github.com", "GitHub Enterprise Server hostname")
	tokenShort := exportCmd.String("t", "", "GitHub personal access token for organization (default \"gh auth token\")")
	tokenLong := exportCmd.String("token", "", "GitHub personal access token for organization (default \"gh auth token\")")
	help := exportCmd.Bool("help", false, "Show help for command")

	// Only print debug information if debug flag is set
	exportCmd.Parse(args)

	// Use either the short or long version of the flags
	debug := *debugShort || *debugLong
	token := *tokenShort
	if *tokenLong != "" {
		token = *tokenLong
	}

	if debug {
		fmt.Println("Arguments after reordering:")
		for i, arg := range args {
			fmt.Printf("  %d: %s\n", i, arg)
		}

		fmt.Println("Parsed arguments:")
		fmt.Printf("  debug: %v\n", debug)
		fmt.Printf("  outputFile: %s\n", outputFile)
		fmt.Printf("  hostname: %s\n", *hostname)
		fmt.Printf("  token: %s\n", token)
		fmt.Printf("  help: %v\n", *help)
	}

	if *help {
		printUsage(exportCmd)
		return nil
	}

	if outputFile == "" {
		fmt.Println("Error: --output/-f is required")
		printUsage(exportCmd)
		return fmt.Errorf("--output/-f is required")
	}

	// Get organization from positional arguments in the reordered args
	if exportCmd.NArg() == 0 {
		fmt.Println("Error: organization is required")
		printUsage(exportCmd)
		return fmt.Errorf("organization is required")
	}

	org := exportCmd.Arg(0)
	fmt.Printf("Organization: %s\n", org)

	// Create a REST client with the provided hostname and token
	var client *api.RESTClient
	var err error

	if token != "" {
		opts := api.ClientOptions{
			Host:      *hostname,
			AuthToken: token,
		}
		client, err = api.NewRESTClient(opts)
	} else {
		client, err = api.DefaultRESTClient()
	}

	if err != nil {
		return fmt.Errorf("failed to create GitHub API client: %w", err)
	}

	fmt.Println("Fetching organization secrets...")
	secrets := apiClient.FetchSecrets(client, org, *hostname)
	utils.WriteCSV(outputFile, secrets)

	fmt.Printf("Successfully exported %d secrets from organization %s to %s\n", len(secrets), org, outputFile)

	return nil
}

func printUsage(cmd *flag.FlagSet) {
	fmt.Println("Export Actions, Dependabot, and/or Codespaces secrets report for an organization")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  gh org-secrets export <organization> [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	cmd.PrintDefaults()
}
