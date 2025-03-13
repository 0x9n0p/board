package manager_test

import (
	"testing"

	"gitea.federated.computer/ross/federated-dash.git/internal/manager"
)

func TestTurnOff(t *testing.T) {
	err := manager.TurnOff(
		manager.GetApplicationDirectory("testdata/nextcloud"),
		manager.Default_ApplicationServiceFileanme,
	)
	if err != nil {
		t.Errorf("Failed to turn off: %v", err)
	}
}
