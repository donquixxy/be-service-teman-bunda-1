package response

import (
	"github.com/shopspring/decimal"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindCartByIdUser struct {
	SubTotal     decimal.Decimal `json:"sub_total"`
	ShippingCost decimal.Decimal `json:"shipping_cost"`
	TotalBill    decimal.Decimal `json:"total_bill"`
	CartItems    []CartItem      `json:"cart_items"`
}

type CartItem struct {
	Id          string          `json:"id"`
	IdProduct   string          `json:"id_product"`
	Price       decimal.Decimal `json:"price"`
	ProductName string          `json:"product_name"`
	Stock       int             `json:"stock"`
	PictureUrl  string          `json:"picture_url"`
	Thumbnail   string          `json:"thumbnail"`
	Qty         int             `json:"qty"`
}

func ToFindCartByIdUser(carts []entity.Cart) (cartResponse FindCartByIdUser) {
	var cartItems []CartItem
	var totalPricePerItem decimal.Decimal
	var SubTotal decimal.Decimal
	var shippingCost decimal.Decimal
	for _, cart := range carts {
		var cartItem CartItem
		cartItem.Id = cart.Id
		cartItem.IdProduct = cart.IdProduct
		cartItem.Price = cart.Product.Price
		cartItem.ProductName = cart.Product.ProductName
		cartItem.Stock = cart.Product.Stock
		cartItem.PictureUrl = cart.Product.PictureUrl
		cartItem.Thumbnail = cart.Product.Thumbnail
		cartItem.Qty = cart.Qty
		totalPricePerItem = cart.Product.Price
		totalPricePerItem = totalPricePerItem.Mul(decimal.NewFromFloat32(float32(cart.Qty)))
		SubTotal = SubTotal.Add(totalPricePerItem)
		cartItems = append(cartItems, cartItem)
	}

	cartResponse.CartItems = cartItems
	cartResponse.SubTotal = SubTotal
	cartResponse.TotalBill = SubTotal.Add(shippingCost)

	return cartResponse
}
