package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
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

func GetAssetsAppPath() (string, error) {
	var appPath string
	var err error
	switch runtime.GOOS {
	case "linux":
		appPath, err = GetLinuxAppPath()
		if err != nil {
			return "", err
		}
		return filepath.Join(appPath, "assets"), nil
	case "windows":
		appPath, err = GetWinAppPath()
		if err != nil {
			return "", err
		}
		return filepath.Join(appPath, "assets"), nil
	case "darwin":
		appPath, err = GetDarwinAppPath()
		if err != nil {
			return "", err
		}
		return filepath.Join(appPath, "assets"), nil
	default:
		return "", errors.New("unsupported platform")
	}
}
