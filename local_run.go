package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	sdkkms "github.com/sftpgo/sdk/kms"
	"github.com/drakkan/sftpgo/v2/internal/kms"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <password_file> <contents_file> <additional_data>\n", os.Args[0])
		os.Exit(1)
	}
	
	passwordFile := os.Args[1]
	contentsFile := os.Args[2]
	additionalData := os.Args[3]
	
	// Read password file
	passwordData, err := os.Open(passwordFile)
	if err != nil {
		fmt.Printf("Failed to open password file: %v\n", err)
		return
	}
	defer passwordData.Close()
	
	passwordBytes, err := io.ReadAll(passwordData)
	if err != nil {
		fmt.Printf("Failed to read password file: %v\n", err)
		return
	}
	password := strings.TrimSpace(string(passwordBytes))
	
	// Read contents file
	contentsData, err := os.Open(contentsFile)
	if err != nil {
		fmt.Printf("Failed to open contents file: %v\n", err)
		return
	}
	defer contentsData.Close()
	
	contentsBytes, err := io.ReadAll(contentsData)
	if err != nil {
		fmt.Printf("Failed to read contents file: %v\n", err)
		return
	}
	contents := strings.TrimSpace(string(contentsBytes))
	
	// Create a local secret directly using the base secret structure
	baseSecret := kms.BaseSecret{
		Status:         sdkkms.SecretStatusSecretBox,
		Payload:        contents,
		Key:            password,
		AdditionalData: additionalData,
		Mode:           0,
	}
	
	// Create local secret provider directly
	localSecret := kms.NewLocalSecret(baseSecret, "", "")
	
	// Decrypt the secret
	err = localSecret.Decrypt()
	if err != nil {
		fmt.Printf("Failed to decrypt: %v\n", err)
		return
	}

  fmt.Printf("%s\n", localSecret.GetPayload())
}
