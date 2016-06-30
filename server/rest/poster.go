package rest

import (
	"github.com/emicklei/go-restful"
	"fmt"
)

type Poster interface {
	Post(request *restful.Request, response *restful.Response)
}

type PostResource struct {
	Name string
	NewFunc func() interface{}
	Create func(interface{}) (interface{}, error)
}

func (p *PostResource) Post(request *restful.Request, response *restful.Response){

	resource := p.NewFunc()

	err := request.ReadEntity(&resource)
	if err != nil{
		fmt.Println("TODO: IMPLEMENT ERROR HANDLER")
		return
	}

	resource, err = p.Create(resource)
	if err != nil {
		fmt.Printf("TODO: IMPLEMENT ERROR HANDLER")
	}

	response.WriteEntity(resource)
}
