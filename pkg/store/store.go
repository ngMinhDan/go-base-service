/*
Package store: Provides method to working with AWS S3
Package Functionality: This package allows you to connect, set, get file with AWS S3
Author: MinhDan <nguyenmd.works@gmail.com>
*/
package store

import (
	"base/pkg/config"
	"strings"
)

// Initialize Function in Store
func init() {
	// Store Configuration Value
	switch strings.ToLower(config.Config.GetString("STORAGE_DRIVER")) {
	case "minio":
		s3Cfg.UseSSL = config.Config.GetBool("STORAGE_USE_SSL")
		s3Cfg.Endpoint = config.Config.GetString("STORAGE_ENDPOINT")
		s3Cfg.AccessKey = config.Config.GetString("STORAGE_ACCESS_KEY")
		s3Cfg.SecretKey = config.Config.GetString("STORAGE_SECRET_KEY")
		s3Cfg.Region = config.Config.GetString("STORAGE_REGION")
		s3Cfg.Bucket = config.Config.GetString("STORAGE_BUCKET")

		s3 = s3Connect()
	}
}
