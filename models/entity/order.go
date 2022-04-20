package entity

type Order struct {
	Id                      string `gorm:"primaryKey;column:id;"`
	NumberOrder             string `gorm:"column:number_order;"`
	IdCustomer              string `gorm:"column:id_customer;"`
	FullName                string `gorm:"column:full_name;"`
	Email                   string `gorm:"column:email;"`
	Address                 string `gorm:"column:address;"`
	TotalBillBeforeDiscount string `gorm:"column:total_bill_before_discount;"`
	TotalBillAfterDiscount  string `gorm:"column:total_bill_after_discount;"`
	OrderSatus              string `gorm:"column:order_status;"`
	PaymentMethod           string `gorm:"column:payment_method;"`
	PaymentStatus           string `gorm:"column:payment_status;"`
	ShippingMethod          string `gorm:"column:shipping_method;"`
	ShippingCost            string `gorm:"column:shipping_cost;"`
	ShippingStatus          string `gorm:"column:shipping_status;"`
	Invoice                 string `gorm:"column:invoice;"`
	InvoiceDate             string `gorm:"column:invoice_date;"`
	PaymentDueDate          string `gorm:"column:payment_due_date;"`
}
