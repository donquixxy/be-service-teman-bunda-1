package services

import (
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
	orderEntity.NumberOrder = utilities.RandomUUID()
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

	orderItem, err := service.OrderItemRepositoryInterface.CreateOrderItems(tx, orderItems)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error create order"}, service.Logger, tx)

	// delete data item in cart
	errDelete := service.CartRepositoryInterface.DeleteAllProductInCartByIdUser(service.DB, idUser, cartItems)
	exceptions.PanicIfErrorWithRollback(errDelete, requestId, []string{"Error delete in cart"}, service.Logger, tx)

	commit := tx.Commit()
	exceptions.PanicIfError(commit.Error, requestId, service.Logger)

	// Send Request To Ipaymu

	// ipaymuVa := "0000007762212544"
	// ipaymuKey := "SANDBOXBA640645-B4FF-488B-A540-7F866791E73E-20220425110704"

	// url, _ := url.Parse("https://sandbox.ipaymu.com/api/v2/payment")

	// postBody, _ := json.Marshal(map[string]interface{}{
	// 	"product":     []string{"Baju"},
	// 	"qty":         []int8{1},
	// 	"price":       []float64{25000},
	// 	"returnUrl":   "http://mywebsite/thank-you-page",
	// 	"cancelUrl":   "http://mywebsite/cancel-page",
	// 	"notifyUrl":   "http://mywebsite/callback",
	// 	"referenceId": orderEntity.NumberOrder,
	// 	"buyerName":   orderEntity.FullName,
	// 	"buyerEmail":  orderEntity.Email,
	// 	"buyerPhone":  orderEntity.Phone,
	// })

	// bodyHash := sha256.Sum256([]byte(postBody))
	// bodyHashToString := hex.EncodeToString(bodyHash[:])
	// stringToSign := "POST:" + ipaymuVa + ":" + strings.ToLower(string(bodyHashToString)) + ":" + ipaymuKey

	// h := hmac.New(sha256.New, []byte(ipaymuKey))
	// h.Write([]byte(stringToSign))
	// signature := hex.EncodeToString(h.Sum(nil))

	// reqBody := ioutil.NopCloser(strings.NewReader(string(postBody)))

	// req := &http.Request{
	// 	Method: "POST",
	// 	URL:    url,
	// 	Header: map[string][]string{
	// 		"Content-Type": {"application/json"},
	// 		"va":           {ipaymuVa},
	// 		"signature":    {signature},
	// 	},
	// 	Body: reqBody,
	// }

	// resp, err := http.DefaultClient.Do(req)

	// if err != nil {
	// 	log.Fatalf("An Error Occured %v", err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// sb := string(body)
	// log.Printf(sb)

	orderResponse = response.ToCreateOrderResponse(order, orderItem)

	return orderResponse
}
