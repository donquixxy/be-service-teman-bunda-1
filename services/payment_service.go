package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"github.com/tensuqiuwulu/be-service-teman-bunda/utilities"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type PaymentServiceInterface interface {
	PaymentStatus(requestId string, paymentStatusRequest *request.PaymentStatusRequest) (paymentStatusResponse response.PaymentStatusResponse)
}

type PaymentServiceImplementation struct {
	ConfigWebserver                        config.Webserver
	DB                                     *gorm.DB
	Validate                               *validator.Validate
	Logger                                 *logrus.Logger
	ConfigPayment                          config.Payment
	OrderRepositoryInterface               mysql.OrderRepositoryInterface
	OrderItemRepositoryInterface           mysql.OrderItemRepositoryInterface
	ProductRepositoryInterface             mysql.ProductRepositoryInterface
	ProductStockHistoryRepositoryInterface mysql.ProductStockHistoryRepositoryInterface
	PaymentLogRepositoryInterface          mysql.PaymentLogRepositoryInterface
}

func NewPaymentService(
	configWebserver config.Webserver,
	DB *gorm.DB,
	validate *validator.Validate,
	logger *logrus.Logger,
	configPayment config.Payment,
	orderRepositoryInterface mysql.OrderRepositoryInterface,
	orderItemRepositoryInterface mysql.OrderItemRepositoryInterface,
	productRepositoryInterface mysql.ProductRepositoryInterface,
	productStockHistoryRepositoryInterface mysql.ProductStockHistoryRepositoryInterface,
	PaymentLogRepositoryInterface mysql.PaymentLogRepositoryInterface) PaymentServiceInterface {
	return &PaymentServiceImplementation{
		ConfigWebserver:                        configWebserver,
		DB:                                     DB,
		Validate:                               validate,
		Logger:                                 logger,
		ConfigPayment:                          configPayment,
		OrderRepositoryInterface:               orderRepositoryInterface,
		OrderItemRepositoryInterface:           orderItemRepositoryInterface,
		ProductRepositoryInterface:             productRepositoryInterface,
		ProductStockHistoryRepositoryInterface: productStockHistoryRepositoryInterface,
		PaymentLogRepositoryInterface:          PaymentLogRepositoryInterface,
	}
}

func (service *PaymentServiceImplementation) PaymentStatus(requestId string, paymentStatusRequest *request.PaymentStatusRequest) (paymentStatusResponse response.PaymentStatusResponse) {
	// Validate request
	request.ValidatePaymentStatusRequest(service.Validate, paymentStatusRequest, requestId, service.Logger)

	order, _ := service.OrderRepositoryInterface.FindOrderById(service.DB, paymentStatusRequest.IdOrder)

	// cek status pembayaran ke ipaymu
	var ipaymu_va = string(service.ConfigPayment.IpaymuVa)
	var ipaymu_key = string(service.ConfigPayment.IpaymuKey)

	url, _ := url.Parse(string(service.ConfigPayment.IpaymuTranscationUrl))
	postBody, _ := json.Marshal(map[string]interface{}{
		"transactionId": paymentStatusRequest.TranscationId,
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
		exceptions.PanicIfError(err, requestId, service.Logger)
	}
	defer resp.Body.Close()

	var dataPaymentStatus modelService.PaymentStatusResponse

	if err := json.NewDecoder(resp.Body).Decode(&dataPaymentStatus); err != nil {
		fmt.Println(err)
		exceptions.PanicIfError(err, requestId, service.Logger)
	}

	if order.OrderSatus == "Menunggu Pembayaran" {
		if dataPaymentStatus.Data.Status == 1 || dataPaymentStatus.Data.Status == 6 {
			tx := service.DB.Begin()

			orderEntity := &entity.Order{}
			orderEntity.OrderSatus = "Menunggu Konfirmasi"
			orderEntity.PaymentStatus = "Sudah Dibayar"
			orderEntity.PaymentSuccessAt = null.NewTime(time.Now(), true)

			_, err := service.OrderRepositoryInterface.UpdateOrderStatus(tx, order.NumberOrder, *orderEntity)
			exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error update order"}, service.Logger, tx)

			// Update product stock
			orderItems, _ := service.OrderItemRepositoryInterface.FindOrderItemsByIdOrder(service.DB, order.Id)
			for _, orderItem := range orderItems {
				productEntity := &entity.Product{}
				productEntityStockHistory := &entity.ProductStockHistory{}
				product, errFindProduct := service.ProductRepositoryInterface.FindProductById(tx, orderItem.IdProduct)
				exceptions.PanicIfErrorWithRollback(errFindProduct, requestId, []string{"product not found"}, service.Logger, tx)

				productEntityStockHistory.IdProduct = orderItem.IdProduct
				productEntityStockHistory.TxDate = time.Now()
				productEntityStockHistory.StockOpname = product.Stock
				productEntityStockHistory.StockOutQty = orderItem.Qty
				productEntityStockHistory.StockFinal = product.Stock - orderItem.Qty
				productEntityStockHistory.Description = "Pembelian " + order.NumberOrder
				productEntityStockHistory.CreatedAt = time.Now()

				_, errAddProductStockHistory := service.ProductStockHistoryRepositoryInterface.AddProductStockHistory(tx, *productEntityStockHistory)
				exceptions.PanicIfErrorWithRollback(errAddProductStockHistory, requestId, []string{"add stock history error"}, service.Logger, tx)

				productEntity.Stock = product.Stock - orderItem.Qty
				_, errUpdateProductStock := service.ProductRepositoryInterface.UpdateProductStock(tx, orderItem.IdProduct, *productEntity)
				exceptions.PanicIfErrorWithRollback(errUpdateProductStock, requestId, []string{"update stock error"}, service.Logger, tx)
			}

			// Create response log
			paymentLogEntity := &entity.PaymentLog{}
			paymentLogEntity.Id = utilities.RandomUUID()
			paymentLogEntity.IdOrder = order.Id
			paymentLogEntity.NumberOrder = order.NumberOrder
			paymentLogEntity.TypeLog = "Respon Success Ipaymu"
			paymentLogEntity.PaymentMethod = order.PaymentMethod
			paymentLogEntity.PaymentChannel = order.PaymentChannel
			paymentLogEntity.Log = fmt.Sprintf("%+v\n", dataPaymentStatus)
			paymentLogEntity.CreatedAt = time.Now()

			_, errCreateLog := service.PaymentLogRepositoryInterface.CreatePaymentLog(tx, *paymentLogEntity)
			exceptions.PanicIfErrorWithRollback(errCreateLog, requestId, []string{"Error create log"}, service.Logger, tx)

			commit := tx.Commit()
			exceptions.PanicIfError(commit.Error, requestId, service.Logger)
		}
	}

	paymentStatusResponse = response.ToPaymentStatusResponse(dataPaymentStatus)
	return paymentStatusResponse

}
