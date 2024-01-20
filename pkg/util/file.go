package util

import (
	"encoding/json"
	"os"
)

func ReadFromJSON(path string, target any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, target)
}
