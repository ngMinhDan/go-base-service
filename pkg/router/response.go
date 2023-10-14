/*
Package router: Provides a global router for handling HTTP requests.
Package Functionality: This package allows you to create a global router, set middleware for the router, and configure responses and handlers for HTTP requests.
File response.go: Provide response for HTTP request
Author: MinhDan <nguyenmd.works@gmail.com>
*/
package router

import (
	"encoding/json"
	"net/http"
)

// Response Struct: Use For Response For HTTP Request
type Response struct {
	// Code Define For Works In System
	// Bakckend And Frontend Will Define
	// To Be Clear What Happens Quickly
	Code string `json:"code"`

	// Message Response
	Message string `json:"message"`
}

// Response With Data
type ResWithData struct {
	Response
	Data any `json:"data"`
}

// Response With Error Message
type ResWithError struct {
	Response
	Error string `json:"error"`
}

// ResponseWrite Function
func ResponseWrite(w http.ResponseWriter, responseCode int, responseData any) {
	// Write Response
	w.Header().Set("Content-Type", "application/json")
	// Response With Standard HTTP Code
	w.WriteHeader(responseCode)

	// Write JSON to Response
	json.NewEncoder(w).Encode(responseData)
}

// ResponseSuccess Function
func ResponseSuccess(w http.ResponseWriter, code, message string) {
	var response Response

	// Set Default Message
	if len(message) == 0 {
		message = "Success"
	}

	// Set Default Code
	if len(code) == 0 {
		code = "200"
	}

	// Set Response Data
	response.Code = code
	response.Message = message

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusOK, response)
}

// ResponseSuccess Function With Any Data Type
func ResponseSuccessWithData(w http.ResponseWriter, code, message string, data ...any) {
	var response ResWithData
	var responseData any

	if len(data) == 1 {
		responseData = data[0]
	} else {
		responseData = data
	}
	// Set Default Message
	if len(message) == 0 {
		message = "Success"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "200"
	}

	// Set Response Data
	response.Code = code
	response.Message = message
	response.Data = responseData

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusOK, response)
}

// ResponseCreated Function With Data
func ResponseCreatedWithData(w http.ResponseWriter, code, message string, data ...any) {
	var responseData any
	var response ResWithData

	if len(data) == 1 {
		responseData = data[0]
	} else {
		responseData = data
	}

	// Set Default Message
	if len(message) == 0 {
		message = "Created successfully"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "201"
	}

	// Set Response Data
	response.Code = code
	response.Message = message
	response.Data = responseData

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusCreated, response)
}

// ResponseCreated Function: Simple Response Created With No Data
func ResponseCreated(w http.ResponseWriter, code string) {
	var response Response

	// Set Response Data
	response.Code = code
	response.Message = "Created"

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusCreated, response)
}

// ResponseUpdated Function
func ResponseUpdated(w http.ResponseWriter, code string) {
	var response Response

	// Set Response Data
	response.Code = code
	response.Message = "Updated"

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusOK, response)
}

// ResponseNoContent Function
func ResponseNoContent(w http.ResponseWriter) {
	w.WriteHeader(204)
}

// ResponseNotFound Function
func ResponseNotFound(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "Not Found"
	}
	// Set Default Message
	if len(code) == 0 {
		code = "404"
	}

	// Set Response Data
	response.Code = code
	response.Message = "Not Found"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusNotFound, response)
}

// ResponseMethodNotAllowed Function
func ResponseMethodNotAllowed(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "Method Not Allowed"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "405"
	}

	// Set Response Data
	response.Code = code
	response.Message = "Method Not Allowed"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusMethodNotAllowed, response)
}

// ResponseBadRequest Function
func ResponseBadRequest(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "Bad Request"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "400"
	}

	// Set Response Data
	response.Code = code
	response.Message = "Bad Request"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusBadRequest, response)
}

// ResponseBadRequest Function
func ResponseForbiddenRequest(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "Access denied"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "403"
	}

	// Set Response Data
	response.Code = code
	response.Message = "Access denied"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusForbidden, response)
}

// ResponseInternalError Function
func ResponseInternalError(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "Internal Server Error"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "500"
	}

	// Set Response Data
	response.Code = code
	response.Message = "Internal Server Error"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusInternalServerError, response)
}

// ResponseBadGateway Function
func ResponseBadGateway(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "Bad Gateway"
	}
	// Set Default Code
	if len(code) == 0 {
		errMessage = "502"
	}

	// Set Response Data
	response.Code = code
	response.Message = "Bad Gateway"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusBadGateway, response)
}

// ResponseUnauthorized Function
func ResponseUnauthorized(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "Unauthorized"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "401"
	}

	// Set Response Data
	response.Code = code
	response.Message = "Unauthorized"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusUnauthorized, response)
}

func ResponseTooManyRequests(w http.ResponseWriter, code, errMessage string) {
	var response ResWithError

	// Set Default Message
	if len(errMessage) == 0 {
		errMessage = "To Many Request"
	}
	// Set Default Code
	if len(code) == 0 {
		code = "429"
	}

	// Set Response Data
	response.Code = code
	response.Message = "To Many Request"
	response.Error = errMessage

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusTooManyRequests, response)
}
