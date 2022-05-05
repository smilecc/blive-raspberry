package services

type Service struct {
	name string
}

type IService interface {
	Start()
}
