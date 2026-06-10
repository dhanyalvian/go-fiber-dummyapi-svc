//- apps/entities/post_entity.go

package entities

import (
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/typesense/typesense-go/v4/typesense/api"
)

type Post struct {
	BaseID

	UserID   string   `json:"user_id" typesense:"index"`
	Title    string   `json:"title" typesense:"index"`
	Body     string   `json:"body" typesense:"index"`
	Tags     []string `json:"tags" typesense:"index"`
	Likes    int      `json:"likes"`
	Dislikes int      `json:"dislikes"`
	Views    int      `json:"views"`

	BaseTimestamp
}

type PostDoc struct {
	Post
}

type RespListPost struct {
	Post
}

type RespDetailPost struct {
	RespListPost
}

func (Post) ColletionName() string {
	return GetCollectionName(COLLECTION_POST)
}

func (Post) TypesenseSchema() ([]api.Field, *string) {
	return utils.DeriveTypesenseFields[Post](), nil
}
