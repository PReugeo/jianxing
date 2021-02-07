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

type PatientProfileSearch struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name" search:"type:contains;column:name;table:patient_profile" comment:"患者姓名"`

	Mobile string `form:"mobile" search:"type:exact;column:mobile;table:patient_profile" comment:"联系电话"`

	CreateBy string `form:"createBy" search:"type:exact;column:create_by;table:patient_profile" comment:"创建人"`
}

func (m *PatientProfileSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *PatientProfileSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *PatientProfileSearch) Generate() dto.Index {
	o := *m
	return &o
}

type PatientProfileControl struct {
	ID uint `uri:"ID" comment:"主键"` // 主键

	Name string `json:"name" comment:"患者姓名"`

	Age string `json:"age" comment:"年龄"`

	Gender string `json:"gender" comment:"患者性别"`

	Region string `json:"region" comment:"省，市，地区"`

	Nationality string `json:"nationality" comment:"民族/国籍"`

	Bitrhday string `json:"bitrhday" comment:"出生日期"`

	Mobile string `json:"mobile" comment:"联系电话"`

	Address string `json:"address" comment:"联系地址"`

	Occupation string `json:"occupation" comment:"职业"`

	Email string `json:"email" comment:"电子邮箱"`
}

func (s *PatientProfileControl) Bind(ctx *gin.Context) error {
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

func (s *PatientProfileControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientProfileControl) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientProfile{

		Model:       gorm.Model{ID: s.ID},
		Name:        s.Name,
		Age:         s.Age,
		Gender:      s.Gender,
		Region:      s.Region,
		Nationality: s.Nationality,
		Bitrhday:    s.Bitrhday,
		Mobile:      s.Mobile,
		Address:     s.Address,
		Occupation:  s.Occupation,
		Email:       s.Email,
	}, nil
}

func (s *PatientProfileControl) GetId() interface{} {
	return s.ID
}

type PatientProfileById struct {
	dto.ObjectById
}

func (s *PatientProfileById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientProfileById) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientProfile{}, nil
}
