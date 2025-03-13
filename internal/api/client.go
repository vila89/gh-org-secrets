package api

import (
	"fmt"
	"strings"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/vila89/gh-org-secrets/internal/models"
)

// Secret represents a GitHub secret with its visibility
type Secret struct {
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
}

// secretResponse matches the GitHub API response structure
type secretResponse struct {
	Secrets []Secret `json:"secrets"`
}

// FetchSecrets fetches all organization secrets and returns them in our models.Secret format
func FetchSecrets(client *api.RESTClient, org string, hostname string) []models.Secret {
	var allSecrets []models.Secret

	// Fetch Actions secrets
	actionSecrets, _ := fetchOrgActionSecrets(client, org)
	for _, s := range actionSecrets {
		allSecrets = append(allSecrets, models.Secret{
			SecretLevel:  "organization",
			SecretType:   "Actions",
			SecretName:   s.Name,
			SecretAccess: getSecretAccess(s.Visibility, hostname),
		})
	}

	// Fetch Dependabot secrets
	dependabotSecrets, _ := fetchOrgDependabotSecrets(client, org)
	for _, s := range dependabotSecrets {
		allSecrets = append(allSecrets, models.Secret{
			SecretLevel:  "organization",
			SecretType:   "Dependabot",
			SecretName:   s.Name,
			SecretAccess: getSecretAccess(s.Visibility, hostname),
		})
	}

	// Fetch Codespaces secrets
	codespacesSecrets, _ := fetchOrgCodespacesSecrets(client, org)
	for _, s := range codespacesSecrets {
		allSecrets = append(allSecrets, models.Secret{
			SecretLevel:  "organization",
			SecretType:   "Codespaces",
			SecretName:   s.Name,
			SecretAccess: getSecretAccess(s.Visibility, hostname),
		})
	}

	return allSecrets
}

// Helper function to get the correct secretAccess based on hostname
func getSecretAccess(visibility string, hostname string) string {
	if visibility == "private" && !strings.EqualFold(hostname, "github.com") {
		return "private and internal repositories"
	}
	return visibility
}

func fetchOrgActionSecrets(client *api.RESTClient, org string) ([]Secret, error) {
	var result secretResponse
	err := client.Get(fmt.Sprintf("orgs/%s/actions/secrets", org), &result)
	if err != nil {
		fmt.Printf("Error fetching Actions secrets: %v\n", err)
		return []Secret{}, err
	}
	return result.Secrets, nil
}

func fetchOrgDependabotSecrets(client *api.RESTClient, org string) ([]Secret, error) {
	var result secretResponse
	err := client.Get(fmt.Sprintf("orgs/%s/dependabot/secrets", org), &result)
	if err != nil {
		fmt.Printf("Error fetching Dependabot secrets: %v\n", err)
		return []Secret{}, err
	}
	return result.Secrets, nil
}

func fetchOrgCodespacesSecrets(client *api.RESTClient, org string) ([]Secret, error) {
	var result secretResponse
	err := client.Get(fmt.Sprintf("orgs/%s/codespaces/secrets", org), &result)
	if err != nil {
		fmt.Printf("Error fetching Codespaces secrets: %v\n", err)
		return []Secret{}, err
	}
	return result.Secrets, nil
}
