package controller

// Sign Up Form
// I use pointer here to check nil of field in struct
type SignupForm struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`

	// You can define more to protect bot
}

// Sign In Form
// I use pointer here to check nil of field in struct
type SigninForm struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

// ChangePassword Form
// I use pointer here to check nil of field in struct
type ChangePasswordForm struct {
	Password    *string `json:"password"`
	NewPassword *string `json:"newPassword"`
}

// ChangeProfile Form
// I use pointer here to check nil of field in struct
type ChangeProfileForm struct {
	Username          *string `json:"username"`
	ProfilePictureURL *string `json:"profilePictureURL"`
}

// UpdateRole Form
type UpdateRoleForm struct {
	// Role: admin, user, mod
	Role *string `json:"role"`
}

// BlockIp Form
type BlockIpForm struct {
	IP       *string `json:"ip"`
	Duration *int    `json:"duration"` // Duration Hour
}
