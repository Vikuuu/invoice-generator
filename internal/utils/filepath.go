package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func GetLinuxAppPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", "parmaan-patr"), nil
}

func GetWinAppPath() (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return "", errors.New("LOCALAPPDATA not set")
	}
	return filepath.Join(localAppData, "parmaan-patr"), nil
}

func GetDarwinAppPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "Application Support", "parmaan-patr"), nil
}
