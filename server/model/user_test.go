package model

import ("fmt"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("User Model", func() {

	BeforeEach(func() {
		fmt.Println("Before Each function")
	})

	Context("On adding a user in the hash map", func() {
		It("successfully adding a user on map", func() {
			user:=User{Name:"jie", Email:"zan.jie.lee@gmail.com", Year:1986}
			userStore := UserStorage{prefix:"user", storage:make(map[string]User)}
			userStore.Create(user)
			res, _  := userStore.Read(user.Email)
			fmt.Println("res", res)

		})

	})

})