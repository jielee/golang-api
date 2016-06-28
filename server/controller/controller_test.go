package controller_test

import ("fmt"
	. "github.com/onsi/ginkgo"
	"github.com/learning/api/server/model"
	"github.com/learning/api/server/controller"
)

var _ = Describe("User Model", func() {

	BeforeEach(func() {
		fmt.Println("Before Each function")
	})

	Context("On adding a user in the hash map", func() {
		It("successfully adding a user on map", func() {
			userStorage := model.NewUserStorage("user")
			userController := controller.NewUsersController(userStorage)

			userController.CRUDer.Create()
		})

	})

})