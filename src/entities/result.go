package entities

type ApiResult[T interface{}] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

func SuccessResult[T interface{}](data T) ApiResult[T] {
	return ApiResult[T]{
		Code: 0,
		Data: data,
	}
}
