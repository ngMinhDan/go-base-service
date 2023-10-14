/*
Package middle: Provides middleware for HTTP request : Check Blocked IP, Rate Limit with IP and Rate Limit For API

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package middle

import (
	"base/pkg/cache"
	"base/pkg/constant"
	"base/pkg/log"
	"base/pkg/router"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type BlockedIP struct {
	IpAddress  string    `json:"ipAddress"`
	UnlockTime time.Time `json:"unLockTime"`
}

// Add Ip Address Into Blocked IP List
// This List Will Be Saved On Redis Server
func (ip BlockedIP) AddBlockedIp() error {
	var listBlockIp []BlockedIP
	// check value of key
	byteData, check, err := cache.Redis.Get(constant.BlockedIpAddressKey)

	// If key not exist
	if check == false && err == nil {
		// Set this address into key
		listBlockIp = append(listBlockIp, ip)
		err = cache.Redis.Set(constant.BlockedIpAddressKey, listBlockIp, time.Duration(constant.BlockedIPDurationTimeToLiveHour*time.Hour))
		if err != nil {
			log.Println(log.LogLevelError, "set-value-in-cache", err.Error())
			return err
		}
	} else {
		if check == false && err != nil {
			log.Println(log.LogLevelError, "get-value-in-cache", err.Error())
			return err
		}
		// Append data into key
		err = json.Unmarshal(byteData, &listBlockIp)
		if err != nil {
			log.Println(log.LogLevelError, "json-unmarshal-cached-data", err.Error())
			return err
		}
		// append new blocked ip address into slice
		listBlockIp = append(listBlockIp, ip)

		err = cache.Redis.Set(constant.BlockedIpAddressKey, listBlockIp, time.Duration(constant.BlockedIPDurationTimeToLiveHour)*time.Hour)
		if err != nil {
			log.Println(log.LogLevelError, "set-value-in-cache", err.Error())
			return err
		}
		return nil
	}

	return nil
}

// Check Ip Address Is Blocked
// Check Ip Is Exist In Blocked List And Compare With Unlocked Time
func IsBlocked(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var listBlockedIp []BlockedIP

		// Get Current Ip Here
		currentIpAddress := strings.Split(r.RemoteAddr, ":")[0]

		// Check Ip Address In List Blocked Ip Address In Cache
		byteData, check, err := cache.Redis.Get(constant.BlockedIpAddressKey)

		// Cache Have Not Key
		if check == false && err == nil {
			next.ServeHTTP(w, r)
		} else {
			if check == false && err != nil {
				log.Println(log.LogLevelError, "get-value-in-cache", err.Error())
				router.ResponseInternalError(w, constant.GetCacheFail, err.Error())
			} else {
				// Check this ip address in list blocked ip
				_ = json.Unmarshal(byteData, &listBlockedIp)
				for _, blockedIp := range listBlockedIp {
					if blockedIp.IpAddress == currentIpAddress {
						// Check Unlock Time And Nơ
						if time.Now().Before(blockedIp.UnlockTime) {
							router.ResponseForbiddenRequest(w, constant.UserHasBeenBlocked, "You have been blocked until "+blockedIp.UnlockTime.String())
							return
						}
					}
				}
				next.ServeHTTP(w, r)
			}
		}
	})
}
