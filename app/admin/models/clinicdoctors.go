package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type ClinicDoctors struct {
    gorm.Model
    models.ControlBy
    
    Name string `json:"name" gorm:"type:varchar(255);comment:医生姓名"` // 
    Department string `json:"department" gorm:"type:varchar(255);comment:所属科室"` // 
    Mobile string `json:"mobile" gorm:"type:varchar(255);comment:手机号"` // 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态"` // 
}

func (ClinicDoctors) TableName() string {
    return "clinic_doctors"
}

func (e *ClinicDoctors) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ClinicDoctors) GetId() interface{} {
	return e.ID
}
