package ipaymu

type IpaymuRepositoryInterface interface {
	VaDirectPayment(paymentData interface{})
}

type IpaymuRepositoryImplementation struct {
}

func NewIpaymuRepositoryInterface() {
}

func (ipaymu *IpaymuRepositoryImplementation) VaDirectPayment(paymentData interface{}) {
}
