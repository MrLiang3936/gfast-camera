/*
* @desc:分析任务管理
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

var AnalysisTask = analysisTaskController{}

type analysisTaskController struct {
	BaseController
}

// List 分析任务列表
func (c *analysisTaskController) List(ctx context.Context, req *system.AnalysisTaskSearchReq) (res *system.AnalysisTaskSearchRes, err error) {
	res, err = service.SysAnalysisTask().List(ctx, req)
	return
}

// Add 添加分析任务
func (c *analysisTaskController) Add(ctx context.Context, req *system.AnalysisTaskAddReq) (res *system.AnalysisTaskAddRes, err error) {
	err = service.SysAnalysisTask().Add(ctx, req)
	return
}

// Edit 修改分析任务
func (c *analysisTaskController) Edit(ctx context.Context, req *system.AnalysisTaskEditReq) (res *system.AnalysisTaskEditRes, err error) {
	err = service.SysAnalysisTask().Edit(ctx, req)
	return
}

// Delete 删除分析任务
func (c *analysisTaskController) Delete(ctx context.Context, req *system.AnalysisTaskDeleteReq) (res *system.AnalysisTaskDeleteRes, err error) {
	err = service.SysAnalysisTask().Delete(ctx, req.Ids)
	return
}

func (c *analysisTaskController) Get(ctx context.Context, req *system.AnalysisTaskGetReq) (res *system.AnalysisTaskGetRes, err error) {
	res, err = service.SysAnalysisTask().Get(ctx, req.Id)
	return
}
