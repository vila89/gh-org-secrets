package utils

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/vila89/gh-org-secrets/internal/models"
)

// WriteCSV writes the secrets to a CSV file
func WriteCSV(filename string, secrets []models.Secret) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"SecretLevel", "SecretType", "SecretName", "SecretAccess"})
	for _, secret := range secrets {
		writer.Write([]string{secret.SecretLevel, secret.SecretType, secret.SecretName, secret.SecretAccess})
	}

	fmt.Printf("Successfully exported %d secrets to %s\n", len(secrets), filename)
}
