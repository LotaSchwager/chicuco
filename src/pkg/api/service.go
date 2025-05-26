package api

import (
	"chicuco/pkg/centers"
	"chicuco/pkg/contacto"
	"net/http"
)

// Api represents the API with its address and endpoints
type Api struct {
	Addr string
}

// Mount creates a new ServeMux with all endpoints
func (a *Api) Mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /contactos", contacto.HandlerContactos) // Funcion para obtener el contacto
	mux.HandleFunc("GET /centros", centers.HandlerCentros)      // Funcion para obtener los centros
	return mux
}

// Run starts the HTTP server
func (a *Api) Run(mux *http.ServeMux) error {
	server := &http.Server{
		Addr:    a.Addr,
		Handler: mux,
	}
	return server.ListenAndServe()
}
