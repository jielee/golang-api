package rest

import (
	"github.com/emicklei/go-restful"
	"fmt"
	"io/ioutil"
)

type CRUDer interface {
	Create(request *restful.Request, response *restful.Response)
	Read(request *restful.Request, response *restful.Response)
	Update(request *restful.Request, response *restful.Response)
	Delete(request *restful.Request, response *restful.Response)
}

type ResourceCruder struct{
	Name string
	NewFunc func() interface{}
	CreateRes func(resource interface{}) (interface{}, error)
	ReadRes func(params map[string]string) (interface{}, error)
	//UpdateRes func(resource interface{}) (interface{}, error)
	//DeleteRes func(id interface{}) error
}

func (r *ResourceCruder) Create(request *restful.Request, response *restful.Response){

	res := r.NewFunc()

	err := request.ReadEntity(res)
	fmt.Println("RESSSSSsssss", res)
	if err != nil {
		fmt.Println("Problem reading resourceeeee from request", err)
		body, _ := ioutil.ReadAll(request.Request.Body)
		fmt.Println(string(body))
		response.WriteEntity("<h1>Hello world</h1>")
		return
	}

	res, err = r.CreateRes(res)
	if err != nil {
		fmt.Println("Problem creating resource")
	}

	response.WriteEntity(res)

}

func (r *ResourceCruder) Read (request *restful.Request, response *restful.Response){
	entity, err := r.ReadRes(request.PathParameters())

	if err == nil {
		err = response.WriteEntity(entity)
	}
}

func (r *ResourceCruder) Update (request *restful.Request, response *restful.Response){
	fmt.Println("TODO IMPLEMENTATION...")
}

func (r *ResourceCruder) Delete (request *restful.Request, response *restful.Response){
	fmt.Println("TODO IMPLEMENTATION...")
}




