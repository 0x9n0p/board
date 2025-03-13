package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var DefaultInstallationConfig = map[string]string{
	Dashboard_Key: Dashboard_StatusYes,
}

func Install(directory string, service string) error {
	if err := os.MkdirAll(directory, 0755); err != nil {
		return fmt.Errorf("could not create application directory (%s): %v", directory, err)
	}

	path := filepath.Clean(filepath.Join(directory, service))

	if err := godotenv.Write(DefaultInstallationConfig, path); err != nil {
		return fmt.Errorf("could not write environment variables (%s): %w", path, err)
	}

	return nil
}
