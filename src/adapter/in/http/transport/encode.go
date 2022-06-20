package transport

type Encoder struct {
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Success interface{} `json:"success"`
}

func Encode(data interface{}, err interface{}, success interface{}) *Encoder {
	return &Encoder{
		Data:    data,
		Error:   err,
		Success: success,
	}
}
