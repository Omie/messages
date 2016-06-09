package messages

import (
	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
	"github.com/omie/messages/api"
)

// initialize Message resource and register web service
func init() {
	log.Debug("--- initializing messages resource")
	ws := new(restful.WebService)
	ws.Path("/messages")
	initWebResource(ws)
	api.AddWebService(ws)
}

// setup endpoints
func initWebResource(ws *restful.WebService) {
	message_id := ws.PathParameter("message_id", "message id").DataType("string")

	ws.Route(ws.GET("/{message_id}").
		To(getMessage).
		Doc("get message").
		Param(message_id).
		Operation("getMessage"))

	ws.Route(ws.POST("/").
		Consumes("application/x-www-form-urlencoded").
		To(createMessage).
		Doc("create message").
		Operation("createMessage"))
}
