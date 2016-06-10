package messages

import (
	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
	apiv1 "github.com/omie/messages/api/v1"
)

// initialize Message resource and register web service
func init() {
	log.Debug("--- initializing messages resource")
	ws := new(restful.WebService)
	ws.Path("/messages")
	initWebResource(ws)
	apiv1.AddWebService(ws)
}

// setup endpoints
func initWebResource(ws *restful.WebService) {
	message_id := ws.PathParameter("message_id", "message id").DataType("string")
	text := ws.PathParameter("text", "message text").DataType("string")

	ws.Route(ws.GET("/{message_id:[a-zA-Z0-9\\-]+}").
		To(getMessage).
		Doc("get message").
		Param(message_id).
		Operation("getMessage"))

	ws.Route(ws.POST("/").
		Consumes("application/x-www-form-urlencoded").
		To(createMessage).
		Doc("create message").
		Param(text).
		Operation("createMessage"))
}
