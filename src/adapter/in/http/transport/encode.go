package transport

import (
	_ "users/docs"
)

type Encoder interface {
	Encode(interface{}, interface{}, bool) interface{}
}

type EncodedSuccess struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
} //@name ResponseSuccess

type EncodedFail struct {
	Error   interface{} `json:"error"`
	Success bool        `json:"success"`
} //@name ResponseFail

type BaseEncode struct{}

func (e *BaseEncode) Encode(data interface{}, err interface{}, success bool) interface{} {

	if err != nil {
		return &EncodedFail{
			Error:   err,
			Success: success,
		}
	}

	return &EncodedSuccess{
		Data:    data,
		Success: success,
	}
}
