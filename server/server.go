package server

import (
	"github.com/emicklei/go-restful"
	"github.com/learning/api/server/model"
	"github.com/learning/api/server/controller"
	"log"
	"net/http"
)

func StartAPI(listenString string) {
	log.Fatal(http.ListenAndServe(listenString, NewApi()))
}

func NewApi() *restful.Container{

	ws := new(restful.WebService)

	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	user := model.NewUserMemoryStorage()

	controllers := []controller.Router{controller.NewUsersController(&user)}

	for _, c := range controllers{
		c.ConfigureRoutes(ws)
	}

	wsContainer := restful.NewContainer()
	wsContainer.Add(ws)

	return wsContainer

}