package transport

type Encoder interface {
	Encode(interface{}, interface{}, interface{}) *Encoded
}

type Encoded struct {
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Success interface{} `json:"success"`
}

type BaseEncode struct{}

func (e *BaseEncode) Encode(data interface{}, err interface{}, success interface{}) *Encoded {
	return &Encoded{
		Data:    data,
		Error:   err,
		Success: success,
	}
}
