package service

import (
	"errors"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/common/service"
	"gorm.io/gorm"
)

type PatientRecords struct {
	service.Service
}

// GetPatientRecordsPage 获取PatientRecords列表
func (e *PatientRecords) GetPatientRecordsPage(c cDto.Index, p *actions.DataPermission, list *[]models.PatientRecords, count *int64) error {
	var err error
	var data models.PatientRecords
	msgID := e.MsgID

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// GetPatientRecords 获取PatientRecords对象
func (e *PatientRecords) GetPatientRecords(d cDto.Control, p *actions.DataPermission, model *models.PatientRecords) error {
	var err error
	var data models.PatientRecords
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	if db.Error != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// GetPatientRecordsByPatientId 根据 patientId 获取病历信息
func (e *PatientRecords) GetPatientRecordsByPatientId(d dto.PatientRecordsControl, p *actions.DataPermission, model *models.PatientRecords) error {
	var err error
	var data models.PatientRecords
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetPatientId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	if db.Error != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// InsertPatientRecords 创建PatientRecords对象
func (e *PatientRecords) InsertPatientRecords(model common.ActiveRecord) error {
	var err error
	var data models.PatientRecords
	msgID := e.MsgID

	err = e.Orm.Model(&data).
		Create(model).Error
	if err != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	return nil
}

// UpdatePatientRecords 修改PatientRecords对象
func (e *PatientRecords) UpdatePatientRecords(c common.ActiveRecord, p *actions.DataPermission) error {
	var err error
	var data models.PatientRecords
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Where(c.GetId()).Updates(c)
	if db.Error != nil {
		log.Errorf("msgID[%s] db error:%s", msgID, err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// RemovePatientRecords 删除PatientRecords
func (e *PatientRecords) RemovePatientRecords(d cDto.Control, c common.ActiveRecord, p *actions.DataPermission) error {
	var err error
	var data models.PatientRecords
	msgID := e.MsgID

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Where(d.GetId()).Delete(c)
	if db.Error != nil {
		err = db.Error
		log.Errorf("MsgID[%s] Delete error: %s", msgID, err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
