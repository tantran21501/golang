package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/tantran21501/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter,r *http.Request, user database.User){
	type parameters struct{
		FeedID uuid.UUID `json:"feed_id"`

	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Error parsing JSON: %v",err))
		return
	}
	feedFollow,err := apiCfg.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,

	})
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Coudn't create feed follow: %v",err))
		return
	}


	responseWithJson(w,201,databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter,r *http.Request,user database.User){
	feedFollows,err := apiCfg.DB.GetFeedFollow(r.Context(),user.ID)
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Coudn't get feed follows: %v",err))
		return
	}
	responseWithJson(w,201,databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter,r *http.Request,user database.User){
	feedFollowIdStr := chi.URLParam(r,"feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Coudn't parse feed follow Id: %v",err))
		return
	}

	
	err = apiCfg.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID:  feedFollowId,
		UserID: user.ID,
	})
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Coudn't delete feed follows: %v",err))
		return
	}
	responseWithJson(w,201,struct{}{})
}
