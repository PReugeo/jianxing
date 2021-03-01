package patientpic

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	common "go-admin/common/models"
	"net/http"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/log"
	"go-admin/tools"
)

type PatientPic struct {
	apis.Api
}

// @Summary 列表患者病历图片信息数据
// @Description 获取JSON
// @Tags 患者病历
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/patientpic [get]
// @Security Bearer
func (e *PatientPic) GetPatientPicList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	d := new(dto.PatientPicSearch)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	req := d.Generate()

	//查询列表
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.PatientPic, 0)
	var count int64
	serviceStudent := service.PatientPic{}
	serviceStudent.MsgID = msgID
	serviceStudent.Orm = db
	err = serviceStudent.GetPatientPicPage(req, p, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// @Summary 根据Id患者病历图片信息数据
// @Description 获取JSON
// @Tags 患者病历
// @Param id path int true "患者病历图片信息 id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/patientpic/{id} [get]
// @Security Bearer
func (e *PatientPic) GetPatientPic(c *gin.Context) {
	control := new(dto.PatientPicById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//查看详情
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.PatientPic

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	servicePatientPic := service.PatientPic{}
	servicePatientPic.MsgID = msgID
	servicePatientPic.Orm = db
	err = servicePatientPic.GetPatientPic(req, p, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, object, "查看成功")
}


// @Summary 新建患者病历图片信息数据
// @Description 获取JSON
// @Tags 患者病历
// @Accept  application/json
// @Product application/json
// @Param data models.PatientPic true "患者病历图片数据"
// @Success 200 {string} string "{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string "{"code": -1, "message": "添加失败"}"
// @Router /api/v1/patientpic [post]
// @Security Bearer
func (e *PatientPic) InsertPatientPic(c *gin.Context) {
	control := new(dto.PatientPicControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//新增操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(tools.GetUserIdUint(c))

	servicePatientPic := service.PatientPic{}
	servicePatientPic.Orm = db
	servicePatientPic.MsgID = msgID
	err = servicePatientPic.InsertPatientPic(object)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

// @Summary 修改患者病历图片数据
// @Description 获取JSON
// @Tags 患者病历
// @Accept  application/json
// @Product application/json
// @Param data body models.PatientPic true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/patientpic/{id} [put]
func (e *PatientPic) UpdatePatientPic(c *gin.Context) {
	control := new(dto.PatientPicControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	req := control.Generate()
	//更新操作
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(tools.GetUserIdUint(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	servicePatientPic := service.PatientPic{}
	servicePatientPic.Orm = db
	servicePatientPic.MsgID = msgID
	err = servicePatientPic.UpdatePatientPic(object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

// @Summary 删除患者病历图片数据
// @Description 删除数据
// @Tags 患者病历
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/patientpic/{id} [delete]
func (e *PatientPic) DeletePatientPic(c *gin.Context) {
	control := new(dto.PatientPicById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//删除操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		log.Errorf("MsgID[%s] Bind error: %s", msgID, err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}

	// 设置编辑人
	object.SetUpdateBy(tools.GetUserIdUint(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	servicePatientPic := service.PatientPic{}
	servicePatientPic.Orm = db
	servicePatientPic.MsgID = msgID
	err = servicePatientPic.RemovePatientPic(req, object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}
