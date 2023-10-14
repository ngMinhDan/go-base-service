/*
Package auth: Provides JWT-based Authentication Middleware.
Package Functionality: Creates JWT tokens using RSA256, HS256, and includes a sample user payload.
File payload.go: Contains functions Create JWT Token with sample payload : User

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package auth

import (
	"base/pkg/constant"
	"encoding/json"
)

// The information you include in Payload (JWT) depends on your usecase and the requirements of your application.
// However, Here are some info I think we should include into Payload
// UserId, Email, Username, Avatar, User's role
// Becasue we dont need get this info by DB to show for user immediately
type Payload struct {
	ID                int     `json:"id"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	ProfilePictureURL *string `json:"profilePictureURL"` // Use pointer to hanlde NULL value
	Role              string  `json:"role"`
}

type TokenData struct {
	TokenString string `json:"tokenString"`
	Type        string `json:"type"`
}

// GetDataFromClaims Function to Get Info from Claims and Unmarshal Into Payload
func (payload *Payload) GetDataFromClaims(claims string) error {
	err := json.Unmarshal([]byte(claims), &payload)
	if err != nil {
		return err
	}
	return nil
}

// CreateTokenDataJWT Function Create TokenData from Payload
func (payload *Payload) CreateTokenDataJWT() (any, error) {
	// Json
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create token from payload data
	token, err := CreateJWTToken(string(data))
	if err != nil {
		return nil, err
	}

	tokenData := TokenData{
		TokenString: token,
		Type:        constant.BearToken,
	}

	return tokenData, nil
}
