package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
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
	UpdateStatusOrder(requestId string, orderRequest *request.CallBackIpaymuRequest) (orderResponse response.UpdateOrderStatusResponse)
}

type OrderServiceImplementation struct {
	ConfigurationWebserver       config.Webserver
	DB                           *gorm.DB
	ConfigJwt                    config.Jwt
	Validate                     *validator.Validate
	Logger                       *logrus.Logger
	ConfigurationIpaymu          *config.Ipaymu
	OrderRepositoryInterface     mysql.OrderRepositoryInterface
	CartRepositoryInterface      mysql.CartRepositoryInterface
	UserRepositoryInterface      mysql.UserRepositoryInterface
	OrderItemRepositoryInterface mysql.OrderItemRepositoryInterface
}

func NewOrderService(
	configurationWebserver config.Webserver,
	DB *gorm.DB,
	configJwt config.Jwt,
	validate *validator.Validate,
	logger *logrus.Logger,
	configIpaymu *config.Ipaymu,
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
		ConfigurationIpaymu:          configIpaymu,
		OrderRepositoryInterface:     orderRepositoryInterface,
		CartRepositoryInterface:      cartRepositoryInterface,
		UserRepositoryInterface:      userRepositoryInterface,
		OrderItemRepositoryInterface: orderItemRepositoryInterface,
	}
}

func (service *OrderServiceImplementation) UpdateStatusOrder(requestId string, orderRequest *request.CallBackIpaymuRequest) (orderResponse response.UpdateOrderStatusResponse) {
	// validate request
	request.ValidateCallBackIpaymuRequest(service.Validate, orderRequest, requestId, service.Logger)

	order, _ := service.OrderRepositoryInterface.FindOrderByNumberOrder(service.DB, orderRequest.ReferenceId)

	if order.Id == "" {
		err := errors.New("order not found")
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"order not found"}, service.Logger)
	}

	orderEntity := &entity.Order{}
	orderEntity.OrderSatus = "Menunggu Konfirmasi"
	if orderRequest.StatusCode == 1 {
		orderEntity.PaymentStatus = "Sudah Dibayar"
	} else {
		orderEntity.PaymentStatus = "Pending"
	}

	orderEntity.PaymentSuccessAt.Time = time.Now()

	orderResult, err := service.OrderRepositoryInterface.UpdateOrderStatus(service.DB, orderRequest.ReferenceId, *orderEntity)
	exceptions.PanicIfError(err, requestId, service.Logger)
	orderResponse = response.ToUpdateOrderStatusResponse(orderResult)
	return orderResponse
}

func (service *OrderServiceImplementation) GenerateNumberOrder() (numberOrder string) {
	now := time.Now()
	orderEntity := &entity.Order{}
	for {
		rand.Seed(time.Now().Unix())
		charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		var output strings.Builder
		length := 8

		for i := 0; i < length; i++ {
			random := rand.Intn(len(charSet))
			randomChar := charSet[random]
			output.WriteString(string(randomChar))
		}

		orderEntity.NumberOrder = "ORDER/" + now.Format("20060102") + "/" + output.String()

		// Check referal code if exist
		checkNumberOrder, _ := service.OrderRepositoryInterface.FindOrderByNumberOrder(service.DB, orderEntity.NumberOrder)
		if checkNumberOrder.Id == "" {
			break
		}
	}
	return orderEntity.NumberOrder
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
	orderEntity.NumberOrder = service.GenerateNumberOrder()
	orderEntity.FullName = user.FamilyMembers.FullName
	orderEntity.Email = user.FamilyMembers.Email
	orderEntity.Address = orderRequest.Address
	orderEntity.Phone = user.FamilyMembers.Phone
	orderEntity.CourierNote = orderRequest.CourierNote
	orderEntity.TotalBill = orderRequest.TotalBill
	orderEntity.OrderSatus = "Menunggu Pembayaran"
	orderEntity.OrderedAt = time.Now()
	orderEntity.PaymentMethod = orderRequest.PaymentMethod
	orderEntity.PaymentStatus = "Belum Dibayar"
	orderEntity.PaymentByPoint = orderRequest.PaymentByPoint
	orderEntity.PaymentByCash = orderRequest.PaymentByCash
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

	orderItem, err := service.OrderItemRepositoryInterface.CreateOrderItems(tx, orderItems)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error create order"}, service.Logger, tx)

	// delete data item in cart
	errDelete := service.CartRepositoryInterface.DeleteAllProductInCartByIdUser(tx, idUser, cartItems)
	exceptions.PanicIfErrorWithRollback(errDelete, requestId, []string{"Error delete in cart"}, service.Logger, tx)

	var ipaymu_va = "0000007762212544" //your ipaymu va
	var ipaymu_key = "SANDBOXBA640645-B4FF-488B-A540-7F866791E73E-20220425110704"

	// Send request to ipaymu
	url, _ := url.Parse("https://sandbox.ipaymu.com/api/v2/payment/direct")

	postBody, _ := json.Marshal(map[string]interface{}{
		"name":           orderEntity.FullName,
		"phone":          orderEntity.Phone,
		"email":          orderEntity.Email,
		"amount":         orderEntity.TotalBill,
		"notifyUrl":      "http://117.53.44.216:9000/api/v1/order/update",
		"expired":        24,
		"expiredType":    "hours",
		"referenceId":    orderEntity.NumberOrder,
		"paymentMethod":  orderRequest.PaymentMethod,
		"paymentChannel": orderRequest.PaymentChannel,
	})

	bodyHash := sha256.Sum256([]byte(postBody))
	bodyHashToString := hex.EncodeToString(bodyHash[:])
	stringToSign := "POST:" + ipaymu_va + ":" + strings.ToLower(string(bodyHashToString)) + ":" + ipaymu_key

	h := hmac.New(sha256.New, []byte(ipaymu_key))
	h.Write([]byte(stringToSign))
	signature := hex.EncodeToString(h.Sum(nil))

	reqBody := ioutil.NopCloser(strings.NewReader(string(postBody)))

	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"va":           {ipaymu_va},
			"signature":    {signature},
		},
		Body: reqBody,
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var dataResponseIpaymu utilities.IpaymuDirectPaymentResponse

	if err := json.NewDecoder(resp.Body).Decode(&dataResponseIpaymu); err != nil {
		fmt.Println(err)
	}

	if dataResponseIpaymu.Status != 200 {
		exceptions.PanicIfErrorWithRollback(errors.New("error response ipaymu"), requestId, []string{"Error response ipaymu"}, service.Logger, tx)
	} else if dataResponseIpaymu.Status == 200 {
		commit := tx.Commit()
		exceptions.PanicIfError(commit.Error, requestId, service.Logger)
	}

	orderResponse = response.ToCreateOrderResponse(order, orderItem, dataResponseIpaymu)

	return orderResponse
}
