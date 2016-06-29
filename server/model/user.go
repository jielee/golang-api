package model

import (
	"time"
	"github.com/learning/api/server/storage"
	"fmt"
)

type Users interface  {
	Create(u *User) (*User, error)
	Read(email string) (*User, error)
}

type User struct {
	Id string 	`json:"id"`
	Name string 	`json:"name"`
	Year int	`json:"year,string"`
	Email string	`json:"email"`
	Created time.Time`json:"time"`
}

type UserMemoryStorage struct {
	Storage storage.MemoryStorage
}

func (s *UserMemoryStorage) Create(user *User) (*User, error){
	user.Created = time.Now()
	s.Storage.Create(fmt.Sprintf("%v", user.Email), user)
	return user, nil
}

//how to transform interface to object s.Storage.GetOne(key).(*User)
func (s *UserMemoryStorage) Read(key string) (*User, error){
	return s.Storage.GetOne(key).(*User), nil
}

func NewUserMemoryStorage() UserMemoryStorage{
	return UserMemoryStorage{Storage: storage.MemoryStorage{Storage:make(map[interface{}]interface{})}}
}



type UserStorage struct {
	prefix string
	storage map[string]*User

}

func NewUserStorage(name string) UserStorage{
	return UserStorage{prefix: name , storage: make(map[string]*User)}

}

func (u *UserStorage) Create(res *User) (*User, error) {
	var err error = nil
	res.Created = time.Now()
	u.storage[res.Email]=res
	return res, err
}

func (u *UserStorage) Read(email string) (*User, error){
	var err error = nil
	res, _ := u.storage[email]
	return res, err
}


