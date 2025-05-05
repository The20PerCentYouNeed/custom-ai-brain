package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func StoragePath(path ...string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	absPath := filepath.Join(cwd, "storage")
	if len(path) > 0 {
		absPath = filepath.Join(absPath, strings.Join(path, "/"))
	}

	return absPath
}
