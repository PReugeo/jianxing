package models

import (
    "gorm.io/gorm"
    "time"

    "go-admin/common/models"
)

type PatientBeginProcess struct {
    gorm.Model
    models.ControlBy

    PatientRecordId string `json:"patientRecordId" gorm:"type:int;comment:病历id"` //
    XrayFlow string `json:"x_rayFlow" gorm:"type:tinyint(1);comment:X光流程是否完成"` //
    XrayBy string `json:"x_rayBy" gorm:"type:varchar(100);comment:X光流程执行人"` //
    XrayTime time.Time `json:"x_rayTime" gorm:"type:datetime;comment:X光执行日期"` //
    PhysicianVisits string `json:"physicianVisits" gorm:"type:tinyint(1);comment:问诊是否完成"` //
    PhysicianVisitsBy string `json:"physicianVisitsBy" gorm:"type:varchar(100);comment:问诊执行人"` //
    PhysicianVisitsTime time.Time `json:"physicianVisitsTime" gorm:"type:datetime;comment:问诊时间"` //
    TreatmentOptions string `json:"treatmentOptions" gorm:"type:tinyint(1);comment:方案制定是否完成"` //
    TreatmentOptionsBy string `json:"treatmentOptionsBy" gorm:"type:varchar(255);comment:方案制定人"` //
    TreatmentOptionTime time.Time `json:"treatmentOptionTime" gorm:"type:datetime;comment:方案制定完成时间"` //
    CT string `json:"cT" gorm:"type:tinyint(1);comment:CT是否完成"` //
    CTBy string `json:"cTBy" gorm:"type:varchar(100);comment:CT完成人"` //
    CTTime time.Time `json:"cTTime" gorm:"type:datetime;comment:CT完成时间"` //
    Toothwash string `json:"toothwash" gorm:"type:tinyint(1);comment:洗牙流程是否完成"` //
    ToothwashBy string `json:"toothwashBy" gorm:"type:varchar(100);comment:洗牙执行人"` //
    ToothwashTime time.Time `json:"toothwashTime" gorm:"type:datetime;comment:洗牙完成时间"` //
    TeethModulus string `json:"teethModulus" gorm:"type:tinyint(1);comment:牙齿取模是否完成"` //
    TeethModulusBy string `json:"teethModulusBy" gorm:"type:varchar(100);comment:牙齿取模完成人"` //
    TeethModulusTime time.Time `json:"teethModulusTime" gorm:"type:datetime;comment:牙齿取模时间"` //
    TeethCheck string `json:"teethCheck" gorm:"type:tinyint(1);comment:牙齿检查是否完成"` //
    TeethCheckBy string `json:"teethCheckBy" gorm:"type:varchar(100);comment:牙齿检查人"` //
    TeethCheckTime time.Time `json:"teethCheckTime" gorm:"type:datetime;comment:牙齿检查时间"` //
    TeethPhoto string `json:"teethPhoto" gorm:"type:tinyint(1);comment:拍照是否完成"` //
    TeethPhotoBy string `json:"teethPhotoBy" gorm:"type:varchar(100);comment:拍照执行人"` //
    TeethPhotoTime time.Time `json:"teethPhotoTime" gorm:"type:datetime;comment:拍照时间"` //
    ToothExtraction string `json:"toothExtraction" gorm:"type:tinyint(1);comment:拔牙是否完成"` //
    ToothExtractionBy string `json:"toothExtractionBy" gorm:"type:varchar(100);comment:拔牙执行人"` //
    ToothExtractionTime time.Time `json:"toothExtractionTime" gorm:"type:datetime;comment:拔牙时间"` //
    Leave string `json:"leave" gorm:"type:tinyint(1);comment:是否已离店"` //
    LeaveTime time.Time `json:"leaveTime" gorm:"type:datetime;comment:离开时间"` //
}

func (PatientBeginProcess) TableName() string {
    return "patient_begin_process"
}

func (e *PatientBeginProcess) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PatientBeginProcess) GetId() interface{} {
	return e.ID
}
