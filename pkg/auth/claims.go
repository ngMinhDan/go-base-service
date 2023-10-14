/*
Package auth: Provides JWT-based Authentication Middleware.
Package Functionality: Creates JWT tokens using RSA256, HS256, and includes a sample user payload.
File claims.go: Contains functions for extracting claims from a JWT token with various algorithms.

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package auth

import (
	"base/pkg/config"
	"base/pkg/constant"
	"base/pkg/crypt"
	"base/pkg/log"
	"strings"

	// Golang implementation of JSON Web Tokens
	jwt "github.com/dgrijalva/jwt-go"
)

// JWTClaimsRSA Function to Get JWT Claims Information
func jwtClaims(data string) (jwt.MapClaims, error) {
	if strings.ToLower(config.Config.GetString("JWT_ALGORITHM")) == "rsa256" {
		return jwtClaimsRSA(data)
	} else {
		// Use HS256 Algorithm
		return jwtClaimsHS(data)
	}
}

// JWTClaimsRSA Function to Get JWT Claims Information
func jwtClaimsRSA(data string) (jwt.MapClaims, error) {
	// Convert Signing Key in Byte Format
	signingKey, err := jwt.ParseRSAPublicKeyFromPEM(crypt.KeyRSACfg.BytePublic)
	if err != nil {
		return nil, err
	}

	// Parse and Validate JWT Token
	parsedToken, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		log.Println(log.LogLevelError, constant.JwtClaimsFail, err.Error())
		return nil, err
	}

	// Get The Claims
	claims := parsedToken.Claims.(jwt.MapClaims)

	// Return The Claims and Error
	return claims, err
}

// JWTClaimsHS Function to Get JWT Claims Information
func jwtClaimsHS(data string) (jwt.MapClaims, error) {
	// Convert Signing Key in Byte Format
	// Define the secret key (change this to your secret key)
	secretKey := []byte(config.Config.GetString("HS_SECRET_KEY"))

	// Parse and Validate the JWT token
	parsedToken, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// If Error Found Then Return Empty Claims and The Error
	if err != nil {
		log.Println(log.LogLevelError, constant.JwtClaimsFail, err.Error())
		return nil, err
	}

	// Get The Claims
	claims := parsedToken.Claims.(jwt.MapClaims)

	// Return The Claims and Error
	return claims, err
}

// GetJWTClaims Function: Decrypt to Get JWT Claims in Plain Text
// When use encrypt data, need decrypt data
func GetJWTClaims(data string) (string, error) {
	// Decrypt Encrypted Claims Using RSA Encryption
	claimsDecrypted, err := crypt.DecryptWithRSA(data)
	if err != nil {
		return "", err
	}

	// Return Decrypted Claims and Error
	return claimsDecrypted, nil
}
