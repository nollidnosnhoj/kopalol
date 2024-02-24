package assets

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed dist/*
var embedFs embed.FS

func Build() fs.FS {
	dist, err := fs.Sub(embedFs, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return dist
}
