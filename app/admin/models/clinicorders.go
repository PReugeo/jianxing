package models

import (
    "gorm.io/gorm"
	"time"

	"go-admin/common/models"
)

type ClinicOrders struct {
    gorm.Model
    models.ControlBy

    OrderNumber string `json:"orderNumber" gorm:"type:int;comment:订单号"` //
    PatientName string `json:"patientName" gorm:"type:varchar(100);comment:患者姓名"` //
    PatientMobile string `json:"patientMobile" gorm:"type:varchar(50);comment:患者手机号"` //
    PatientAge string `json:"patientAge" gorm:"type:varchar(20);comment:患者年龄"` //
    Type string `json:"type" gorm:"type:tinyint(1);comment:就诊情况"` //
    OrderTime time.Time `json:"orderTime" gorm:"type:datetime;comment:预约时间"` //
    OrderDoctor string `json:"orderDoctor" gorm:"type:varchar(50);comment:预约医生"` //
    OrderProject string `json:"orderProject" gorm:"type:varchar(100);comment:预约项目"` //
    OrderStatus string `json:"orderStatus" gorm:"type:varchar(25);comment:预约情况"` //
    Remark string `json:"remark" gorm:"type:text;comment:备注"` //
}

func (ClinicOrders) TableName() string {
    return "clinic_orders"
}

func (e *ClinicOrders) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ClinicOrders) GetId() interface{} {
	return e.ID
}
