package models

import (
    "gorm.io/gorm"
	"time"

	"go-admin/common/models"
)

type ClinicDoctorWorks struct {
    gorm.Model
    models.ControlBy

    DocorId string `json:"docorId" gorm:"type:int;comment:医生工号"` //
    OrderNumber string `json:"orderNumber" gorm:"type:int;comment:订单号"` //
    PatientName string `json:"patientName" gorm:"type:varchar(100);comment:患者姓名"` //
    TreatmentItems string `json:"treatmentItems" gorm:"type:varchar(255);comment:治疗项目"` //
    IsWork string `json:"isWork" gorm:"type:tinyint(1);comment:是否已经开始"` //
    TimeStart time.Time `json:"timeStart" gorm:"type:datetime;comment:工作开始时间"` //
    TimeEnd time.Time `json:"timeEnd" gorm:"type:datetime;comment:工作结束时间"` //
}

func (ClinicDoctorWorks) TableName() string {
    return "clinic_doctor_works"
}

func (e *ClinicDoctorWorks) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ClinicDoctorWorks) GetId() interface{} {
	return e.ID
}
