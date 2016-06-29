package controller

import (
	"github.com/learning/api/server/model"
	"github.com/learning/api/server/rest"
	"github.com/emicklei/go-restful"
)

type UsersController struct {
	rest.Getter
	rest.Poster
}

func NewUsersController(userStorage model.Users) UsersController{
	name:="user"

	getter := &rest.GetResource{
		Name: name,
		GetFunc: func(params map[string]string) (interface{}, error) {
			return userStorage.Read(params["user-id"])
		},
	}

	poster := &rest.PostResource{
		Name: name,
		NewFunc: func() interface{}{
			return &model.User{}
		},
		Create: func(user interface{})(interface{}, error) {
			return userStorage.Create(user.(*model.User))
		},
	}

	return UsersController{Getter: getter, Poster: poster}
}

func (u UsersController) ConfigureOn(ws *restful.WebService){
	ws.	Doc("Users").
		Route(ws.POST("/users").
		To(u.Post).
		Doc("List").
		Operation("createUser").
		Param(ws.BodyParameter("body", "json").DataType("string")))

	ws.	Doc("User Information").
		Route(ws.GET("/users/{user-id}").To(u.Get).
		// docs
		Doc("Get user information").
		Operation("ReadUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))
}