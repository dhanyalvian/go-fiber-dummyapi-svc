//- apps/dto/base_dto.go

package dto

import "go-fiber-dummyapi-svc/pkgs/date"

type ReqListFilterDateRange struct {
	Start date.Date `json:"start"`
	End   date.Date `json:"end"`
}
