package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonathassk/travel_back_go/types"
	"github.com/jonathassk/travel_back_go/utils"
	"net/http"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	// get request body
	var payload types.RegistrationType
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User registered: %+v", payload)
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var payload types.RegistrationType
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	print("print do erro", err)

	if err == nil {
		utils.WriteError(w, http.StatusConflict, err)
		return
	}

	user, err := h.store.CreateUser(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	println("criacao")
	utils.WriteJson(w, http.StatusCreated, user)
}
