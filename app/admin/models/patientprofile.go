package models

import (
	"gorm.io/gorm"

	"go-admin/common/models"
)

type PatientProfile struct {
	gorm.Model
	models.ControlBy

	Name        string `json:"name" gorm:"type:varchar(255);comment:患者姓名"`         //
	Age         string `json:"age" gorm:"type:int;comment:年龄"`                     //
	Gender      string `json:"gender" gorm:"type:varchar(20);comment:患者性别"`        //
	Region      string `json:"region" gorm:"type:varchar(255);comment:省，市，地区"`     //
	Nationality string `json:"nationality" gorm:"type:varchar(255);comment:民族/国籍"` //
	Bitrhday    string `json:"bitrhday" gorm:"type:date;comment:出生日期"`             //
	Mobile      string `json:"mobile" gorm:"type:varchar(35);comment:联系电话"`        //
	Address     string `json:"address" gorm:"type:varchar(255);comment:联系地址"`      //
	Occupation  string `json:"occupation" gorm:"type:varchar(255);comment:职业"`     //
	Email       string `json:"email" gorm:"type:varchar(255);comment:电子邮箱"`        //
}

func (PatientProfile) TableName() string {
	return "patient_profile"
}

func (e *PatientProfile) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PatientProfile) GetId() interface{} {
	return e.ID
}
