package manager

import (
	"path/filepath"
	"strings"
)

var InstallationDirectory string

func GetApplicationDirectory(title string) string {
	return filepath.Clean(filepath.Join(InstallationDirectory, strings.ToLower(title)))
}
