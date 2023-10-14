package controller

import (
	"base/pkg/auth"
	"base/pkg/constant"
	"base/pkg/db"
	"base/pkg/log"
	"base/pkg/router"
	"base/service/users/model"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const DefaultGeneratePasswordCost = 10

type TokenData struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

// Signin Function: Check Input, Create JWT Token
func Signin(w http.ResponseWriter, r *http.Request) {
	var signinForm SigninForm
	var userPayload auth.Payload

	var hashedPassword string

	_ = json.NewDecoder(r.Body).Decode(&signinForm)

	// Get the hashedPassword from database
	query := fmt.Sprintf("SELECT password, id, email, username, profile_picture_url, role FROM users WHERE email = '%s' limit 1", *signinForm.Email)

	err := db.PSQL.QueryRow(query).Scan(&hashedPassword, &userPayload.ID, &userPayload.Email,
		&userPayload.Username, &userPayload.ProfilePictureURL, &userPayload.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			router.ResponseBadRequest(w, constant.UserLoginFailed, err.Error())
			return
		} else {
			log.Println(log.LogLevelError, "query-email-database", err.Error())
			router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
			return
		}
	}

	// Compare the password with the hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(*signinForm.Password))
	if err != nil {
		// Login failed
		router.ResponseBadRequest(w, constant.UserLoginFailed, err.Error())
		return
	}

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

	// Create token data
	tokenData := TokenData{
		Token: token,
		Type:  constant.BearToken,
	}
	router.ResponseSuccessWithData(w, constant.UserLoginSuccess, "", tokenData)
	return
}

// CreateAccount Function: Check Input, Hash Password, Create Account
func Signup(w http.ResponseWriter, r *http.Request) {
	// Singup Form Register
	var signUpForm = SignupForm{}
	_ = json.NewDecoder(r.Body).Decode(&signUpForm)

	// Check Input
	if signUpForm.Username == nil || signUpForm.Email == nil || signUpForm.Password == nil {
		router.ResponseBadRequest(w, constant.MissingFieldInputed, "missing email || password || username")
		return
	}

	// Valid Email
	_, err := mail.ParseAddress(*signUpForm.Email)
	if err != nil {
		router.ResponseBadRequest(w, constant.EmailIsNotValid, "email is not valid")
		return
	}
	// Check Host of Email
	emailPart := strings.SplitN(*signUpForm.Email, "@", 2)
	emailHost := emailPart[1]
	mx, err := net.LookupMX(emailHost)
	if err != nil {
		router.ResponseBadRequest(w, constant.EmailHostIsNotExist, "email is not valid")
		return
	}
	if len(mx) == 0 {
		router.ResponseBadRequest(w, constant.EmailHostIsNotExist, "email is not valid")
		return
	}
	newUser := model.User{
		Username: signUpForm.Username,
		Email:    signUpForm.Email,
		Password: nil,
		IsActive: nil,
		Role:     nil,
	}
	// Check Email Exist
	emailExits, err := newUser.CheckEmailExist()
	if err != nil {
		log.Println(log.LogLevelError, "query-email-exits", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
		return
	}

	if emailExits {
		router.ResponseBadRequest(w, constant.EmailExist, "this email is existed")
		return
	}
	// Generate Hashed Password, Use Generate Cost : 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*signUpForm.Password), DefaultGeneratePasswordCost)
	if err != nil {
		router.ResponseInternalError(w, constant.GenerateFromPasswordFaild, err.Error())
		return
	}
	password := string(hashedPassword)
	isActive := true
	role := "user"

	// You Can Confirm Email or SMS Code Here To Update IsActive
	// In This Sample Services, I dont
	newUser.Password = &password
	newUser.IsActive = &isActive
	newUser.Role = &role

	err = newUser.CreateNewUser()
	if err != nil {
		router.ResponseInternalError(w, constant.UserRegisterFail, err.Error())
		return
	}

	router.ResponseCreatedWithData(w, constant.UserRegisterSuccsess, constant.UserRegisterSuccsess)
	return
}

// Signout Function
// Actualy, I Should Use Add JWT Token Into Blacklist, To Check Auth -- Soon Available
// In This Version I use
// Frontend application should remove the JWT token from the client-side storage.
func Signout(w http.ResponseWriter, r *http.Request) {
	// Actualy, I Should Use Add JWT Token Into Blacklist, To Check Auth -- Soon Available
	router.ResponseSuccess(w, constant.UserLogoutSuccsess, constant.UserLogoutSuccsess)
	return
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Form Change Password
	var changePassword ChangePasswordForm
	_ = json.NewDecoder(r.Body).Decode(&changePassword)

	if changePassword.NewPassword == nil || changePassword.Password == nil {
		router.ResponseBadRequest(w, constant.MissingFieldInputed, constant.MissingFieldInputed)
		return
	}

	// Get Info Of User From JWT Token
	var user model.User
	claims := r.Header.Get(constant.JwtClaimsHeader)

	err := json.Unmarshal([]byte(claims), &user)

	if err != nil {
		router.ResponseInternalError(w, constant.UnmarshalFail, err.Error())
		return
	}

	// Get Hashed Password
	// Get the username, email hashedPassword from database
	var hashed_password string

	query := fmt.Sprintf("SELECT password FROM users WHERE id = '%d'", user.ID)
	err = db.PSQL.QueryRow(query).Scan(&hashed_password)

	if err != nil {
		log.Println(log.LogLevelError, "query-password-fail", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
		return
	}
	//Compare the password with the hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(*changePassword.Password))
	if err != nil {
		router.ResponseBadRequest(w, constant.UserPasswordWrong, err.Error())
		log.Println(log.LogLevelDebug, "Logout: CompareHashAndPassword", err)
		return
	}

	// Gen new passwrod
	// Generate new hashed password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(*changePassword.NewPassword), DefaultGeneratePasswordCost)
	if err != nil {
		router.ResponseInternalError(w, constant.GenerateFromPasswordFaild, err.Error())
		return
	}

	// Update Password Into Database
	err = user.UpdatePassword(string(newHashedPassword))
	if err != nil {
		log.Println(log.LogLevelError, "update-new-passowrd-db", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
		return
	}

	router.ResponseSuccess(w, constant.UserUpdatePasswordSuccsess, constant.UserUpdatePasswordSuccsess)
	return
}
