package messages

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/schema"

	"github.com/omie/messages/api"
	"github.com/omie/messages/models/message"
)

var decoder = schema.NewDecoder()

// get message endpoint handler function
func getMessage(request *restful.Request, response *restful.Response) {
	log.Debug("--- getMessage")
	messageId := request.PathParameter("message_id")
	msg, err := message.GetByUUID(messageId)

	if err != nil {
		if err == message.ErrMessageNotFound {
			api.ErrorResponse(response, http.StatusNotFound, err.Error())
		} else {
			api.ErrorResponse(response, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(msg.Text))
}

// create message endpoint handler function
func createMessage(request *restful.Request, response *restful.Response) {
	log.Debug("--- createMessage")

	// parse form to retrive input Text
	err := request.Request.ParseForm()
	if err != nil {
		api.ErrorResponse(response, http.StatusBadRequest, err.Error())
		return
	}

	newMessage := new(message.Message)
	err = decoder.Decode(newMessage, request.Request.PostForm)
	if err != nil {
		api.ErrorResponse(response, http.StatusBadRequest, err.Error())
		return
	}

	// insert the message into the database
	err = newMessage.Create()
	if err != nil {
		if err == message.ErrValidationError {
			api.ErrorResponse(response, http.StatusBadRequest, err.Error())
		} else {
			api.ErrorResponse(response, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response.WriteHeader(http.StatusCreated)
	response.WriteAsJson(newMessage)
}
