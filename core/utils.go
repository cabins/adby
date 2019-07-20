package core

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func GetCacheFolder() string {
	homedir, err := homedir.Dir()

	if err != nil {
		return ""
	}

	return filepath.Join(homedir, ".adby", "cache")
}
