package handlers

import (
	"encoding/json"
	"net/http"
	"projeto-modelo/internal/dto"
	"projeto-modelo/internal/entity"
	"projeto-modelo/internal/infra/database"
	"time"

	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
	Jwt *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB: db,
		Jwt: jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userEntity, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(userEntity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	user, err := h.UserDB.FindByEmail(input.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	if !user.ValidatePassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]any{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
    


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

