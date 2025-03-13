package delivery

import (
	"net/http"

	"gitea.federated.computer/ross/federated-dash.git/internal/manager"
)

func TurnOff(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Please give a correct title", http.StatusBadRequest)
		return
	}

	err := manager.TurnOff(
		manager.GetApplicationDirectory(title),
		manager.Default_ApplicationServiceFileanme,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
