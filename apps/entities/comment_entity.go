//- apps/entities/comment_entity.go

package entities

import (
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/typesense/typesense-go/v4/typesense/api"
)

type Comment struct {
	BaseID

	UserID string `json:"user_id" typesense:"index"`
	PostID string `json:"post_id" typesense:"index"`
	Body   string `json:"body" typesense:"index"`
	Likes  int    `json:"likes"`

	BaseTimestamp
}

type CommentDoc struct {
	Comment
}

type RespListComment struct {
	Comment
}

type RespDetailComment struct {
	RespListComment
}

func (Comment) ColletionName() string {
	return GetCollectionName(COLLECTION_COMMENT)
}

func (Comment) TypesenseSchema() ([]api.Field, *string) {
	return utils.DeriveTypesenseFields[Comment](), nil
}
