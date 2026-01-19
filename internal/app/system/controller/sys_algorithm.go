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
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Algorithm = algorithmController{}

type algorithmController struct {
	BaseController
}

// List 列表
func (c *algorithmController) List(ctx context.Context, req *system.AlgorithmSearchReq) (res *system.AlgorithmSearchRes, err error) {
	res, err = service.SysAlgorithm().List(ctx, req)
	return
}

// Add 添加
func (c *algorithmController) Add(ctx context.Context, req *system.AlgorithmAddReq) (res *system.AlgorithmAddRes, err error) {
	err = service.SysAlgorithm().Add(ctx, req)
	return
}

// Edit 修改
func (c *algorithmController) Edit(ctx context.Context, req *system.AlgorithmEditReq) (res *system.CameraEditRes, err error) {
	err = service.SysAlgorithm().Edit(ctx, req)
	return
}

// Delete 删除
func (c *algorithmController) Delete(ctx context.Context, req *system.AlgorithmDeleteReq) (res *system.CameraDeleteRes, err error) {
	err = service.SysAlgorithm().Delete(ctx, req.Ids)
	return
}

// Get 查询详情
func (c *algorithmController) Get(ctx context.Context, req *system.AlgorithmGetReq) (res *system.AlgorithmGetRes, err error) {
	res, err = service.SysAlgorithm().Get(ctx, req.Id)
	return
}
