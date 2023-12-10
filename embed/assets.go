// Package embed provides funcs to embed FS
package embed

import (
	"embed"
	"io/fs"
	"log"
)

// EmbeddedAssets :
// return FS with assets in good directory
func EmbeddedAssets(embedFS embed.FS, subDir string) fs.FS {
	embeddedFiles, err := fs.Sub(embedFS, subDir)
	if err != nil {
		log.Fatal(err)
	}
	return embeddedFiles
}
