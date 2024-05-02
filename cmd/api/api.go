package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/jonathassk/travel_back_go/service/user"
	"net/http"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	return http.ListenAndServe(s.addr, router)
}
