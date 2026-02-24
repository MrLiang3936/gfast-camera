/*
* @desc:告警日志管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/7 23:12
 */

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var AlertLog = alertLogController{}

type alertLogController struct {
	BaseController
}

// List 告警日志列表
func (c *alertLogController) List(ctx context.Context, req *system.AlertLogSearchReq) (res *system.AlertLogSearchRes, err error) {
	res, err = service.SysAlertLog().List(ctx, req)
	return
}

// Add 添加告警日志
func (c *alertLogController) Add(ctx context.Context, req *system.AlertLogAddReq) (res *system.AlertLogAddRes, err error) {
	err = service.SysAlertLog().Add(ctx, req)
	return
}

// Edit 修改告警日志
func (c *alertLogController) Edit(ctx context.Context, req *system.AlertLogEditReq) (res *system.AlertLogEditRes, err error) {
	err = service.SysAlertLog().Edit(ctx, req)
	return
}

// Delete 删除告警日志
func (c *alertLogController) Delete(ctx context.Context, req *system.AlertLogDeleteReq) (res *system.AlertLogDeleteRes, err error) {
	err = service.SysAlertLog().Delete(ctx, req.AlertIds)
	return
}

// Get 查询告警日志详情
func (c *alertLogController) Get(ctx context.Context, req *system.AlertLogGetReq) (res *system.AlertLogGetRes, err error) {
	data, err := service.SysAlertLog().Get(ctx, req.AlertId)
	if err != nil {
		return nil, err
	}
	res = &system.AlertLogGetRes{
		Data: data,
	}
	return
}

// BatchUpdateStatus 批量更新告警日志状态
func (c *alertLogController) BatchUpdateStatus(ctx context.Context, req *system.AlertLogBatchUpdateStatusReq) (res *system.AlertLogBatchUpdateStatusRes, err error) {
	err = service.SysAlertLog().BatchUpdateStatus(ctx, req.AlertIds, req.Status)
	return
}
