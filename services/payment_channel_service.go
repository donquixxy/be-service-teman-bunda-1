package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type PaymentChannelServiceInterface interface {
	FindListPaymentChannel(requestId string) (paymentListChannelResponses []response.FindListPaymentChannelResponse)
	FindListPaymentChannelv2(requestId string) (paymentListChannelResponses []response.FindListPaymentChannelResponse)
}

type PaymentChannelServiceImplementation struct {
	ConfigWebserver                  config.Webserver
	DB                               *gorm.DB
	Logger                           *logrus.Logger
	PaymentMethodRepositoryInterface mysql.PaymentMethodRepositoryInterface
	BankVaRepositoryInterface        mysql.BankVaRepositoryInterface
	BankTransferRepositoryInterface  mysql.BankTransferRepositoryInterface
}

func NewPaymentChannelService(
	configWebserver config.Webserver,
	DB *gorm.DB,
	logger *logrus.Logger,
	bankVaRepositoryInterface mysql.BankVaRepositoryInterface,
	bankTransferRepositoryInterface mysql.BankTransferRepositoryInterface) PaymentChannelServiceInterface {
	return &PaymentChannelServiceImplementation{
		ConfigWebserver:                 configWebserver,
		DB:                              DB,
		Logger:                          logger,
		BankVaRepositoryInterface:       bankVaRepositoryInterface,
		BankTransferRepositoryInterface: bankTransferRepositoryInterface,
	}
}

func (service *PaymentChannelServiceImplementation) FindListPaymentChannel(requestId string) (paymentListChannelResponses []response.FindListPaymentChannelResponse) {
	var listPayments []modelService.ListPaymentChannelPayment
	bankVas, _ := service.BankVaRepositoryInterface.FindAllBankVa(service.DB)
	for _, bankVa := range bankVas {
		listPayment := &modelService.ListPaymentChannelPayment{}
		listPayment.PaymentMethod = "va"
		listPayment.BankCode = bankVa.BankCode
		listPayment.BankName = bankVa.BankName
		listPayment.BankLogo = bankVa.BankLogo
		listPayment.AdminFee = bankVa.AdminFee
		listPayment.AdminFeePercentage = bankVa.AdminFeePercentage
		listPayments = append(listPayments, *listPayment)
	}

	bankTrfs, _ := service.BankTransferRepositoryInterface.FindAllBankTransfer(service.DB)
	for _, bankTrf := range bankTrfs {
		listPayment := &modelService.ListPaymentChannelPayment{}
		listPayment.PaymentMethod = "trf"
		listPayment.BankCode = bankTrf.BankCode
		listPayment.BankName = bankTrf.BankName
		listPayment.BankLogo = bankTrf.BankLogo
		listPayments = append(listPayments, *listPayment)
	}

	bankVa, _ := service.BankVaRepositoryInterface.FindBankVaByBankCode(service.DB, "qris")
	if bankVa.Id != "" {
		listPayment := &modelService.ListPaymentChannelPayment{}
		listPayment.PaymentMethod = "qris"
		listPayment.BankCode = bankVa.BankCode
		listPayment.BankName = bankVa.BankName
		listPayment.BankLogo = bankVa.BankLogo
		listPayment.AdminFee = bankVa.AdminFee
		listPayment.AdminFeePercentage = bankVa.AdminFeePercentage
		listPayments = append(listPayments, *listPayment)
	}

	bankCod, _ := service.BankVaRepositoryInterface.FindBankVaByBankCode(service.DB, "cod")
	if bankCod.Id != "" {
		listPayment := &modelService.ListPaymentChannelPayment{}
		listPayment.PaymentMethod = "cod"
		listPayment.BankCode = bankCod.BankCode
		listPayment.BankName = bankCod.BankName
		listPayment.BankLogo = bankCod.BankLogo
		listPayments = append(listPayments, *listPayment)
	}

	cc, _ := service.BankVaRepositoryInterface.FindBankVaByBankCode(service.DB, "cc")
	if cc.Id != "" {
		listPayment := &modelService.ListPaymentChannelPayment{}
		listPayment.PaymentMethod = "cc"
		listPayment.BankCode = cc.BankCode
		listPayment.BankName = cc.BankName
		listPayment.BankLogo = cc.BankLogo
		listPayment.AdminFee = cc.AdminFee
		listPayment.AdminFeePercentage = cc.AdminFeePercentage
		listPayments = append(listPayments, *listPayment)
	}

	paymentListChannelResponses = response.ToFindPaymentMethodResponses(listPayments)
	return paymentListChannelResponses
}

func (service *PaymentChannelServiceImplementation) FindListPaymentChannelv2(requestId string) (paymentListChannelResponses []response.FindListPaymentChannelResponse) {
	return

}
