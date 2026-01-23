/*
* @desc:岗位管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/7 23:12
 */

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var AlertConfig = alertConfigController{}

type alertConfigController struct {
	BaseController
}

// List 报警配置列表
func (c *alertConfigController) List(ctx context.Context, req *system.AlertConfigSearchReq) (res *system.AlertConfigSearchRes, err error) {
	res, err = service.SysAlertConfig().List(ctx, req)
	return
}

// Add 添加报警配置
func (c *alertConfigController) Add(ctx context.Context, req *system.AlertConfigAddReq) (res *system.AlertConfigAddRes, err error) {
	err = service.SysAlertConfig().Add(ctx, req)
	return
}

// Edit 修改报警配置
func (c *alertConfigController) Edit(ctx context.Context, req *system.AlertConfigEditReq) (res *system.AlertConfigEditRes, err error) {
	err = service.SysAlertConfig().Edit(ctx, req)
	return
}

// Get 查询报警配置
func (c *alertConfigController) Get(ctx context.Context, id int) (res *entity.SysAlertConfig, err error) {
	res, err = service.SysAlertConfig().Get(ctx, id)
	return
}
