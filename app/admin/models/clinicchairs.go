package models

import (
	"gorm.io/gorm"

	"go-admin/common/models"
)

type ClinicChairs struct {
	gorm.Model
	models.ControlBy

	UseStatus   string `json:"useStatus" gorm:"type:tinyint;comment:是否在使用"`      //
	Location    string `json:"location" gorm:"type:varchar(150);comment:椅子所处位置"` //
	OrderNumber string `json:"orderNumber" gorm:"type:int;comment:订单号"`          //
	Head        string `json:"head" gorm:"type:varchar(255);comment:负责人"`        //
}

func (ClinicChairs) TableName() string {
	return "clinic_chairs"
}

func (e *ClinicChairs) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ClinicChairs) GetId() interface{} {
	return e.ID
}
