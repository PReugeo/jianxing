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

type PatientPicSearch struct {
	dto.Pagination     `search:"-"`
    PatientId string `form:"patientId" search:"type:exact;column:patient_id;table:patient_pic" comment:"患者id"`

    CreateBy string `form:"createBy" search:"type:contains;column:create_by;table:patient_pic" comment:"创建人"`


}

func (m *PatientPicSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *PatientPicSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *PatientPicSearch) Generate() dto.Index {
	o := *m
	return &o
}

type PatientPicControl struct {

    ID uint `uri:"ID" comment:""` //

    PatientId string `json:"patientId" comment:"患者id"`


    XrayPic string `json:"x_rayPic" comment:"X光图片存储信息最新"`


    XrayTime string `json:"x_rayTime" comment:"X光拍摄时间"`


    CTPic string `json:"cTPic" comment:"CT图片最新"`


    CTTime string `json:"cTTime" comment:"CT拍摄时间"`


    TeethPic string `json:"teethPic" comment:"牙齿照片"`


    TeethTime string `json:"teethTime" comment:"牙齿拍照时间"`

}

func (s *PatientPicControl) Bind(ctx *gin.Context) error {
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

func (s *PatientPicControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientPicControl) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientPic{
        Model:     gorm.Model{ID: s.ID},
        PatientId:  s.PatientId,
        XrayPic:  s.XrayPic,
        XrayTime:  s.XrayTime,
        CTPic:  s.CTPic,
        CTTime:  s.CTTime,
        TeethPic:  s.TeethPic,
        TeethTime:  s.TeethTime,
	}, nil
}

// GetId 获取患者 ID 查找到其图片
func (s *PatientPicControl) GetId() interface{} {
	return s.PatientId
}

type PatientPicById struct {
	dto.ObjectById
}

func (s *PatientPicById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *PatientPicById) GenerateM() (common.ActiveRecord, error) {
	return &models.PatientPic{}, nil
}
