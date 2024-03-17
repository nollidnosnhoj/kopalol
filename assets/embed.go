package assets

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed dist/*
var publicEmbedFs embed.FS

func BuildPublicDistFs() fs.FS {
	dist, err := fs.Sub(publicEmbedFs, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return dist
}
