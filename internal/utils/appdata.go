package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetKonciergeDBPath() string {
	var baseDir string

	switch runtime.GOOS {
	case "windows":
		baseDir = os.Getenv("LOCALAPPDATA")
		if baseDir == "" {
			baseDir = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local")
		}
	default: // Linux, macOS
		baseDir = os.Getenv("XDG_DATA_HOME")
		if baseDir == "" {
			baseDir = filepath.Join(os.Getenv("HOME"), ".local", "share")
		}
	}

	appDir := filepath.Join(baseDir, "koncierge")

	// Ensure the folder exists
	err := os.MkdirAll(appDir, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create app data directory: %v", err)
	}

	return filepath.Join(appDir, "koncierge.sqlite")
}
