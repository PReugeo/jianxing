package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type PatientRecordsSearch struct {
	dto.Pagination `search:"-"`
	PatientId      string `form:"patientId" search:"type:exact;column:patient_id;table:patient_records" comment:"患者"`

	AttendingPhysician string `form:"attendingPhysician" search:"type:contains;column:attending_physician;table:patient_records" comment:"主治医生"`

	MedicalRecord string `form:"medicalRecord" search:"type:exact;column:medical_record;table:patient_records" comment:"病历号"`
}

func (m *PatientRecordsSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *PatientRecordsSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *PatientRecordsSearch) Generate() dto.Index {
	o := *m
	return &o
}

type PatientRecordsControl struct {
	ID uint `uri:"ID" comment:""` //

	PatientId string `json:"patientId" comment:"患者"`

	AttendingPhysician string `json:"attendingPhysician" comment:"主治医生"`

	MedicalRecord string `json:"medicalRecord" comment:"病历号"`

	AllergyMedication string `json:"allergyMedication" comment:"过敏药物"`

	RecentUsedDrugs string `json:"recentUsedDrugs" comment:"最近使用药物"`

	CheckPurpose string `json:"checkPurpose" comment:"检查目的"`

	RecentMedicalDate string `json:"recentMedicalDate" comment:"最近看牙日期"`

	IsDrugAllergy string `json:"isDrugAllergy" comment:"麻醉药物过敏"`

	ChronicDisease string `json:"chronicDisease" comment:"慢性病"`

	IsPregnancy string `json:"isPregnancy" comment:"怀孕"`

	SignPic string `json:"signPic" comment:"手写签名"`

	SignDate string `json:"signDate" comment:"签名日期"`

	InterrogationType string `json:"interrogationType" comment:"问诊类型"`

	PhysicianVisitsRes string `json:"physicianVisitsRes" comment:"问诊"`

	TreatmentOptionsRes string `json:"treatmentOptionsRes" comment:"治疗方案"`

	Contraindications string `json:"contraindications" comment:"禁忌症"`

	CurrentProcess string `json:"currentProcess" comment:"当前流程"`

	NextProcess string `json:"nextProcess" comment:"下一个流程"`
}

func (s *PatientRecordsControl) Bind(ctx *gin.Context) error {
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

func (s *PatientRecordsControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientRecordsControl) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientRecords{

		Model:               gorm.Model{ID: s.ID},
		PatientId:           s.PatientId,
		AttendingPhysician:  s.AttendingPhysician,
		MedicalRecord:       s.MedicalRecord,
		AllergyMedication:   s.AllergyMedication,
		RecentUsedDrugs:     s.RecentUsedDrugs,
		CheckPurpose:        s.CheckPurpose,
		RecentMedicalDate:   s.RecentMedicalDate,
		IsDrugAllergy:       s.IsDrugAllergy,
		ChronicDisease:      s.ChronicDisease,
		IsPregnancy:         s.IsPregnancy,
		SignPic:             s.SignPic,
		SignDate:            s.SignDate,
		InterrogationType:   s.InterrogationType,
		PhysicianVisitsRes:  s.PhysicianVisitsRes,
		TreatmentOptionsRes: s.TreatmentOptionsRes,
		Contraindications:   s.Contraindications,
		CurrentProcess:      s.CurrentProcess,
		NextProcess:         s.NextProcess,
	}, nil
}

func (s *PatientRecordsControl) GetId() interface{} {
	return s.ID
}

func (s *PatientRecordsControl) GetPatientId() interface{} {
	pId, err := strconv.ParseInt(s.PatientId, 10, 64)
	if err != nil {
		return nil
	}
	return pId
}

type PatientRecordsById struct {
	dto.ObjectById
}

func (s *PatientRecordsById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientRecordsById) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientRecords{}, nil
}
