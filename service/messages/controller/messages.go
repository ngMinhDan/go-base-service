package controller

import (
	"base/pkg/constant"
	"base/pkg/db"
	"base/pkg/log"
	"base/pkg/router"
	"base/pkg/utils"
	"base/service/messages/model"
	"encoding/json"
	"net/http"
)

// CreateMessage : Create Message From Form, save into DB and sent to message broker
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	messageForm := MessageForm{}
	_ = json.NewDecoder(r.Body).Decode(&messageForm)

	if messageForm.Content == nil {
		router.ResponseBadGateway(w, constant.MissingFieldInputed, constant.MissingFieldInputed)
		return
	}
	timeNow := utils.TimeNow()
	mess := model.Message{
		Content:   messageForm.Content,
		CreatedAt: &timeNow,
	}

	err := mess.CreateNewMessage()
	if err != nil {
		log.Println(log.LogLevelError, "query-db-message", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, constant.QueryDatabaseFail)
		return
	}

	// Send Message To Message Broker
	err = mess.ProducerMessage()
	if err != nil {
		log.Println(log.LogLevelError, "publish-message-broker-fail", err.Error())
	}

	router.ResponseCreatedWithData(w, "", "", *messageForm.Content)
	return
}

// GetAllMessage: Get All Message
func GetAllMessages(w http.ResponseWriter, r *http.Request) {

	// Get List All Of User's Info
	query := "SELECT id, content, created_at FROM messages"
	rows, err := db.PSQL.Query(query)

	if err != nil {
		log.Println(log.LogLevelError, "query-get-all-messages", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
		return
	}
	// Iterate through the result rows and scan into User objects
	var messages []model.Message
	for rows.Next() {
		var mess model.Message

		err := rows.Scan(&mess.ID, &mess.Content, &mess.CreatedAt)

		if err != nil {
			log.Println(log.LogLevelError, "scan-row-to-messages", err.Error())
			router.ResponseInternalError(w, constant.ScanDatabaseToObject, err.Error())
			return
		}
		messages = append(messages, mess)
	}
	router.ResponseSuccessWithData(w, "", "", messages)
	return
}
