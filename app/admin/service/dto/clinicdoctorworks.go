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

type ClinicDoctorWorksSearch struct {
	dto.Pagination     `search:"-"`
    DocorId string `form:"docorId" search:"type:exact;column:docor_id;table:clinic_doctor_works" comment:"医生工号"`

    OrderNumber string `form:"orderNumber" search:"type:exact;column:order_number;table:clinic_doctor_works" comment:"订单号"`

    PatientName string `form:"patientName" search:"type:contains;column:patient_name;table:clinic_doctor_works" comment:"患者姓名"`


}

func (m *ClinicDoctorWorksSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ClinicDoctorWorksSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *ClinicDoctorWorksSearch) Generate() dto.Index {
	o := *m
	return &o
}

type ClinicDoctorWorksControl struct {

    ID uint `uri:"ID" comment:""` //

    DocorId string `json:"docorId" comment:"医生工号"`


    OrderNumber string `json:"orderNumber" comment:"订单号"`


    PatientName string `json:"patientName" comment:"患者姓名"`


    TreatmentItems string `json:"treatmentItems" comment:"治疗项目"`


    IsWork string `json:"isWork" comment:"是否已经开始"`


    TimeStart time.Time `json:"timeStart" comment:"工作开始时间"`


    TimeEnd time.Time `json:"timeEnd" comment:"工作结束时间"`

}

func (s *ClinicDoctorWorksControl) Bind(ctx *gin.Context) error {
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

func (s *ClinicDoctorWorksControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicDoctorWorksControl) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicDoctorWorks{

        Model:     gorm.Model{ID: s.ID},
        DocorId:  s.DocorId,
        OrderNumber:  s.OrderNumber,
        PatientName:  s.PatientName,
        TreatmentItems:  s.TreatmentItems,
        IsWork:  s.IsWork,
        TimeStart:  s.TimeStart,
        TimeEnd:  s.TimeEnd,
	}, nil
}

func (s *ClinicDoctorWorksControl) GetId() interface{} {
	return s.ID
}

type ClinicDoctorWorksById struct {
	dto.ObjectById
}

func (s *ClinicDoctorWorksById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicDoctorWorksById) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicDoctorWorks{}, nil
}
