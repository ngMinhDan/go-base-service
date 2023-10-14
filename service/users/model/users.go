package model

import (
	"base/pkg/db"
	"errors"
	"fmt"
)

//

// User represents the data structure for a user.
type User struct {
	ID                int     `json:"id"`
	Username          *string `json:"username"`
	Email             *string `json:"email"`
	Password          *string `json:"password"` // password_hashed
	ProfilePictureURL *string `json:"profilePictureURL"`
	Role              *string `json:"role"`
	IsActive          *bool   `json:"isActive"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         *string `json:"updatedAt"`
}

// Check Email Exist
func (user User) CheckEmailExist() (bool, error) {
	var exist bool
	query := `SELECT exists(select 1 FROM users WHERE email = $1)`
	err := db.PSQL.QueryRow(query, *user.Email).Scan(&exist)
	return exist, err
}

// Create New User
func (user User) CreateNewUser() error {
	query := `INSERT INTO users (username, email, password, role, is_active) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.PSQL.QueryRow(query, *user.Username, *user.Email, *user.Password, *user.Role, *user.IsActive).Scan(&user.ID)
	return err
}

// Update Password
func (user User) UpdatePassword(newPassword string) error {
	query := fmt.Sprintf("Update users SET password = '%s', updated_at = now() WHERE id = '%d'", newPassword, user.ID)
	_, err := db.PSQL.Query(query)
	return err
}

// Update Role User
func (user User) UpdateRole(role string) error {
	query := fmt.Sprintf("Update users SET role = '%s', updated_at = now() WHERE id = '%d'", role, user.ID)
	_, err := db.PSQL.Query(query)
	return err
}

// Update Profile : Username or Avatar
func (user User) UpdateProfile(username, avatar string) error {
	if username != "" && avatar != "" {
		query := fmt.Sprintf("Update users SET username = '%s', profile_picture_url = '%s', updated_at = now() WHERE id = '%d'",
			username, avatar, user.ID)
		_, err := db.PSQL.Query(query)
		return err
	}
	if username != "" {
		query := fmt.Sprintf("Update users SET username = '%s', updated_at = now() WHERE id = '%d'",
			username, user.ID)
		_, err := db.PSQL.Query(query)
		return err
	}
	if avatar != "" {
		query := fmt.Sprintf("Update users SET avatar = '%s', updated_at = now() WHERE id = '%d'",
			avatar, user.ID)
		_, err := db.PSQL.Query(query)
		return err
	}
	return errors.New("UpdateProfile Wrong !")
}
