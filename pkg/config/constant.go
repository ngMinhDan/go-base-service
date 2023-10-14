package config

// Constants defined for log package
const (
	// Config Env
	DevelopmentEnv    = "development"
	ProductionEnv     = "production"
	DefaultConfigPath = "./config"
	DefaultYmlFormat  = "yml"

	// Server
	DefaultNameService = "base-service"
	DefaultIpService   = "0.0.0.0"
	DefaultPortService = "8000"

	DefaultStorePath         = "./config/stores"
	DefaultServerUploadPath  = "./config/uploads"
	DefaultServerUploadLimit = 8

	// Router
	DefaultRouteBasePath = ""
	DefaultLogLevel      = "info"

	// Key For Crypt
	DefaultPrivateKeyFile = "./config/keys/private.key"
	DefaultPublicKeyFile  = "./config/keys/public.key"

	// JWT
	DefaultJWTExpirationTime = 96 // Hours
	DefaultJWTAlgorithm      = "RSA256"

	// Log
	DefaultLogConsole = "CONSOLE"
	DefaultLogFormat  = "TEXT"

	// Cache
	DefaultEnableCache = false

	// Elasicseach
	DefaultEnableElasticseach = false
)
