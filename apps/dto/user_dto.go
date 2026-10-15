//- apps/dto/user_dto.go

package dto

type ReqListUserFilter struct {
	Gender *[]string `json:"gender"`
}
