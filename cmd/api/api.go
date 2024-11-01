package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/QuocAnh189/DemoEcom/services/cart"
	"github.com/QuocAnh189/DemoEcom/services/order"
	"github.com/QuocAnh189/DemoEcom/services/product"
	"github.com/QuocAnh189/DemoEcom/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (server *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "Successfully updated"})
	}).Methods(http.MethodGet)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(server.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(server.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(server.db)
	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Print("Server is running on ", server.addr)

	return http.ListenAndServe(server.addr,router)
}