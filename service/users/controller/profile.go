package controller

import (
	"base/pkg/auth"
	"base/pkg/constant"
	"base/pkg/log"
	"base/pkg/router"
	"base/service/users/model"
	"encoding/json"
	"net/http"
)

// GetCurrentProfile Function : Get Request'User Profile
// Need check Jwt Token and Get Info From JWT Token
func GetCurrentProfile(w http.ResponseWriter, r *http.Request) {
	var userPayload auth.Payload

	// Use This Info In JWT Token To Show
	// Don't Use Database
	claims := r.Header.Get(constant.JwtClaimsHeader)
	err := json.Unmarshal([]byte(claims), &userPayload)

	if err != nil {
		router.ResponseInternalError(w, constant.UnmarshalFail, err.Error())
		return
	}
	router.ResponseSuccessWithData(w, "", "", userPayload)
	return
}

// GetCurrentProfile Function : Get Request'User Profile
// Need check Jwt Token and Get Info From JWT Token
func UpdateCurrentProfile(w http.ResponseWriter, r *http.Request) {
	var userPayload auth.Payload
	var changeProfile ChangeProfileForm

	_ = json.NewDecoder(r.Body).Decode(&changeProfile)
	if changeProfile.Username == nil || changeProfile.ProfilePictureURL == nil {
		router.ResponseBadGateway(w, constant.MissingFieldInputed, constant.MissingFieldInputed)
		return
	}

	claims := r.Header.Get(constant.JwtClaimsHeader)
	err := json.Unmarshal([]byte(claims), &userPayload)

	// Map payload Object To Model
	var user = model.User{
		ID: userPayload.ID,
	}

	// Check Field Of Object
	if userPayload.ProfilePictureURL == nil {
		userPayload.ProfilePictureURL = changeProfile.ProfilePictureURL
	} else {
		if *userPayload.ProfilePictureURL != *changeProfile.ProfilePictureURL {
			userPayload.ProfilePictureURL = changeProfile.ProfilePictureURL
		}
	}
	if userPayload.Username != *changeProfile.Username {
		userPayload.Username = *changeProfile.Username
	}

	err = user.UpdateProfile(userPayload.Username, *userPayload.ProfilePictureURL)
	if err != nil {
		log.Println(log.LogLevelError, "update-profile-db", err.Error())
		router.ResponseInternalError(w, constant.UpdateProfileFail, err.Error())
		return
	}
	// Update JWT TOKEN

	// Must convert userPayload to string of json
	userPayloadJson, err := json.Marshal(userPayload)
	if err != nil {
		router.ResponseInternalError(w, constant.UnmarshalFail, err.Error())
		log.Println(log.LogLevelDebug, "Login: Marshal Payload Data To Json", err)
		return
	}

	token, err := auth.CreateJWTToken(string(userPayloadJson))
	if err != nil {
		log.Println(log.LogLevelDebug, "Login: CreateJWTToken", err)
		router.ResponseInternalError(w, constant.CreateJWTTokenFail, err.Error())
		return
	}

	tokendata := TokenData{
		Token: token,
		Type:  constant.BearToken,
	}

	router.ResponseSuccessWithData(w, constant.UpdateProfileSuccsess, constant.UpdateProfileSuccsess, tokendata)
	return
}
