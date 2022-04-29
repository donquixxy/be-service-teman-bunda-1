package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindOrderByIdOrderResponse struct {
	ShippingCost   float64             `json:"shipping_cost"`
	TotalBill      float64             `json:"total_bill"`
	SubTotal       float64             `json:"sub_total"`
	OrderStatus    string              `json:"order_status"`
	PaymentByPoint float64             `json:"payment_by_point"`
	PaymentByCash  float64             `json:"payment_by_cash"`
	OrderItems     []OrderItemResponse `json:"order_items"`
}

type OrderItemResponse struct {
	Id          string  `json:"id"`
	IdProduct   string  `json:"id_product"`
	Price       float64 `json:"price"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	PictureUrl  string  `json:"picture_url"`
	Thumbnail   string  `json:"thumbnail"`
	Qty         int     `json:"qty"`
	FlagPromo   string  `json:"flag_promo"`
}

func ToFindOrderByIdOrder(order entity.Order, orderItems []entity.OrderItem) (orderResponse FindOrderByIdOrderResponse) {

	var totalPricePerItem float64
	var orderItemsResponses []OrderItemResponse
	for _, orderItem := range orderItems {
		var orderItemResponse OrderItemResponse
		orderItemResponse.Id = orderItem.Id
		orderItemResponse.IdProduct = orderItem.IdProduct
		orderItemResponse.Price = orderItem.Price
		orderItemResponse.ProductName = orderItem.ProductName
		orderItemResponse.Description = orderItem.Description
		orderItemResponse.PictureUrl = orderItem.PictureUrl
		orderItemResponse.Qty = orderItem.Qty
		orderItemResponse.Thumbnail = orderItem.Thumbnail
		orderItemResponse.FlagPromo = orderItem.FlagPromo
		totalPricePerItem = totalPricePerItem + (orderItem.Price * float64(orderItem.Qty))
		orderItemsResponses = append(orderItemsResponses, orderItemResponse)
	}

	orderResponse.OrderItems = orderItemsResponses
	orderResponse.TotalBill = order.TotalBill
	orderResponse.PaymentByPoint = order.PaymentByPoint
	orderResponse.PaymentByCash = order.PaymentByCash
	orderResponse.OrderStatus = order.OrderSatus
	orderResponse.ShippingCost = order.ShippingCost
	orderResponse.SubTotal = totalPricePerItem
	return orderResponse
}
