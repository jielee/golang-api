package controller

import (
	"github.com/learning/api/server/model"
	"github.com/learning/api/server/rest"
	"github.com/emicklei/go-restful"
)

type UsersController struct {
	rest.Getter
	rest.Poster
	rest.Deleter
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

	deleter := &rest.DeleteResource{
		Name: name,
		DeleteFunc: func(params map[string]string) bool{
			return userStorage.Delete(params["user-id"])
		},
	}

	return UsersController{Getter: getter, Poster: poster, Deleter: deleter}
}

func (u UsersController) ConfigureRoutes(ws *restful.WebService){
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

	ws.	Doc("delete").
		Route(ws.DELETE("/users/{user-id}").
		To(u.Delete).Doc("Delete user information").
		Operation("Delete User").Param(ws.PathParameter("user-id", "user id").DataType("string")))
}