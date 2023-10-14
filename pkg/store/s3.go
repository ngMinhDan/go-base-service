/*
Package store: Provides method to working with AWS S3
Package Functionality: This package allows you to connect, set, get file with AWS S3
Author: MinhDan <nguyenmd.works@gmail.com>
*/
package store

import (
	"errors"
	"mime/multipart"
	"strconv"
	"strings"

	"base/pkg/config"
	"base/pkg/log"

	minio "github.com/minio/minio-go"
)

// StoreS3 Configuration Struct
type storeS3Config struct {
	UseSSL    bool
	Endpoint  string
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
}

// S3 Configuration Variable
var s3Cfg storeS3Config

// S3 Variable
var s3 *minio.Client

// S3 Connect Function
func s3Connect() *minio.Client {
	switch strings.ToLower(config.Config.GetString("STORAGE_DRIVER")) {
	case "minio":
		conn, err := minio.New(s3Cfg.Endpoint, s3Cfg.AccessKey, s3Cfg.SecretKey, s3Cfg.UseSSL)
		if err != nil {
			log.Println(log.LogLevelFatal, "store-s3-connect", err.Error())
		}
		return conn
	default:
		return nil
	}
}

// S3UploadFile Function to Upload File to S3 Storage
func S3UploadFile(fileName string, fileSize int64, fileType string, fileStream multipart.File) error {
	switch strings.ToLower(config.Config.GetString("STORAGE_DRIVER")) {
	case "minio":
		// Check If Bucket Exists
		bucketExists, err := s3.BucketExists(s3Cfg.Bucket)
		if err != nil {
			return err
		}

		// If Bucket Not Exists Then Create Bucket
		if !bucketExists {
			err := s3.MakeBucket(s3Cfg.Bucket, s3Cfg.Region)
			if err != nil {
				return err
			}
		}

		// Try to Upload File into Bucket
		nSize, err := s3.PutObject(s3Cfg.Bucket, fileName, fileStream, fileSize, minio.PutObjectOptions{ContentType: fileType})
		if err != nil {
			return err
		}

		log.Println(log.LogLevelInfo, "store-s3-upload-file", "successfully uploaded '"+fileName+"' with size "+strconv.FormatInt(nSize, 10))
		return nil
	default:
		return errors.New("No storage driver defined")
	}
}

// S3GetFileLink Function to Get Link for Uploaded File in S3 Storage
func S3GetFileLink(fileName string) (string, error) {
	switch strings.ToLower(config.Config.GetString("STORAGE_DRIVER")) {
	case "minio":
		if !s3Cfg.UseSSL {
			return "http://" + s3Cfg.Endpoint + "/" + s3Cfg.Bucket + "/" + fileName, nil
		}
		return "https://" + s3Cfg.Endpoint + "/" + s3Cfg.Bucket + "/" + fileName, nil
	default:
		return "", errors.New("No storage driver defined")
	}
}
