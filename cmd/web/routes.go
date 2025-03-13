package main

import (
	"net/http"

	manager_delivery "gitea.federated.computer/ross/federated-dash.git/internal/manager/delivery"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	//Set up static file server
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//Set up routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /", app.notFound)
	mux.HandleFunc("GET /vpn/{$}", app.vpn)
	mux.HandleFunc("GET /apps/{$}", app.appList)
	mux.HandleFunc("GET /apps/turnoff", manager_delivery.TurnOff)

	return mux
}
