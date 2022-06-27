package out

type WebResponse struct {
	Header  HeaderWebResponse  `json:"header"`
	Payload PayloadWebResponse `json:"payload"`
}

type HeaderWebResponse struct {
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

type PayloadWebResponse struct {
	Status PayloadStatusResponse `json:"status"`
	Data   interface{}           `json:"data"`
	Other  interface{}           `json:"other"`
}

type PayloadStatusResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}
