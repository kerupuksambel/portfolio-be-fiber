package utils

import (
	"log"
	"os"
	"path/filepath"
)

func ExecPath() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	return filepath.Dir(execPath)
}
