/*
Package constant: Provides code for facilitating team development.
Package Functionality: This code defines constant variables to ensure clear and quick understanding of responses.
Both backend and frontend developers must be aware of these constants.
You can modify this code to suit your team's needs; it's a simple and flexible resource.

Author: MinhDan <nguyenmd.works@gmail.com>
*/

package constant

const (

	// Authentication Code
	UserLoginSuccess           = "LoginSuccess"
	UserLoginFailed            = "LoginFailed"
	UserNotFound               = "UserNotFound"
	UserLogoutSuccsess         = "UserLogoutSuccsess"
	UserRegisterSuccsess       = "UserRegisterSuccsess"
	UserRegisterFail           = "UserRegisterFail"
	UserUpdatePasswordSuccsess = "UserUpdatePasswordSuccsess"
	UserPasswordWrong          = "UserPasswordWrong"
	UserHasBeenBlocked         = "UserHasBeenBlocked"

	// User Profile
	UpdateProfileFail     = "UpdateProfileFail"
	UpdateProfileSuccsess = "UpdateProfileSuccsess"

	// Admin Code
	UpdateRoleSuccsess = "UpdateRoleSuccsess"
	UpdateRoleFail     = "UpdateRoleFail"

	// Database Status Code
	DatabaseConnectionFail = "DatabaseConnectionFail"
	QueryDatabaseFail      = "QueryDatabaseFail"
	ScanDatabaseToObject   = "ScanDatabaseToObject"

	// JWT
	CreateJWTTokenFail = "CreateJWTTokenFail"
	JwtClaimsFail      = "JwtClaimsFail"
	AuthPayloadFail    = "AuthPayloadFail"
	JWTIsNotValid      = "JWTIsNotValid"
	BearToken          = "Bearer"
	JwtClaimsHeader    = "X-JWT-Claims"

	// Logic Status Code
	UnmarshalFail             = "UnmarshalFail"
	MarshalFail               = "MarshalFail"
	GenerateFromPasswordFaild = "GenerateFromPasswordFaild"
	ConvertTypeError          = "ConvertTypeError"
	TooManyRequests           = "TooManyRequests"

	// User Action Input
	MissingFieldInputed = "MissingFieldInputed"
	WrongInputed        = "WrongInputed"
	EmailIsNotValid     = "EmailIsNotValid"
	EmailHostIsNotExist = "EmailHostIsNotExist"
	EmailExist          = "EmailExist"

	// Cache Code
	SetCacheFail        = "SetCacheFail"
	GetCacheFail        = "GetCacheFail"
	BlockedIpAddressKey = "BlockedIpAddressKey"
	// Store value of this key in cache 100000 hours
	BlockedIPDurationTimeToLiveHour = 100000
	CacheInternalError              = "CacheInternalError"
	CacheConnectionFail             = "CacheConnectionFail"

	// Elasticsearch Code
	ElasticSearchError    = "ElasticSearchError"
	SearchWithNotResults  = "SearchWithNotResults"
	SearchWithResults     = "SearchWithResults"
	ElasticConnectionFail = "ElasticConnectionFail"

	// Log Code
	ConsoleOuput = "console"
	JsonFormat   = "json"
)
