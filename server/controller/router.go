package controller

import "github.com/emicklei/go-restful"

type Router interface {
	ConfigureRoutes(ws *restful.WebService)
}
