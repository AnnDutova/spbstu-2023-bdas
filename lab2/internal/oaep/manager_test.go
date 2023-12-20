package oaep_test

import (
	"testing"

	"github.com/AnnDutova/bdas/lab2/internal/oaep"
)

func TestGenerateEncryptDecrypt(t *testing.T) {
	bits := 2048

	// Генерация ключевой пары
	privateKey, publicKey, err := oaep.GenerateKeyPair(bits)
	if err != nil {
		t.Fatalf("Error generating key pair: %v", err)
	}

	// Шифрование и дешифрование данных
	originalData := []byte("Hello, World!")

	encryptedData, err := oaep.Encrypt(publicKey, originalData)
	if err != nil {
		t.Fatalf("Error encrypting data: %v", err)
	}

	decryptedData, err := oaep.Decrypt(privateKey, encryptedData)
	if err != nil {
		t.Fatalf("Error decrypting data: %v", err)
	}

	// Проверка, что оригинальные данные совпадают с дешифрованными
	if string(originalData) != string(decryptedData) {
		t.Errorf("Decrypted data doesn't match the original data.")
	}
}

func TestSignVerify(t *testing.T) {
	bits := 2048

	// Генерация ключевой пары
	privateKey, publicKey, err := oaep.GenerateKeyPair(bits)
	if err != nil {
		t.Fatalf("Error generating key pair: %v", err)
	}

	// Подписание и верификация данных
	data := []byte("Hello, World!")

	signature, err := oaep.Sign(privateKey, data)
	if err != nil {
		t.Fatalf("Error signing data: %v", err)
	}

	verified := oaep.Verify(publicKey, data, signature)
	if !verified {
		t.Error("Signature verification failed.")
	}
}
