/*
Package auth: Provides JWT-based Authentication Middleware.
Package Functionality: Creates JWT tokens using RSA256, HS256, and includes a sample user payload.
File jwt.go: Contains functions for Create JWT Token, JWT Middleware for check auth and valid token

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package auth

import (
	"net/http"
	"strings"
	"time"

	// Golang implementation of JSON Web Tokens
	jwt "github.com/dgrijalva/jwt-go"

	"base/pkg/config"
	"base/pkg/constant"
	"base/pkg/crypt"
	"base/pkg/log"
	"base/pkg/router"
)

// ResGetJWT Struct
type ResGetJWT struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

// JWT Claims Data Struct
type jwtClaimsData struct {
	Data string `json:"data"`

	// JWT Standard Info
	jwt.StandardClaims
}

// JWT Function as Midleware for JWT Authorization
func JWT(next http.Handler) http.Handler {
	// Return Next HTTP Handler Function, If Authorization is Valid
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse HTTP Header Authorization
		authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		// Check HTTP Header Authorization Section
		// Authorization Section Length Should Be 2
		// The First Authorization Section Should Be "Bearer"
		if len(authHeader) != 2 || authHeader[0] != "Bearer" {
			log.Println(log.LogLevelWarn, "http-access", "unauthorized method "+r.Method+" at URI "+r.RequestURI)
			router.ResponseUnauthorized(w, "", "unauthorized method "+r.Method+" at URI "+r.RequestURI)
			return
		}

		// The Second Authorization Section Should Be The Credentials Payload
		// jwtToken: xxxx.yyyyy.zzzzz
		authPayload := authHeader[1]
		if len(authPayload) == 0 {
			router.ResponseBadRequest(w, constant.AuthPayloadFail, "Credentials Payload Not Exits")
			return
		}

		// Get Authorization Claims From JWT Token
		authClaims, err := jwtClaims(authPayload)

		if err != nil {
			router.ResponseBadRequest(w, constant.JwtClaimsFail, "JWT Token Is Wrong Or Invalid")
			return
		}

		// We Can Encrypt Claims Using RSA Encryption ***RSA****
		// If We Encrypt, We Will Decrypt Claims
		// If We don't you, you dont need !

		// claimsEncrypted, err := crypt.EncryptWithRSA(authClaims["data"].(string))
		// claimsDecrypted, err := crypt.DecryptWithRSA(string_data)

		// Set Claims to HTTP Header
		r.Header.Set(constant.JwtClaimsHeader, authClaims["data"].(string))

		// Call Next Handler Function With Current Request
		next.ServeHTTP(w, r)
	})
}

func CreateJWTToken(payload interface{}) (string, error) {
	if strings.ToLower(config.Config.GetString("JWT_ALGORITHM")) == "rsa256" {
		return CreateJWTTokenByRSA(payload)
	}
	if strings.ToLower(config.Config.GetString("JWT_ALGORITHM")) == "hs256" {
		return CreateJWTTokenByHS256(payload)
	}

	// Need Return Here
	log.Println(log.LogLevelError, "Config", "Error, Wrong Config ! Only Supported RSA AND HS")
	return "", nil
}

// CreateJWTTokenByRSA Function: to Generate JWT Token From Payload Data
func CreateJWTTokenByRSA(payload interface{}) (string, error) {

	// Convert Signing Key in Byte Format
	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM(crypt.KeyRSACfg.BytePrivate)
	if err != nil {
		return "", err
	}

	// Create JWT Token With RS256 Method And Set JWT Claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaimsData{
		payload.(string),
		jwt.StandardClaims{
			// Add Expire Time Into Token
			ExpiresAt: time.Now().Add(time.Duration(config.Config.GetInt64("JWT_EXPIRATION_TIME_HOURS")) * time.Hour).Unix(),
		},
	})

	// Generate JWT Token String With Signing Key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	// Return The JWT Token String and Error
	return tokenString, nil
}

// CreateJWTTokenByRSA Function: to Generate JWT Token From Payload Data
func CreateJWTTokenByHS256(payload interface{}) (string, error) {

	// Define the secret key (change this to your secret key)
	secretKey := []byte(config.Config.GetString("HS_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaimsData{
		payload.(string),
		jwt.StandardClaims{
			// Add Expire Time Into Token
			ExpiresAt: time.Now().Add(time.Duration(config.Config.GetInt64("JWT_EXPIRATION_TIME_HOURS")) * time.Hour).Unix(),
		},
	})

	// Sign the token with the secret key to create the digital signature
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// Return The JWT Token String and Error
	return tokenString, nil
}
