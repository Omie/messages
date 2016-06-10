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

	_text, ok := request.Request.PostForm["text"]
	if !ok {
		api.ErrorResponse(response, http.StatusBadRequest, "text is required")
		return
	}
	text := _text[0]

	// insert the message into the database
	newMessage, err := message.Create(text)
	if err != nil {
		if err == message.ErrValidationError {
			api.ErrorResponse(response, http.StatusBadRequest, err.Error())
		} else {
			api.ErrorResponse(response, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response.WriteHeader(http.StatusCreated)
	// we only have to return uuid as id and not other fields of the newMessage
	// so only return that much
	response.WriteAsJson(map[string]string{"id": newMessage.Uuid})
}
