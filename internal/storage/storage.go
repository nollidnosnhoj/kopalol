package storage

import "io"

type Storage interface {
	Upload(filename string, source io.Reader) error
}
