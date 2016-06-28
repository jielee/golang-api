package controller

import "github.com/emicklei/go-restful"

type Controller interface {
	ConfigureOn(ws *restful.WebService)
}
