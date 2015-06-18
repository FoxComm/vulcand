package tomlng

import (
	"os"
	"path/filepath"
)

func pathAbs(inPath string) (string, error) {

	path, err := filepath.Abs(os.ExpandEnv(inPath))
	if err != nil {
		return "", err
	}

	return filepath.Clean(path), nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
