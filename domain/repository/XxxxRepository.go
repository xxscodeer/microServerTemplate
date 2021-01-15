package repository

import (
	"microServerTemplate/domain/model"
	"microServerTemplate/tools"
)

//XxxxRepository 实例化数据层对象
type XxxxRepository struct {}

type IXxxxRepository interface {
	CreateXxxx(Xxxx model.Xxxx)error
	FindXxxx(XxxxName string)(model.Xxxx,error)
	UpdateXxxx(Xxxx model.Xxxx)error
	DelXxxx(XxxxName string)error
}

func (r XxxxRepository) CreateXxxx(Xxxx model.Xxxx)error {
	return tools.DbEngine.Create(Xxxx).Error
}

func (r XxxxRepository) FindXxxx(XxxxName string)(Xxxx model.Xxxx,err error) {
	return Xxxx,tools.DbEngine.Where("Xxxx_name = ?",XxxxName).Find(&Xxxx).Error
}

func (r XxxxRepository) UpdateXxxx(Xxxx model.Xxxx)error  {
	return tools.DbEngine.Model(&model.Xxxx{}).Updates(Xxxx).Error
}
func (r XxxxRepository) DelXxxx(XxxxName string)error {
	return tools.DbEngine.Delete(XxxxName).Error
}

