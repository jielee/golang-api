package rest

import "github.com/emicklei/go-restful"

type Deleter interface {
	Delete(request *restful.Request, response *restful.Response)
}

type DeleteResource struct {
	Name string
	DeleteFunc func (id interface{}) bool
}

func (r *DeleteResource) Delete (request *restful.Request, response *restful.Response) {

}
