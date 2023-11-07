package model

import (
	"base/pkg/db"
	"errors"
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
	query := `UPDARE users SET password = $1, updated_at = now() WHERE id = $2 RETURNING id`
	err := db.PSQL.QueryRow(query, newPassword, user.ID).Scan(&user.ID)
	return err
}

// Update Role User
func (user User) UpdateRole(role string) error {
	query := `UPDATE users SET role = $1, updated_at = now() WHERE id = $2 RETURNING id`
	err := db.PSQL.QueryRow(query, role, user.ID).Scan(&user.ID)
	return err
}

// Update Profile : Username or Avatar
func (user User) UpdateProfile(username, avatar string) error {
	if username != "" && avatar != "" {
		query := `UPDATE users SET username = $1, profile_picture_url = $2, updated_at = now() WHERE id = $3 RETURNING id`
		err := db.PSQL.QueryRow(query, username, avatar, user.ID).Scan(&user.ID)
		return err
	}
	if username != "" {
		query := `UPDATE users SET username = $1, updated_at = now() WHERE id = $2 RETURNING id`
		err := db.PSQL.QueryRow(query, username, user.ID).Scan(&user.ID)
		return err
	}
	if avatar != "" {
		query := `UPDATE users SET avatar = $1, updated_at = now() WHERE id = $2 RETURNING id`
		err := db.PSQL.QueryRow(query, avatar, user.ID).Scan(&user.ID)
		return err
	}
	return errors.New("UpdateProfile Query Wrong !")
}
