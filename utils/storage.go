package utils

import (
	"log"
	"os"
	"path/filepath"
)

func StoragePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(cwd, "storage")
	return path
}
