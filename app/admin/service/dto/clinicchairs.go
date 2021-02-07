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

type ClinicChairsSearch struct {
	dto.Pagination `search:"-"`
	OrderNumber    string `form:"orderNumber" search:"type:exact;column:order_number;table:clinic_chairs" comment:"订单号"`

	Head string `form:"head" search:"type:contains;column:head;table:clinic_chairs" comment:"负责人"`

	CreateBy string `form:"createBy" search:"type:exact;column:create_by;table:clinic_chairs" comment:"创建人"`
}

func (m *ClinicChairsSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ClinicChairsSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *ClinicChairsSearch) Generate() dto.Index {
	o := *m
	return &o
}

type ClinicChairsControl struct {
	ID uint `uri:"ID" comment:""` //

	UseStatus string `json:"useStatus" comment:"是否在使用"`

	Location string `json:"location" comment:"椅子所处位置"`

	OrderNumber string `json:"orderNumber" comment:"订单号"`

	Head string `json:"head" comment:"负责人"`
}

func (s *ClinicChairsControl) Bind(ctx *gin.Context) error {
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

func (s *ClinicChairsControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicChairsControl) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicChairs{

		Model:       gorm.Model{ID: s.ID},
		UseStatus:   s.UseStatus,
		Location:    s.Location,
		OrderNumber: s.OrderNumber,
		Head:        s.Head,
	}, nil
}

func (s *ClinicChairsControl) GetId() interface{} {
	return s.ID
}

type ClinicChairsById struct {
	dto.ObjectById
}

func (s *ClinicChairsById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ClinicChairsById) GenerateM() (common.ActiveRecord, error) {
	return &models.ClinicChairs{}, nil
}
