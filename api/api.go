package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

var Container *restful.Container

// initialize the container
func init() {
	ws := new(restful.WebService)
	ws.Path("/")
	Container = restful.NewContainer()
	Container.Add(ws)
}

// helper function to let controllers register their services
func AddWebService(ws *restful.WebService) {
	Container.Add(ws)
}

// helper function to emit error
func ErrorResponse(response *restful.Response, httpStatus int, err string) {
	log.Error(err)
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(httpStatus, err)
}
