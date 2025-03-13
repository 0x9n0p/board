package manager

import "strings"

func GetApplicationDirectory(title string) string {
	return strings.ToLower(title)
}
