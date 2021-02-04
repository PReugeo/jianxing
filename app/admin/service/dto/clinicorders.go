package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type ClinicOrdersSearch struct {
	dto.Pagination     `search:"-"`
    PatientName string `form:"patientName" search:"type:contains;column:patient_name;table:clinic_orders" comment:"患者姓名"`

    PatientMobile string `form:"patientMobile" search:"type:exact;column:patient_mobile;table:clinic_orders" comment:"患者手机号"`

    OrderDoctor string `form:"orderDoctor" search:"type:contains;column:order_doctor;table:clinic_orders" comment:"预约医生"`


}

func (m *ClinicOrdersSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ClinicOrdersSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *ClinicOrdersSearch) Generate() dto.Index {
	o := *m
	return &o
}

type ClinicOrdersControl struct {

    ID uint `uri:"ID" comment:""` //

    OrderNumber string `json:"orderNumber" comment:"订单号"`


    PatientName string `json:"patientName" comment:"患者姓名"`


    PatientMobile string `json:"patientMobile" comment:"患者手机号"`


    PatientAge string `json:"patientAge" comment:"患者年龄"`


    Type string `json:"type" comment:"就诊情况"`


    OrderTime time.Time `json:"orderTime" comment:"预约时间"`


    OrderDoctor string `json:"orderDoctor" comment:"预约医生"`


    OrderProject string `json:"orderProject" comment:"预约项目"`


    OrderStatus string `json:"orderStatus" comment:"预约情况"`


    Remark string `json:"remark" comment:"备注"`

}

func (s *ClinicOrdersControl) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBindUri(s)
    if err != nil {
        log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
        return err
    }
    err = ctx.ShouldBind(s)
    if err != nil {
        log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
    }
    return err
}

func (s *ClinicOrdersControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicOrdersControl) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicOrders{

        Model:     gorm.Model{ID: s.ID},
        OrderNumber:  s.OrderNumber,
        PatientName:  s.PatientName,
        PatientMobile:  s.PatientMobile,
        PatientAge:  s.PatientAge,
        Type:  s.Type,
        OrderTime:  s.OrderTime,
        OrderDoctor:  s.OrderDoctor,
        OrderProject:  s.OrderProject,
        OrderStatus:  s.OrderStatus,
        Remark:  s.Remark,
	}, nil
}

func (s *ClinicOrdersControl) GetId() interface{} {
	return s.ID
}

type ClinicOrdersById struct {
	dto.ObjectById
}

func (s *ClinicOrdersById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicOrdersById) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicOrders{}, nil
}
