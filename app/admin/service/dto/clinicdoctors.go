package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type ClinicDoctorsSearch struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name" search:"type:contains;column:name;table:clinic_doctors" comment:"医生姓名"`

    Department string `form:"department" search:"type:exact;column:department;table:clinic_doctors" comment:"所属科室"`

    Mobile string `form:"mobile" search:"type:exact;column:mobile;table:clinic_doctors" comment:"手机号"`

    
}

func (m *ClinicDoctorsSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ClinicDoctorsSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *ClinicDoctorsSearch) Generate() dto.Index {
	o := *m
	return &o
}

type ClinicDoctorsControl struct {
    
    ID uint `uri:"ID" comment:""` // 

    Name string `json:"name" comment:"医生姓名"`
    

    Department string `json:"department" comment:"所属科室"`
    

    Mobile string `json:"mobile" comment:"手机号"`
    

    Status string `json:"status" comment:"状态"`
    
}

func (s *ClinicDoctorsControl) Bind(ctx *gin.Context) error {
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

func (s *ClinicDoctorsControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicDoctorsControl) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicDoctors{
	
        Model:     gorm.Model{ID: s.ID},
        Name:  s.Name,
        Department:  s.Department,
        Mobile:  s.Mobile,
        Status:  s.Status,
	}, nil
}

func (s *ClinicDoctorsControl) GetId() interface{} {
	return s.ID
}

type ClinicDoctorsById struct {
	dto.ObjectById
}

func (s *ClinicDoctorsById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicDoctorsById) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicDoctors{}, nil
}
