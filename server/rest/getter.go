package rest

import "github.com/emicklei/go-restful"

type Getter interface {
	Get(request *restful.Request, response *restful.Response)
}

type GetResource struct {
	Name string
	GetFunc func(params map[string]string) (interface{}, error)
}

func (r *GetResource) Get(request *restful.Request, response *restful.Response){
	resource, err := r.GetFunc(request.PathParameters())

	if err == nil {
		response.WriteEntity(resource)
	} else {
		response.WriteEntity(err)
	}
}
