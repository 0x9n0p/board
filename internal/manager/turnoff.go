package manager

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	Default_ApplicationServiceFileanme = ApplicationServiceFilename
	ApplicationServiceFilename         = "service"
)

func TurnOff(directory string, service string) error {
	path := filepath.Clean(filepath.Join(directory, ApplicationServiceFilename))

	variables, err := godotenv.Read()
	if err != nil {
		return fmt.Errorf("could not read environment variables (%s): %w", path, err)
	}

	status, found := variables[Dashboard_Key]
	if !found {
		status = Dashboard_StatusYes
	}

	if status != Dashboard_StatusYes {
		// TODO: debug log
		return nil
	}

	variables[Dashboard_Key] = Dashboard_StatusNo

	if err := godotenv.Write(variables, path); err != nil {
		return fmt.Errorf("could not write environment variables (%s): %w", path, err)
	}

	// TODO: debug log
	return nil
}
