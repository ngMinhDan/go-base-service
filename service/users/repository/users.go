package repository

import (
	"base/pkg/db"
	"base/pkg/log"
	"base/service/users/model"
)

// GetAllUsers : Get All Users From Database
func GetAllProfiles() ([]model.User, error) {
	query := "SELECT id, username, email, profile_picture_url, role, is_active, created_at, updated_at FROM users"
	rows, err := db.PSQL.Query(query)

	if err != nil {
		log.Println(log.LogLevelError, "query-get-all-users", err.Error())
		return nil, err
	}

	defer rows.Close()

	// Iterate through the result rows and scan into User objects
	var users []model.User
	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.ProfilePictureURL, &user.Role,
			&user.IsActive, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			log.Println(log.LogLevelError, "scan-row-to-user", err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
