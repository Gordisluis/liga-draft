package user

import (
	"api/types"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestUserServiceHabler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore) 

	t.Run(" deberia dar error si el user es invalido", func(t *testing.T) {
		payload := types.RegisterUserPayLoad{
			FirstName: "user",
			LastName:  "123",
			Email:     "",
			Password:  "test",
		}
		marshalled := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer((marshalled)))
		if err != nil {
			t.Fatal(err)
		}	

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("se esperaba status code &d, obtuvo &d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User)
error {
	return nil
}