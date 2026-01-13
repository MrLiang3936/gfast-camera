/*
* @desc:摄像头管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysCamera

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysCamera(New())
}

func New() *sSysCamera {
	return &sSysCamera{}
}

type sSysCamera struct {
}

// List 摄像头列表
func (s *sSysCamera) List(ctx context.Context, req *system.CameraSearchReq) (res *system.CameraSearchRes, err error) {
	res = new(system.CameraSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysCamera.Ctx(ctx)
		if req != nil {
			if req.CameraName != "" {
				m = m.Where("camera_name like ?", "%"+req.CameraName+"%")
			}
			//if req.Status != "" {
			//	m = m.Where("status", gconv.Uint(req.Status))
			//}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取摄像头失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		res.CurrentPage = req.PageNum
		err = m.Page(req.PageNum, req.PageSize).Order("created_at desc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取摄像头失败")
	})
	return
}

func (s *sSysCamera) Add(ctx context.Context, req *system.CameraAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysCamera.Ctx(ctx).Insert(do.SysCamera{
			GroupId:    req.GroupId,
			CameraName: req.CameraName,
			GpsInfo:    req.GpsInfo,
			DeviceType: req.DeviceType,
			StreamUrl:  req.StreamUrl,
			PreviewImg: req.PreviewImg,
			Remark:     req.Remark,
			CreateBy:   service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加摄像头失败")
	})
	return
}

func (s *sSysCamera) Edit(ctx context.Context, req *system.CameraEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysCamera.Ctx(ctx).WherePri(req.CameraId).Update(do.SysCamera{
			CameraId:   req.CameraId,
			GroupId:    req.GroupId,
			CameraName: req.CameraName,
			GpsInfo:    req.GpsInfo,
			DeviceType: req.DeviceType,
			StreamUrl:  req.StreamUrl,
			PreviewImg: req.PreviewImg,
			Remark:     req.Remark,
			UpdateBy:   service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改摄像头失败")
	})
	return
}

func (s *sSysCamera) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysCamera.Ctx(ctx).Where(dao.SysCamera.Columns().CameraId+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sSysCamera) Get(ctx context.Context, id int) (res *system.CameraGetRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysCamera.Ctx(ctx).Where(dao.SysCamera.Columns().CameraId, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取摄像头数据失败")
	})
	return
}
