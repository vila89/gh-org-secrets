# gh-org-secrets

GitHub CLI extension to export Actions, Dependabot, and Codespaces secrets from an organization into a CSV file.

## Installation
```sh
gh extension install vila89/gh-org-secrets
```

## Usage
```sh
gh org-secrets export <organization> -f [secrets.csv]
```

Options:

```
-d, --debug              Enable debug logging
-f, --output string      Path and name of CSV file to export secrets to (required)
    --hostname string    GitHub Enterprise Server hostname (default "github.com")
-t, --token string       GitHub personal access token (default "gh auth token")
```
