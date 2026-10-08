//- apps/dto/order_dto.go

package dto

type ReqListOrderFilter struct {
	OrderDate     *ReqListFilterDateRange `json:"orderDate"`
	Status        *[]string               `json:"status"`
	PaymentMethod *[]string               `json:"paymentMethod"`
}
