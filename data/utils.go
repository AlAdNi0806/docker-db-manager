// Package data provides encrypted JSON storage.
package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"golang.org/x/crypto/argon2"
)

const (
	filename = "data.json.enc"
	password = "your-very-strong-secret-password" // 🔑 CHANGE THIS!
)

// SecureStore handles encrypted JSON storage.
type SecureStore struct {
	filePath string
	key      []byte
}

// NewSecureStore creates a new encrypted store.
func NewSecureStore() (*SecureStore, error) {
	// Get XDG-compliant data directory
	dataDir, err := xdg.DataFile("manage-db")
	if err != nil {
		return nil, fmt.Errorf("failed to get data dir: %w", err)
	}
	filePath := filepath.Join(filepath.Dir(dataDir), filename)

	// Derive encryption key from password
	key := deriveKey(password)

	return &SecureStore{
		filePath: filePath,
		key:      key,
	}, nil
}

// deriveKey creates a 32-byte key using Argon2
func deriveKey(password string) []byte {
	salt := []byte("manage-db-salt-1234567890ab") // 16 bytes
	return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}

// Load decrypts and unmarshals data into the provided struct.
// If file doesn't exist, returns nil without error.
func (s *SecureStore) Load(v interface{}) error {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(s.filePath), 0700); err != nil {
		return err
	}

	data, err := os.ReadFile(s.filePath)
	if os.IsNotExist(err) {
		return nil // no data yet
	}
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	decrypted, err := s.decrypt(data)
	if err != nil {
		return fmt.Errorf("failed to decrypt: %w", err)
	}

	if len(decrypted) == 0 {
		return nil
	}

	if err := json.Unmarshal(decrypted, v); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
}

// Save marshals data to JSON, encrypts it, and writes to disk.
func (s *SecureStore) Save(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	encrypted, err := s.encrypt(data)
	if err != nil {
		return fmt.Errorf("failed to encrypt: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(s.filePath), 0700); err != nil {
		return err
	}

	if err := os.WriteFile(s.filePath, encrypted, 0600); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// encrypt uses AES-GCM to encrypt data
func (s *SecureStore) encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decrypt uses AES-GCM to decrypt data
func (s *SecureStore) decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < 12 { // GCM nonce is 12 bytes
		return nil, fmt.Errorf("ciphertext too short")
	}

	block, err := aes.NewCipher(s.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return plaintext, nil
}
