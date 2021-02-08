package models

import (
	"gorm.io/gorm"

	"go-admin/common/models"
)

type PatientRecords struct {
	gorm.Model
	models.ControlBy

	PatientId           string `json:"patientId" gorm:"type:int;comment:患者"`                    //
	AttendingPhysician  string `json:"attendingPhysician" gorm:"type:int;comment:主治医生"`         //
	MedicalRecord       string `json:"medicalRecord" gorm:"type:bigint;comment:病历号"`            //
	AllergyMedication   string `json:"allergyMedication" gorm:"type:varchar(100);comment:过敏药物"` //
	RecentUsedDrugs     string `json:"recentUsedDrugs" gorm:"type:varchar(255);comment:最近使用药物"` //
	CheckPurpose        string `json:"checkPurpose" gorm:"type:varchar(100);comment:检查目的"`      //
	RecentMedicalDate   string `json:"recentMedicalDate" gorm:"type:date;comment:最近看牙日期"`       //
	IsDrugAllergy       string `json:"isDrugAllergy" gorm:"type:tinyint(1);comment:麻醉药物过敏"`     //
	ChronicDisease      string `json:"chronicDisease" gorm:"type:varchar(100);comment:慢性病"`     //
	IsPregnancy         string `json:"isPregnancy" gorm:"type:tinyint(1);comment:怀孕"`           //
	SignPic             string `json:"signPic" gorm:"type:varchar(255);comment:手写签名"`           //
	SignDate            string `json:"signDate" gorm:"type:date;comment:签名日期"`                  //
	InterrogationType   string `json:"interrogationType" gorm:"type:tinyint(1);comment:问诊类型"`   //
	PhysicianVisitsRes  string `json:"physicianVisitsRes" gorm:"type:text;comment:问诊"`          //
	TreatmentOptionsRes string `json:"treatmentOptionsRes" gorm:"type:text;comment:治疗方案"`       //
	Contraindications   string `json:"contraindications" gorm:"type:text;comment:禁忌症"`          //
	CurrentProcess      string `json:"currentProcess" gorm:"type:varchar(255);comment:当前流程"`    //
	NextProcess         string `json:"nextProcess" gorm:"type:varchar(255);comment:下一个流程"`      //
}

func (PatientRecords) TableName() string {
	return "patient_records"
}

func (e *PatientRecords) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PatientRecords) GetId() interface{} {
	return e.ID
}
