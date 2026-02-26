/*
* @desc:岗位管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/7 23:12
 */

package controller

import (
	"context"
	"strings"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Camera = cameraController{}

type cameraController struct {
	BaseController
}

// List 摄像头列表
func (c *cameraController) List(ctx context.Context, req *system.CameraSearchReq) (res *system.CameraSearchRes, err error) {
	res, err = service.SysCamera().List(ctx, req)
	return
}

// Add 添加摄像头
func (c *cameraController) Add(ctx context.Context, req *system.CameraAddReq) (res *system.CameraAddRes, err error) {
	req.PreviewImg = "/camera/" + req.PreviewImg
	err = service.SysCamera().Add(ctx, req)
	return
}

// Edit 修改摄像头
func (c *cameraController) Edit(ctx context.Context, req *system.CameraEditReq) (res *system.CameraEditRes, err error) {
	// 判断PreviewImg是否以"/camera/"开头，不是就加上
	if req.PreviewImg != "" && !strings.HasPrefix(req.PreviewImg, "/camera/") {
		req.PreviewImg = "/camera/" + req.PreviewImg
	}
	err = service.SysCamera().Edit(ctx, req)
	return
}

// Delete 删除摄像头
func (c *cameraController) Delete(ctx context.Context, req *system.CameraDeleteReq) (res *system.CameraDeleteRes, err error) {
	err = service.SysCamera().Delete(ctx, req.Ids)
	return
}

// Get 查询摄像头详情
func (c *cameraController) Get(ctx context.Context, req *system.CameraGetReq) (res *model.CameraGetMediaRes, err error) {
	res, err = service.SysCamera().Get(ctx, req.Id)
	return
}
