// - apps/entities/enums/order_enum.go

package enums

type PaymentMethod string

const (
	PaymentMethodCreditCard    PaymentMethod = "Credit Card"
	PaymentMethodDebitCard     PaymentMethod = "Debit Card"
	PaymentMethodBankTransfer  PaymentMethod = "Bank Transfer"
	PaymentMethodPayPal        PaymentMethod = "PayPal"
)
