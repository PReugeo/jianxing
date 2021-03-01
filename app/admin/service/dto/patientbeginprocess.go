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

type PatientBeginProcessSearch struct {
	dto.Pagination     `search:"-"`
    PatientRecordId string `form:"patientRecordId" search:"type:exact;column:patient_record_id;table:patient_begin_process" comment:"病历id"`

    CreateBy string `form:"createBy" search:"type:exact;column:create_by;table:patient_begin_process" comment:"创建人"`


}

func (m *PatientBeginProcessSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *PatientBeginProcessSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *PatientBeginProcessSearch) Generate() dto.Index {
	o := *m
	return &o
}

type PatientBeginProcessControl struct {

    ID uint `uri:"ID" comment:""` //

    PatientRecordId string `json:"patientRecordId" comment:"病历id"`


    XrayFlow string `json:"x_rayFlow" comment:"X光流程是否完成"`


    XrayBy string `json:"x_rayBy" comment:"X光流程执行人"`


    XrayTime time.Time `json:"x_rayTime" comment:"X光执行日期"`


    PhysicianVisits string `json:"physicianVisits" comment:"问诊是否完成"`


    PhysicianVisitsBy string `json:"physicianVisitsBy" comment:"问诊执行人"`


    PhysicianVisitsTime time.Time `json:"physicianVisitsTime" comment:"问诊时间"`


    TreatmentOptions string `json:"treatmentOptions" comment:"方案制定是否完成"`


    TreatmentOptionsBy string `json:"treatmentOptionsBy" comment:"方案制定人"`


    TreatmentOptionTime time.Time `json:"treatmentOptionTime" comment:"方案制定完成时间"`


    CT string `json:"cT" comment:"CT是否完成"`


    CTBy string `json:"cTBy" comment:"CT完成人"`


    CTTime time.Time `json:"cTTime" comment:"CT完成时间"`


    Toothwash string `json:"toothwash" comment:"洗牙流程是否完成"`


    ToothwashBy string `json:"toothwashBy" comment:"洗牙执行人"`


    ToothwashTime time.Time `json:"toothwashTime" comment:"洗牙完成时间"`


    TeethModulus string `json:"teethModulus" comment:"牙齿取模是否完成"`


    TeethModulusBy string `json:"teethModulusBy" comment:"牙齿取模完成人"`


    TeethModulusTime time.Time `json:"teethModulusTime" comment:"牙齿取模时间"`


    TeethCheck string `json:"teethCheck" comment:"牙齿检查是否完成"`


    TeethCheckBy string `json:"teethCheckBy" comment:"牙齿检查人"`


    TeethCheckTime time.Time `json:"teethCheckTime" comment:"牙齿检查时间"`


    TeethPhoto string `json:"teethPhoto" comment:"拍照是否完成"`


    TeethPhotoBy string `json:"teethPhotoBy" comment:"拍照执行人"`


    TeethPhotoTime time.Time `json:"teethPhotoTime" comment:"拍照时间"`


    ToothExtraction string `json:"toothExtraction" comment:"拔牙是否完成"`


    ToothExtractionBy string `json:"toothExtractionBy" comment:"拔牙执行人"`


    ToothExtractionTime time.Time `json:"toothExtractionTime" comment:"拔牙时间"`


    Leave string `json:"leave" comment:"是否已离店"`


    LeaveTime time.Time `json:"leaveTime" comment:"离开时间"`

}

func (s *PatientBeginProcessControl) Bind(ctx *gin.Context) error {
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

func (s *PatientBeginProcessControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientBeginProcessControl) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientBeginProcess{

        Model:     gorm.Model{ID: s.ID},
        PatientRecordId:  s.PatientRecordId,
        XrayFlow:  s.XrayFlow,
        XrayBy:  s.XrayBy,
        XrayTime:  s.XrayTime,
        PhysicianVisits:  s.PhysicianVisits,
        PhysicianVisitsBy:  s.PhysicianVisitsBy,
        PhysicianVisitsTime:  s.PhysicianVisitsTime,
        TreatmentOptions:  s.TreatmentOptions,
        TreatmentOptionsBy:  s.TreatmentOptionsBy,
        TreatmentOptionTime:  s.TreatmentOptionTime,
        CT:  s.CT,
        CTBy:  s.CTBy,
        CTTime:  s.CTTime,
        Toothwash:  s.Toothwash,
        ToothwashBy:  s.ToothwashBy,
        ToothwashTime:  s.ToothwashTime,
        TeethModulus:  s.TeethModulus,
        TeethModulusBy:  s.TeethModulusBy,
        TeethModulusTime:  s.TeethModulusTime,
        TeethCheck:  s.TeethCheck,
        TeethCheckBy:  s.TeethCheckBy,
        TeethCheckTime:  s.TeethCheckTime,
        TeethPhoto:  s.TeethPhoto,
        TeethPhotoBy:  s.TeethPhotoBy,
        TeethPhotoTime:  s.TeethPhotoTime,
        ToothExtraction:  s.ToothExtraction,
        ToothExtractionBy:  s.ToothExtractionBy,
        ToothExtractionTime:  s.ToothExtractionTime,
        Leave:  s.Leave,
        LeaveTime:  s.LeaveTime,
	}, nil
}

// GetId 获取病历 Id 来获取其流程
func (s *PatientBeginProcessControl) GetId() interface{} {
	return s.PatientRecordId
}

type PatientBeginProcessById struct {
	dto.ObjectById
}

func (s *PatientBeginProcessById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientBeginProcessById) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientBeginProcess{}, nil
}
