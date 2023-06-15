package commands

import (
	"errors"
	"os"
	"strings"
)

func GetSettings(args []string) (map[string]bool, error) {
	var settings map[string]bool

	for _, a := range args {
		if strings.HasPrefix(a, "--") {
			if len(a) < 3 {
				return nil, errors.New("invalid setting")
			}

			settings[a[2:]] = true
		}
	}

	return settings, nil
}

func GetText(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
