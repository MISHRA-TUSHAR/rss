package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TUSHAR-MISHRA/rss/internal/auth"
	"github.com/TUSHAR-MISHRA/rss/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// type parameters struct {
	// 	Name string `json:"name"`
	// }
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	user, err := apiCfg.Db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	respondWithJSON(w, 201, databaseUserToUser(user))

}

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
		return
	}

	user, err := cfg.Db.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
