package core

import (
	"fmt"
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

func byte2kb(i int) (int, string) {
	s := int(i / 1024)
	return s, fmt.Sprintf("%dKB", s)
}

func kb2mb(i int) (int, string) {
	s := int(i / 1024)
	return s, fmt.Sprintf("%dMB", s)
}

func mb2gb(i int) (int, string) {
	s := int(i / 1024)
	return s, fmt.Sprintf("%dGB", s)
}

func Byte2Human(i int) string {
	kb, kbs := byte2kb(i)
	if kb > 1024 {
		mb, mbs := kb2mb(kb)
		if mb > 1024 {
			_, gbs := mb2gb(mb)
			return gbs
		}
		return mbs
	}
	return kbs
}
