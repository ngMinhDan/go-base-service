package controller

import (
	"base/pkg/constant"
	"base/pkg/db"
	"base/pkg/log"
	"base/pkg/middle"
	"base/pkg/router"
	"base/service/users/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
)

// Get List All Profiles For Admin Site
// This Is A Sample API
// In Production We Shouldn't Create Get All API
func GetAllProfile(w http.ResponseWriter, r *http.Request) {
	var reqUser model.User

	// Check Role Of User's Request
	claims := r.Header.Get(constant.JwtClaimsHeader)
	err := json.Unmarshal([]byte(claims), &reqUser)
	if err != nil {
		router.ResponseInternalError(w, constant.UnmarshalFail, err.Error())
		return
	}
	// Check Role Of User's Request
	if strings.ToLower(*reqUser.Role) != "admin" {
		router.ResponseForbiddenRequest(w, "", "")
		return
	}

	// Get List All Of User's Info
	query := "SELECT id, username, email, profile_picture_url, role, is_active, created_at, updated_at FROM users"
	rows, err := db.PSQL.Query(query)

	if err != nil {
		log.Println(log.LogLevelError, "query-get-all-users", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
		return
	}
	// Iterate through the result rows and scan into User objects
	var users []model.User
	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.ProfilePictureURL, &user.Role,
			&user.IsActive, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			log.Println(log.LogLevelError, "scan-row-to-user", err.Error())
			router.ResponseInternalError(w, constant.ScanDatabaseToObject, err.Error())
			return
		}
		users = append(users, user)
	}
	router.ResponseSuccessWithData(w, "", "", users)
	return
}

func UpdateRoleMember(w http.ResponseWriter, r *http.Request) {
	var reqUser model.User

	// Check Role Of User's Request
	claims := r.Header.Get(constant.JwtClaimsHeader)
	err := json.Unmarshal([]byte(claims), &reqUser)
	if err != nil {
		router.ResponseInternalError(w, constant.UnmarshalFail, err.Error())
		return
	}
	// Check Role Of User's Request
	if strings.ToLower(*reqUser.Role) != "admin" {
		router.ResponseForbiddenRequest(w, "", "")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		router.ResponseBadGateway(w, constant.ConvertTypeError, err.Error())
		return
	}

	var updateRoleForm UpdateRoleForm
	_ = json.NewDecoder(r.Body).Decode(&updateRoleForm)

	if updateRoleForm.Role == nil {
		router.ResponseBadRequest(w, constant.MissingFieldInputed, constant.MissingFieldInputed)
		return
	}
	if *updateRoleForm.Role != "admin" && *updateRoleForm.Role != "user" && *updateRoleForm.Role != "mod" {
		router.ResponseBadRequest(w, constant.WrongInputed, constant.WrongInputed)
		return
	}

	user := model.User{
		ID: id,
	}

	err = user.UpdateRole(*updateRoleForm.Role)
	if err != nil {
		log.Println(log.LogLevelError, "update-role-db", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
		return
	}

	// Because JWT still save old role
	// This user need logout to update jwt token

	router.ResponseSuccess(w, constant.UpdateRoleSuccsess, "")
	return
}

// Block Ip Address
func BlockIp(w http.ResponseWriter, r *http.Request) {
	var reqUser model.User
	var blockForm BlockIpForm

	// Check Role Of User's Request
	claims := r.Header.Get(constant.JwtClaimsHeader)
	err := json.Unmarshal([]byte(claims), &reqUser)
	if err != nil {
		router.ResponseInternalError(w, constant.UnmarshalFail, err.Error())
		return
	}
	// Check Role Of User's Request
	if strings.ToLower(*reqUser.Role) != "admin" {
		router.ResponseForbiddenRequest(w, "", "")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&blockForm)
	if blockForm.Duration == nil || blockForm.IP == nil {
		router.ResponseBadRequest(w, constant.MissingFieldInputed, constant.MissingFieldInputed)
		return
	}

	ip := middle.BlockedIP{
		IpAddress:  *blockForm.IP,
		UnlockTime: time.Now().Add(time.Duration(*blockForm.Duration) * time.Hour),
	}
	// Set IP Address Into Blocked IP On Cache
	err = ip.AddBlockedIp()
	if err != nil {
		router.ResponseInternalError(w, constant.SetCacheFail, err.Error())
		return
	}

	router.ResponseSuccessWithData(w, "", "", ip)
	return
}

// Check Block Ip Address
// Use middle to check blocked
func CheckBlockedIp(w http.ResponseWriter, r *http.Request) {
	router.ResponseSuccess(w, "", "You Are Not Blocked !")
}
