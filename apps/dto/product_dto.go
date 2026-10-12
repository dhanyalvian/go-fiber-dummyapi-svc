//- apps/dto/product_dto.go

package dto

type ReqListProductFilter struct {
	Brand    *[]string `json:"brand"`
	Category *[]string `json:"category"`
}
