/*
* @desc:告警日志管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysAlertLog

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysAlertLog(New())
}

func New() *sSysAlertLog {
	return &sSysAlertLog{}
}

type sSysAlertLog struct {
}

// List 告警日志列表
func (s *sSysAlertLog) List(ctx context.Context, req *system.AlertLogSearchReq) (res *system.AlertLogSearchRes, err error) {
	res = new(system.AlertLogSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAlertLog.Ctx(ctx)
		m = m.LeftJoin(
			dao.SysAnalysisTask.Table(), fmt.Sprintf("%s.task_id = %s.id", dao.SysAlertLog.Table(), dao.SysAnalysisTask.Table()))
		m = m.LeftJoin(
			dao.SysAlgorithm.Table(), fmt.Sprintf("%s.en_name = %s.en_name", dao.SysAlertLog.Table(), dao.SysAlgorithm.Table()))
		m = m.LeftJoin(
			dao.SysCamera.Table(), fmt.Sprintf("%s.camera_id = %s.camera_id", dao.SysAlertLog.Table(), dao.SysCamera.Table()))

		if req != nil {
			if req.TaskId != nil {
				m = m.Where(dao.SysAlertLog.Columns().TaskId, req.TaskId)
			}
			if req.CameraId != nil {
				m = m.Where(dao.SysAlertLog.Columns().CameraId, req.CameraId)
			}
			if req.AlgorithmId != nil {
				m = m.Where(dao.SysAlertLog.Columns().AlgorithmId, req.AlgorithmId)
			}
			if req.CnName != "" {
				m = m.Where(dao.SysAlgorithm.Columns().CnName+" like ?", "%"+req.CnName+"%")
			}
			if req.EnName != "" {
				m = m.Where(dao.SysAlertLog.Columns().EnName+" like ?", "%"+req.EnName+"%")
			}
			if req.Degree != nil {
				m = m.Where(dao.SysAlertLog.Columns().Degree, req.Degree)
			}
			if req.Status != nil {
				m = m.Where(dao.SysAlertLog.Columns().Status, req.Status)
			}
			if len(req.DateRange) != 0 {
				m = m.Where(fmt.Sprintf("%s.create_at >=? AND %s.create_at <=?", dao.SysAlertLog.Table(), dao.SysAlertLog.Table()), req.DateRange[0], req.DateRange[1])
			}
		}
		m = m.Where(dao.SysAlertLog.Columns().DeleteFlag, 0)
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取告警日志失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		res.CurrentPage = req.PageNum
		err = m.
			Fields(fmt.Sprintf("%s.*,%s.alert_id as id, %s.name as task_name, %s.cn_name, %s.camera_name",
				dao.SysAlertLog.Table(),
				dao.SysAlertLog.Table(),
				dao.SysAnalysisTask.Table(),
				dao.SysAlgorithm.Table(),
				dao.SysCamera.Table())).
			Page(req.PageNum, req.PageSize).
			Order(fmt.Sprintf("%s.create_at desc", dao.SysAlertLog.Table())).
			Scan(&res.List)

		liberr.ErrIsNil(ctx, err, "获取告警日志失败")
	})
	return
}

// Add 添加告警日志
func (s *sSysAlertLog) Add(ctx context.Context, req *system.AlertLogAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err := dao.SysAlertLog.Ctx(ctx).Insert(do.SysAlertLog{
			TaskId:      req.TaskId,
			CameraId:    req.CameraId,
			AlgorithmId: req.AlgorithmId,
			EnName:      req.EnName,
			Degree:      req.Degree,
			ImageUrl:    req.ImageUrl,
			ImageName:   req.ImageName,
			Status:      req.Status,
			Remark:      req.Remark,
			CreatedBy:   service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加告警日志失败")
	})
	return
}

// Edit 修改告警日志
func (s *sSysAlertLog) Edit(ctx context.Context, req *system.AlertLogEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlertLog.Ctx(ctx).WherePri(req.AlertId).Update(do.SysAlertLog{
			TaskId:      req.TaskId,
			CameraId:    req.CameraId,
			AlgorithmId: req.AlgorithmId,
			EnName:      req.EnName,
			Degree:      req.Degree,
			ImageUrl:    req.ImageUrl,
			ImageName:   req.ImageName,
			Status:      req.Status,
			Remark:      req.Remark,
			UpdatedBy:   service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改告警日志失败")
	})
	return
}

// Delete 删除告警日志
func (s *sSysAlertLog) Delete(ctx context.Context, alertIds []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlertLog.Ctx(ctx).Where(dao.SysAlertLog.Columns().AlertId+" in(?)", alertIds).Update(do.SysAlertLog{
			DeleteFlag: 1,
		})
		liberr.ErrIsNil(ctx, err, "删除告警日志失败")
	})
	return
}

// Get 获取告警日志详情
func (s *sSysAlertLog) Get(ctx context.Context, alertId int) (res *entity.SysAlertLog, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysAlertLog.Ctx(ctx).Where(dao.SysAlertLog.Columns().AlertId, alertId).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取告警日志详情失败")
	})
	return
}

// BatchUpdateStatus 批量更新告警日志状态
func (s *sSysAlertLog) BatchUpdateStatus(ctx context.Context, alertIds []int, status int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlertLog.Ctx(ctx).
			Where(dao.SysAlertLog.Columns().AlertId+" in(?)", alertIds).
			Update(do.SysAlertLog{
				Status:    status,
				UpdatedBy: service.Context().GetUserId(ctx),
			})
		liberr.ErrIsNil(ctx, err, fmt.Sprintf("批量更新告警日志状态失败，状态码: %d", status))
	})
	return
}
