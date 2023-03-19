package models

type BaseResponse[I any] struct {
	Success     bool `json:"success"`
	Status_code int  `json:"status_code"`
	Items       I    `json:"items"`
}

type BaseErrorResponse struct {
	Success     bool        `json:"success"`
	Status_code int         `json:"status_code"`
	Error_msg   interface{} `json:"error_msg"`
}

func NewBaseResponse[I any](data I, code int) *BaseResponse[I] {
	return &BaseResponse[I]{
		Success:     true,
		Status_code: code,
		Items:       data,
	}
}

func NewBaseErrorResponse(err interface{}, code int) *BaseErrorResponse {
	return &BaseErrorResponse{
		Success:     false,
		Status_code: code,
		Error_msg:   err,
	}
}
