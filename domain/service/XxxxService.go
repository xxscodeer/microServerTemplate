package service

import (
	"microServerTemplate/domain/model"
	"microServerTemplate/domain/repository"
)

type XxxxService struct {
	iXxxxRepository repository.IXxxxRepository
}

type IXxxxService interface {
	AddXxxx(Xxxx model.Xxxx)error
	FindXxxxByName(XxxxName string)(model.Xxxx,error)
	UpdateXxxx(Xxxx model.Xxxx)error
	DelXxxx(XxxxName string)error
}

func (s XxxxService) AddXxxx(Xxxx model.Xxxx)error {
	return s.iXxxxRepository.CreateXxxx(Xxxx)
}

func (s XxxxService)FindXxxxByName(XxxxName string)(model.Xxxx,error)  {
	return s.iXxxxRepository.FindXxxx(XxxxName)
}
func (s XxxxService) UpdateXxxx(Xxxx model.Xxxx)error {
	return s.iXxxxRepository.UpdateXxxx(Xxxx)
}

func (s XxxxService)DelXxxx(XxxxName string)error  {
	return s.iXxxxRepository.DelXxxx(XxxxName)
}
