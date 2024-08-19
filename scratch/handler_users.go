package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tantran21501/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter,r *http.Request){
	type parameters struct{
		Name string `json:"name"`

	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Error parsing JSON: %v",err))
		return
	}
	user,err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Coudn't create user: %v",err))
		return
	}


	responseWithJson(w,201,databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter,r *http.Request, user database.User){
	responseWithJson(w,201,databaseUserToUser(user)) 

}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter,r *http.Request, user database.User){
	posts,err := apiCfg.DB.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		responseWithError(w,400,fmt.Sprintf("Coudn't create posts: %v",err))
		return
	}
	responseWithJson(w,201,databasePostsToPosts(posts)) 

}