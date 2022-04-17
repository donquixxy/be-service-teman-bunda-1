package response

type Response struct {
	Code  string      `json:"code"`
	Mssg  string      `json:"message"`
	Data  interface{} `json:"data"`
	Error []string    `json:"error"`
}
