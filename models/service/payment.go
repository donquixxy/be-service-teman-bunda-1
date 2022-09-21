package service

type PaymentResponse struct {
	Status  int
	Message string
	Data    Data
}

type PaymentStatusResponse struct {
	Status  int
	Data    PaymentStatus
	Message string
}

type PaymentCreditCardResponse struct {
	Status  int
	Data    CreditCardData
	Message string
}

type CreditCardData struct {
	SessionId string
	Url       string
}

type PaymentStatus struct {
	TransactionId  int
	SessionId      string
	ReferenceId    string
	RelatedId      int
	Sender         string
	Recevier       string
	Amount         float64
	Fee            float64
	Status         int
	StatusDesc     string
	Type           int
	TypeDesc       string
	Notes          string
	CreatedDate    string
	ExpiredDate    string
	SuccessDate    string
	SettlementDate string
}

type Data struct {
	SessionId     string
	TransactionId int
	ReferenceId   string
	Via           string
	Channel       string
	PaymentNo     string
	PaymentName   string
	Total         float64
	Fee           float64
	Expired       string
}

type ListPaymentChannelPayment struct {
	PaymentMethod      string
	BankCode           string
	BankName           string
	BankLogo           string
	AdminFee           float64
	AdminFeePercentage float64
}
