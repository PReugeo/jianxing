package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type PatientPic struct {
    gorm.Model
    models.ControlBy
    
    PatientId string `json:"patientId" gorm:"type:int;comment:患者id"` // 
    XrayPic string `json:"x_rayPic" gorm:"type:varchar(150);comment:X光图片存储信息最新"` // 
    XrayTime string `json:"x_rayTime" gorm:"type:date;comment:X光拍摄时间"` // 
    CTPic string `json:"cTPic" gorm:"type:varchar(255);comment:CT图片最新"` // 
    CTTime string `json:"cTTime" gorm:"type:date;comment:CT拍摄时间"` // 
    TeethPic string `json:"teethPic" gorm:"type:varchar(255);comment:牙齿照片"` // 
    TeethTime string `json:"teethTime" gorm:"type:date;comment:牙齿拍照时间"` // 
}

func (PatientPic) TableName() string {
    return "patient_pic"
}

func (e *PatientPic) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PatientPic) GetId() interface{} {
	return e.ID
}
