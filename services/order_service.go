package services

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"github.com/tensuqiuwulu/be-service-teman-bunda/utilities"
	"gorm.io/gorm"
)

type OrderServiceInterface interface {
	CreateOrder(requestId string, idUser string, orderRequest *request.CreateOrderRequest) (orderResponse response.CreateOrderResponse)
}

type OrderServiceImplementation struct {
	ConfigurationWebserver       config.Webserver
	DB                           *gorm.DB
	ConfigJwt                    config.Jwt
	Validate                     *validator.Validate
	Logger                       *logrus.Logger
	OrderRepositoryInterface     mysql.OrderRepositoryInterface
	CartRepositoryInterface      mysql.CartRepositoryInterface
	UserRepositoryInterface      mysql.UserRepositoryInterface
	OrderItemRepositoryInterface mysql.OrderItemRepositoryInterface
}

func NewOrderService(
	configurationWebserver config.Webserver,
	DB *gorm.DB, configJwt config.Jwt,
	validate *validator.Validate,
	logger *logrus.Logger,
	orderRepositoryInterface mysql.OrderRepositoryInterface,
	cartRepositoryInterface mysql.CartRepositoryInterface,
	userRepositoryInterface mysql.UserRepositoryInterface,
	orderItemRepositoryInterface mysql.OrderItemRepositoryInterface) OrderServiceInterface {
	return &OrderServiceImplementation{
		ConfigurationWebserver:       configurationWebserver,
		DB:                           DB,
		ConfigJwt:                    configJwt,
		Validate:                     validate,
		Logger:                       logger,
		OrderRepositoryInterface:     orderRepositoryInterface,
		CartRepositoryInterface:      cartRepositoryInterface,
		UserRepositoryInterface:      userRepositoryInterface,
		OrderItemRepositoryInterface: orderItemRepositoryInterface,
	}
}

func (service *OrderServiceImplementation) CreateOrder(requestId string, idUser string, orderRequest *request.CreateOrderRequest) (orderResponse response.CreateOrderResponse) {

	// Validate request
	request.ValidateCreateOrderRequest(service.Validate, orderRequest, requestId, service.Logger)

	// Get data user
	user, _ := service.UserRepositoryInterface.FindUserById(service.DB, idUser)

	// Create Order
	tx := service.DB.Begin()
	exceptions.PanicIfError(tx.Error, requestId, service.Logger)

	orderEntity := &entity.Order{}
	orderEntity.Id = utilities.RandomUUID()
	orderEntity.IdUser = user.Id
	orderEntity.FullName = user.FamilyMembers.FullName
	orderEntity.Email = user.FamilyMembers.Email
	orderEntity.Address = orderRequest.Address
	orderEntity.Phone = user.FamilyMembers.Phone
	orderEntity.CourierNote = orderRequest.CourierNote
	orderEntity.TotalBill = orderRequest.TotalBill
	orderEntity.OrderSatus = "Menunggu"
	orderEntity.OrderedAt = time.Now()
	orderEntity.PaymentMethod = "VA"
	orderEntity.PaymentStatus = "Menunggu"
	orderEntity.PaymentByPoint = orderRequest.PaymentByPoint
	orderEntity.ShippingMethod = "Kurir"
	orderEntity.ShippingCost = orderRequest.ShippingCost
	orderEntity.ShippingStatus = "Menunggu"
	order, err := service.OrderRepositoryInterface.CreateOrder(tx, *orderEntity)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error create order"}, service.Logger, tx)

	// Get data cart
	cartItems, _ := service.CartRepositoryInterface.FindCartByIdUser(service.DB, idUser)

	// Create order items
	var totalPriceProduct float64
	var orderItems []entity.OrderItem
	for _, cartItem := range cartItems {
		orderItemEntity := &entity.OrderItem{}
		orderItemEntity.Id = utilities.RandomUUID()
		orderItemEntity.IdOrder = orderEntity.Id
		orderItemEntity.NoSku = cartItem.Product.NoSku
		orderItemEntity.ProductName = cartItem.Product.ProductName
		orderItemEntity.PictureUrl = cartItem.Product.PictureUrl
		orderItemEntity.Description = cartItem.Product.Description
		orderItemEntity.Weight = cartItem.Product.Weight
		orderItemEntity.Volume = cartItem.Product.Volume
		orderItemEntity.Qty = cartItem.Qty
		if cartItem.Product.ProductDiscount.FlagPromo == "true" {
			orderItemEntity.Price = cartItem.Product.ProductDiscount.Nominal
			totalPriceProduct = cartItem.Product.ProductDiscount.Nominal
		} else {
			orderItemEntity.Price = cartItem.Product.Price
			totalPriceProduct = cartItem.Product.Price
		}

		orderItemEntity.TotalPrice = totalPriceProduct * (float64(cartItem.Qty))
		orderItemEntity.CreatedAt = time.Now()
		orderItems = append(orderItems, *orderItemEntity)
	}

	fmt.Println("id", orderItems)

	orderItem, err := service.OrderItemRepositoryInterface.CreateOrderItems(tx, orderItems)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error create order"}, service.Logger, tx)

	// delete data item in cart
	

	commit := tx.Commit()
	exceptions.PanicIfError(commit.Error, requestId, service.Logger)
	orderResponse = response.ToCreateOrderResponse(order, orderItem)

	return orderResponse
}
