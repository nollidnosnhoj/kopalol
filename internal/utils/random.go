package utils

import (
	"github.com/jaevor/go-nanoid"
	"go.step.sm/crypto/randutil"
)

func GenerateRandomId(size int) (string, error) {
	generator, err := nanoid.Custom("1234567890abcdefghijklmnopqrstuvwxyz", 10)
	if err != nil {
		return "", err
	}
	id := generator()
	return id, nil
}

func GenerateDeletionKey() (string, error) {
	return randutil.Alphanumeric(20)
}
