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
	Description string          `json:"description"`
	Stock       int             `json:"stock"`
	PictureUrl  string          `json:"picture_url"`
	Thumbnail   string          `json:"thumbnail"`
	Qty         int             `json:"qty"`
	FlagPromo   string          `json:"flag_promo"`
}

func ToFindCartByIdUser(carts []entity.Cart) (cartResponse FindCartByIdUser) {
	var cartItems []CartItem
	var totalPricePerItem decimal.Decimal
	var SubTotal decimal.Decimal
	var shippingCost decimal.Decimal
	for _, cart := range carts {
		var cartItem CartItem
		if cart.Product.ProductDiscount.FlagPromo == "true" {
			totalPricePerItem = cart.Product.ProductDiscount.Nominal
		} else {
			totalPricePerItem = cart.Product.Price
		}
		cartItem.Id = cart.Id
		cartItem.IdProduct = cart.IdProduct
		totalPricePerItem = totalPricePerItem.Mul(decimal.NewFromFloat32(float32(cart.Qty)))
		cartItem.Price = totalPricePerItem
		cartItem.ProductName = cart.Product.ProductName
		cartItem.Description = cart.Product.Description
		cartItem.Stock = cart.Product.Stock
		cartItem.PictureUrl = cart.Product.PictureUrl
		cartItem.Thumbnail = cart.Product.Thumbnail
		cartItem.Qty = cart.Qty
		cartItem.FlagPromo = cart.Product.ProductDiscount.FlagPromo
		SubTotal = SubTotal.Add(totalPricePerItem)
		cartItems = append(cartItems, cartItem)
	}

	cartResponse.CartItems = cartItems
	cartResponse.SubTotal = SubTotal
	cartResponse.TotalBill = SubTotal.Add(shippingCost)

	return cartResponse
}
