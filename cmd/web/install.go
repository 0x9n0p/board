package main

import (
	"fmt"

	"gitea.federated.computer/ross/federated-dash.git/internal/manager"
)

func InstallApplications(apps []AppLink) error {
	for _, app := range apps {
		err := manager.Install(
			manager.GetApplicationDirectory(app.Title),
			manager.Default_ApplicationServiceFileanme,
		)
		if err != nil {
			return fmt.Errorf("install app %s: %v", app.Title, err)
		}
	}

	return nil
}
