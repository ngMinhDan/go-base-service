/*
Package config: Working with config, enviroment variable and set default value
Author: MinhDan <nguyenmd.works@gmail.com>
*/
package config

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config Variable - Can access from outside package
var Config *viper.Viper

// Initialize Function in Helper Configuration
func init() {
	// Set Configuration File Value
	configEnv := strings.ToLower(os.Getenv("CONFIG_ENV"))
	if len(configEnv) == 0 {
		configEnv = DevelopmentEnv
	}

	// Set Configuration Path Value
	configFilePath := strings.ToLower(os.Getenv("CONFIG_FILE_PATH"))
	if len(configFilePath) == 0 {
		configFilePath = DefaultConfigPath
	}

	// Set Configuration Type Value
	configFileType := strings.ToLower(os.Getenv("CONFIG_FILE_TYPE"))
	if len(configFileType) == 0 {
		configFileType = DefaultYmlFormat
	}

	// Set Configuration Prefix Value
	configPrefix := strings.ToUpper(configEnv)

	// Initialize Configuratior :
	Config = viper.New()

	// Set Configuratior Configuration
	Config.SetConfigName(configEnv)
	Config.AddConfigPath(configFilePath)
	Config.SetConfigType(configFileType)

	// Set Configurator Environment
	Config.SetEnvPrefix(configPrefix)
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set Configurator to Auto Bind Configuration Variables to
	// Environment Variables
	Config.AutomaticEnv()

	// Set Configurator to Load Configuration File
	configLoadFile()

	// Set Configurator to Set Default Value and
	// Parse Configuration Variables
	configLoadValues()
}

// ConfigLoadFile Function to Load Configuration from File
func configLoadFile() {
	// Load Configuration File
	err := Config.ReadInConfig()
	if err != nil {
		log.Println("{\"label\":\"config-load-file\",\"level\":\"warning\",\"msg\":\"error loading config file, " +
			err.Error() + "\",\"service\":\"" + Config.GetString("SERVER_NAME") +
			"\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")
	}
}

// ConfigLoadValues Function to Load Configuration Values
func configLoadValues() {
	// Server Name Value
	Config.SetDefault("SERVER_NAME", DefaultNameService)

	// Server IP Value
	Config.SetDefault("SERVER_IP", DefaultIpService)

	// Server Port Value
	Config.SetDefault("SERVER_PORT", DefaultPortService)

	// Server Store Path Value
	Config.SetDefault("SERVER_STORE_PATH", DefaultStorePath)

	// Server Upload Path Value
	Config.SetDefault("SERVER_UPLOAD_PATH", DefaultServerUploadPath)

	// Server Upload Limit Value
	Config.SetDefault("SERVER_UPLOAD_LIMIT", DefaultServerUploadLimit)
	Config.Set("SERVER_UPLOAD_LIMIT", (Config.GetInt64("SERVER_UPLOAD_LIMIT")+1)*int64(math.Pow(1024, 2)))

	// Router Base Path
	Config.SetDefault("ROUTER_BASE_PATH", DefaultRouteBasePath)

	// Server Log Level Value
	Config.SetDefault("SERVER_LOG_LEVEL", DefaultLogLevel)

	// CORS Allowed Origin Value
	Config.SetDefault("CORS_ALLOWED_ORIGIN", "*")

	// CORS Allowed Header Value
	Config.SetDefault("CORS_ALLOWED_HEADER", "*")

	// CORS Allowed Method Value
	Config.SetDefault("CORS_ALLOWED_METHOD", "*")

	// Crypt RSA Private Key File Value
	Config.SetDefault("CRYPT_PRIVATE_KEY_FILE", DefaultPrivateKeyFile)

	// Crypt RSA Public Key File Value
	Config.SetDefault("CRYPT_PUBLIC_KEY_FILE", DefaultPublicKeyFile)

	// JWT Expiration Time Value
	Config.SetDefault("JWT_EXPIRATION_TIME_HOURS", DefaultJWTExpirationTime)

	// CONFIG LOG OUTPUT DEFAULT : CONSOLE
	Config.SetDefault("LOG_OUTPUT", DefaultLogConsole)

	// CONFIG LOG FORMAT DEFAULT : TEXT
	Config.SetDefault("LOG_FORMAT", DefaultLogFormat)

	// CONFIG JWT TOKEN ALGORITHM  : RSA
	Config.SetDefault("JWT_ALGORITHM", DefaultJWTAlgorithm)

	// CONFIG ENABLE CACHE API : FALSE
	Config.SetDefault("ENABLE_CACHE_API", DefaultEnableCache)

	// CONFIG ENABLE ELASTIC SEARCH : FALSE
	Config.SetDefault("ENABLE_ELASTIC_SEARCH", DefaultEnableElasticseach)
}
