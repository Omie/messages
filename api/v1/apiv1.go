package v1

import (
	"github.com/emicklei/go-restful"
	"github.com/omie/messages/api"
)

// helper function to prefix version path and add WebService in parent container
func AddWebService(ws *restful.WebService) {
	api.AddWebService(ws)
}
