package main

import (
	"flag"
	"fmt"

	"github.com/AnnDutova/bdas/lab2/internal/oaep"
)

var (
	bitSize int
)

func init() {
	flag.IntVar(&bitSize, "bitSize", 2048, "bit size for generating RSA private key")
}

func main() {
	privateKey, publicKey, err := oaep.GenerateKeyPair(bitSize)
	if err != nil {
		fmt.Println("Error generating key pair:", err)
		return
	}

	originalData := []byte("Hello, World!")

	encryptedData, err := oaep.Encrypt(publicKey, originalData)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	decryptedData, err := oaep.Decrypt(privateKey, encryptedData)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Original Data:", string(originalData))
	fmt.Println("Decrypted Data:", string(decryptedData))

	signature, err := oaep.Sign(privateKey, originalData)
	if err != nil {
		fmt.Println("Error signing data:", err)
		return
	}

	verified := oaep.Verify(publicKey, originalData, signature)
	if verified {
		fmt.Println("Signature verified successfully.")
	} else {
		fmt.Println("Signature verification failed.")
	}
}
