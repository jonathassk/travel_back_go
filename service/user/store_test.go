package user

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jonathassk/travel_back_go/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserServiceHandler_RegisterNewUser(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Register new user", func(t *testing.T) {
		payload := createValidUser()
		marshalledPayload, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalledPayload))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.register).Methods("POST")
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}
	})

}

func createValidUser() types.RegistrationType {
	return types.RegistrationType{
		FirstName: "Jonathas",
		LastName:  "Santos",
		Email:     "jonathas09@gmail.com",
		Password:  "password",
		City:      "Uberlandia",
		Country:   "Brazil",
		Currency:  "BRL",
		Language:  "PT-BR",
	}
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u *types.RegistrationType) (*types.User, error) {
	return nil, nil
}
