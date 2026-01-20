/*
* @desc:分析任务摄像头算法关联管理
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

var AnalysisTaskCameraAlgorithm = analysisTaskCameraAlgorithmController{}

type analysisTaskCameraAlgorithmController struct {
	BaseController
}

// List 分析任务摄像头算法关联列表
func (c *analysisTaskCameraAlgorithmController) List(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmSearchReq) (res *system.AnalysisTaskCameraAlgorithmSearchRes, err error) {
	res, err = service.SysAnalysisTaskCameraAlgorithm().List(ctx, req)
	return
}

// Add 添加分析任务摄像头算法关联
func (c *analysisTaskCameraAlgorithmController) Add(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmAddReq) (res *system.AnalysisTaskCameraAlgorithmAddRes, err error) {
	err = service.SysAnalysisTaskCameraAlgorithm().Add(ctx, req)
	return
}

func (c *analysisTaskCameraAlgorithmController) AddBatch(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmAddBatchReq) (res *system.AnalysisTaskCameraAlgorithmAddBatchRes, err error) {
	err = service.SysAnalysisTaskCameraAlgorithm().AddBatch(ctx, req)
	return
}

// Edit 修改分析任务摄像头算法关联
func (c *analysisTaskCameraAlgorithmController) Edit(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmEditReq) (res *system.AnalysisTaskCameraAlgorithmEditRes, err error) {
	err = service.SysAnalysisTaskCameraAlgorithm().Edit(ctx, req)
	return
}

// Delete 删除分析任务摄像头算法关联
func (c *analysisTaskCameraAlgorithmController) Delete(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmDeleteReq) (res *system.AnalysisTaskCameraAlgorithmDeleteRes, err error) {
	err = service.SysAnalysisTaskCameraAlgorithm().Delete(ctx, req.Ids)
	return
}

// Delete 批量删除分析任务摄像头算法关联
func (c *analysisTaskCameraAlgorithmController) DeleteBatch(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmDeleteBatchReq) (res *system.AnalysisTaskCameraAlgorithmDeleteBatchRes, err error) {
	err = service.SysAnalysisTaskCameraAlgorithm().DeleteBatch(ctx, req)
	return
}
