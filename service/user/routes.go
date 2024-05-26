package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonathassk/travel_back_go/types"
	"github.com/jonathassk/travel_back_go/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
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
	router.HandleFunc("/update", h.UpdateUser).Methods("PUT")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	// get request body
	var payload types.LoginType
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err, status := verifyUserLogin(payload, *user)
	if err != nil {
		utils.WriteError(w, status, err)
		return
	}

	token, err := utils.CreateNewToken(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJson(w, http.StatusOK, "login successful, token: "+token)
	if err != nil {
		return
	}
	return
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var payload types.RegistrationType
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	err := verifyUserPayload(payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userEmail, err := h.store.GetUserByEmail(payload.Email)
	if userEmail != nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	user, err := h.store.CreateUser(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	token, err := utils.CreateNewToken(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	result := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	err = utils.WriteJson(w, http.StatusCreated, result)
	if err != nil {
		return
	}

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("no token provided"))
		return
	}
	tokenString = tokenString[7:] // remove "Bearer " from token
	err := utils.VerifyToken(tokenString)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	err = utils.WriteJson(w, http.StatusOK, "user updated")
	if err != nil {
		return
	}
	return
}

func verifyUserPayload(payload types.RegistrationType) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	for _, field := range []string{payload.FirstName, payload.LastName, payload.Email, payload.Password, payload.City, payload.Country, payload.Currency, payload.Language} {
		if len(field) < 2 {
			return fmt.Errorf("all fields must have at least 2 characters")
		}
	}
	if payload.Language[2] != '-' {
		return fmt.Errorf("language must be in the format 'XX-XX'")
	}
	if emailRegex.MatchString(payload.Email) == false {
		return fmt.Errorf("invalid email")
	}
	if len(payload.Password) < 8 {
		return fmt.Errorf("password must have at least 8 characters")
	}
	return nil
}

func verifyUserLogin(payload types.LoginType, user types.User) (error error, status int) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(payload.Email) == false {
		return fmt.Errorf("invalid email"), http.StatusBadRequest
	}
	if len(payload.Password) < 8 {
		return fmt.Errorf("password must have at least 8 characters"), http.StatusBadRequest
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return fmt.Errorf("invalid password"), http.StatusUnauthorized
	}

	return nil, http.StatusOK
}
