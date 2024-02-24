package assets

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed dist/*
var assetsEmbedFs embed.FS

func BuildAssets() fs.FS {
	dist, err := fs.Sub(assetsEmbedFs, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return dist
}
