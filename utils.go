package main

import (
	"os"
	"runtime"

	"github.com/Vikuuu/invoice_generator/internal/utils"
)

func createAppDir() (string, error) {
	var basePath string
	var err error
	switch runtime.GOOS {
	case "linux":
		basePath, err = utils.GetLinuxAppPath()
	case "windows":
		basePath, err = utils.GetWinAppPath()
	case "darwin":
		basePath, err = utils.GetDarwinAppPath()
	default:
		return "", ErrUnsupportedPlatform
	}

	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(basePath, 0o755); err != nil {
		return "", err
	}
	return basePath, nil
}

func createAppAssetsDir() error {
	path, err := utils.GetAssetsAppPath()
	if err != nil {
		return err
	}

	if err = os.MkdirAll(path, 0o755); err != nil {
		return err
	}
	return nil
}

func getInvoiceDefaultStorePath() (string, error) {
	var path string
	var err error
	switch runtime.GOOS {
	case "linux":
		path, err = utils.GetLinuxDocumentsPath()
	case "windows":
		path, err = utils.GetWinDocumentsPath()
	case "darwin":
		path, err = utils.GetDarwinDocumentsPath()
	default:
		return "", ErrUnsupportedPlatform
	}

	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(path, 0o755); err != nil {
		return "", err
	}
	return path, nil
}
