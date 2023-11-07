package controller

import (
	"base/pkg/constant"
	"base/pkg/log"
	"base/pkg/router"
	"base/pkg/utils"
	"base/service/messages/model"
	"base/service/messages/repository"
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

	messages, err := repository.GetAllMessages()
	if err != nil {
		log.Println(log.LogLevelError, "get-all-messages-fail", err.Error())
		router.ResponseInternalError(w, constant.QueryDatabaseFail, err.Error())
		return
	}
	router.ResponseSuccessWithData(w, "", "", messages)

	return
}
