package images

import "github.com/jaevor/go-nanoid"

func GenerateID() (string, error) {
	generator, err := nanoid.Custom("1234567890abcdefghijklmnopqrstuvwxyz", 10)
	if err != nil {
		return "", err
	}
	id := generator()
	return id, nil
}
