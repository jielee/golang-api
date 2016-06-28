package controller

import (
	"github.com/learning/api/server/model"
	"github.com/learning/api/server/rest"
	"github.com/emicklei/go-restful"
)

type UsersController struct {
	rest.CRUDer
}

func NewUsersController(userStorage model.Users) UsersController{
	name:="user"

	cruder := &rest.ResourceCruder{
		Name:	name,

		NewFunc: func () interface{} {
			return &model.User{}
		},

		CreateRes: func(user interface{}) (interface{}, error) {
			return  userStorage.Create(user.(*model.User))
		},

		ReadRes: func(params map[string]string) (interface{}, error){
			return userStorage.Read(params["user-id"])
		},
	}
	return UsersController{CRUDer: cruder}
}

func (u UsersController) ConfigureOn(ws *restful.WebService){
	ws.	Doc("Users").
		Route(ws.POST("/users").
		To(u.Create).
		Doc("List").
		Operation("createUser").
		Param(ws.BodyParameter("body", "json").DataType("string")))

	ws.	Doc("User Information").
		Route(ws.GET("/users/{user-id}").To(u.Read).
		// docs
		Doc("Get user information").
		Operation("ReadUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))
}