package models

// Secret represents a GitHub secret with its metadata
type Secret struct {
	SecretLevel  string // organization or repository
	SecretType   string // Actions, Dependabot, or Codespaces
	SecretName   string // Name of the secret
	SecretAccess string // Visibility of the secret (all, private, or selected repositories)
}
