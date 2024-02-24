package images

import "github.com/jaevor/go-nanoid"

func GenerateID() (string, error) {
	generator, err := nanoid.Standard(8)
	if err != nil {
		return "", err
	}
	id := generator()
	return id, nil
}
