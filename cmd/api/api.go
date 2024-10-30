package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (server *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// subrouter := router.PathPrefix("/api/v1").Subrouter()

	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Anh Quoc"))
	}).Methods("GET")

	log.Print("Server is running on ", server.addr)

	return http.ListenAndServe(server.addr,router)
}