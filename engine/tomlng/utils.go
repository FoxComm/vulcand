package tomlng

import (
	"os"
	"path/filepath"
	"reflect"
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

// returns list of string keys for map[string]interface{}
// returns empty list for type mismatching
func mapStringKeys(m interface{}) []string {
	v := reflect.ValueOf(m)
	keys := []string{}

	if v.Kind() != reflect.Map {
		return keys
	}

	valueKeys := v.MapKeys()
	for _, k := range valueKeys {
		if k.Kind() != reflect.String {
			return keys
		}
		keys = append(keys, k.String())
	}
	return keys
}
