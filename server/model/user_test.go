package model

import ("fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Model", func() {

	BeforeEach(func() {
		fmt.Println("Before Each function")
	})

	Context("On adding a user in the map", func() {
		It("successfully adding a user on map", func() {
			user:=User{Name:"jie", Email:"zan.jie.lee@gmail.com", Year:1986}
			userStore := UserStorage{prefix:"user", storage:make(map[string]User)}
			userStore.Create(user)
			res, _  := userStore.Read(user.Email)
			Expect(res.Email).To(Equal(user.Email))
		})

	})

})