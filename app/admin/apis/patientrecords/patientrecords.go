package patientrecords

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

type PatientRecords struct {
	apis.Api
}

// @Summary 列表患者病历信息数据
// @Description 获取JSON
// @Tags 患者病历
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/patientrecords [get]
// @Security Bearer
func (e *PatientRecords) GetPatientRecordsList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	d := new(dto.PatientRecordsSearch)
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

	list := make([]models.PatientRecords, 0)
	var count int64
	serviceStudent := service.PatientRecords{}
	serviceStudent.MsgID = msgID
	serviceStudent.Orm = db
	err = serviceStudent.GetPatientRecordsPage(req, p, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// @Summary 获取患者病历信息数据
// @Description 获取JSON
// @Tags 患者病历
// @Param id path int true "患者病历 id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/patientrecords/{id} [get]
// @Security Bearer
func (e *PatientRecords) GetPatientRecords(c *gin.Context) {
	control := new(dto.PatientRecordsById)
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
	var object models.PatientRecords

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	servicePatientRecords := service.PatientRecords{}
	servicePatientRecords.MsgID = msgID
	servicePatientRecords.Orm = db
	err = servicePatientRecords.GetPatientRecords(req, p, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, object, "查看成功")
}

// @Summary 获取患者病历信息数据
// @Description 获取JSON
// @Tags 患者病历
// @Param id path int true "患者 id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/patientrecords/p/{id} [get]
// @Security Bearer
func (e *PatientRecords) GetPatientRecordsByPatientId(c *gin.Context) {
	control := new(dto.PatientRecordsControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	// 根据患者 id 查看详情
	req := *control
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.PatientRecords

	p := actions.GetPermissionFromContext(c)

	servicePatientRecords := service.PatientRecords{}
	servicePatientRecords.MsgID = msgID
	servicePatientRecords.Orm = db
	err = servicePatientRecords.GetPatientRecordsByPatientId(req, p, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(c, object, "查询成功")
}

// @Summary 新建患者病历信息数据
// @Description 获取JSON
// @Tags 患者病历
// @Accept  application/json
// @Product application/json
// @Param data models.PatientRecords true "患者病历数据"
// @Success 200 {string} string "{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string "{"code": -1, "message": "添加失败"}"
// @Router /api/v1/patientrecords [post]
// @Security Bearer
func (e *PatientRecords) InsertPatientRecords(c *gin.Context) {
	control := new(dto.PatientRecordsControl)
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

	servicePatientRecords := service.PatientRecords{}
	servicePatientRecords.Orm = db
	servicePatientRecords.MsgID = msgID
	err = servicePatientRecords.InsertPatientRecords(object)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

// @Summary 修改患者病历数据
// @Description 获取JSON
// @Tags 患者病历
// @Accept  application/json
// @Product application/json
// @Param data body models.PatientRecords true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/patientrecords/{id} [put]
func (e *PatientRecords) UpdatePatientRecords(c *gin.Context) {
	control := new(dto.PatientRecordsControl)
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

	servicePatientRecords := service.PatientRecords{}
	servicePatientRecords.Orm = db
	servicePatientRecords.MsgID = msgID
	err = servicePatientRecords.UpdatePatientRecords(object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

// @Summary 删除患者病历数据
// @Description 删除数据
// @Tags 患者病历
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/patientrecords/{id} [delete]
func (e *PatientRecords) DeletePatientRecords(c *gin.Context) {
	control := new(dto.PatientRecordsById)
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

	servicePatientRecords := service.PatientRecords{}
	servicePatientRecords.Orm = db
	servicePatientRecords.MsgID = msgID
	err = servicePatientRecords.RemovePatientRecords(req, object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}
