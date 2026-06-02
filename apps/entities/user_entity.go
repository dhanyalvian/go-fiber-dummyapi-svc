//- apps/entities/user_entity.go

package entities

import (
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/typesense/typesense-go/v4/typesense/api"
)

type Gender string

// 2. Define constants
const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
)

type User struct {
	BaseID

	Firstname string `json:"firstname" typesense:"index,sort"`
	Lastname  string `json:"lastname" typesense:"index,sort"`
	Gender    Gender `json:"gender" typesense:"facet,sort"`
	Avatar    string `json:"avatar" typesense:"optional"`

	Email string `json:"email" typesense:"index,sort"`
	Phone string `json:"phone" typesense:"index,sort"`

	Password     string `json:"password"`
	PasswordHash string `json:"passwordHash"`
}

type UserDoc struct {
	User
}

type RespListUser struct {
	BaseID

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Gender    Gender `json:"gender"`
}

type RespDetailUser struct {
	RespListUser

	Password     string `json:"password"`
	PasswordHash string `json:"passwordHash"`
}

func (User) ColletionName() string {
	return GetCollectionName(COLLECTION_USER)
}

func (User) TypesenseSchema() ([]api.Field, *string) {
	return utils.DeriveTypesenseFields[User](), nil
}
