package model

type Xxxx struct {
	XxxxID int64	`json:"XxxxId" gorm:"primary_key;not_null;auto_increment"`
	XxxxName string	`json:"XxxxName" gorm:"unique;not_null"`
}
