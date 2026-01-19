/*
* @desc:分析任务摄像头算法关联管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysAnalysisTaskCameraAlgorithm

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
	service.RegisterSysAnalysisTaskCameraAlgorithm(New())
}

func New() *sSysAnalysisTaskCameraAlgorithm {
	return &sSysAnalysisTaskCameraAlgorithm{}
}

type sSysAnalysisTaskCameraAlgorithm struct {
}

// List 分析任务摄像头算法关联列表
func (s *sSysAnalysisTaskCameraAlgorithm) List(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmSearchReq) (res *system.AnalysisTaskCameraAlgorithmSearchRes, err error) {
	res = new(system.AnalysisTaskCameraAlgorithmSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx) // 假设存在此表
		if req != nil {
			if req.CameraId > 0 {
				m = m.Where("camera_id", req.CameraId)
			}
			if req.AlgorithmId > 0 {
				m = m.Where("algorithm_id", req.AlgorithmId)
			}
		}
		total, err := m.Count()
		liberr.ErrIsNil(ctx, err, "获取分析任务摄像头算法关联失败")
		res.Total = int64(total)

		page := req.Page
		if page == 0 {
			page = 1
		}
		size := req.Size
		if size == 0 {
			size = consts.PageSize
		}
		res.Page = page
		res.Size = size

		err = m.Page(page, size).Order("id desc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取分析任务摄像头算法关联失败")
	})
	return
}

func (s *sSysAnalysisTaskCameraAlgorithm) Add(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).Insert(do.SysAnalysisTaskCameraAlgorithm{ // 假设存在此表
			CameraId:    req.CameraId,
			AlgorithmId: req.AlgorithmId,
			Remark:      req.Remark,
			CreateBy:    service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加分析任务摄像头算法关联失败")
	})
	return
}

func (s *sSysAnalysisTaskCameraAlgorithm) Edit(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).WherePri(req.Id).Update(do.SysAnalysisTaskCameraAlgorithm{ // 假设存在此表
			CameraId:    req.CameraId,
			AlgorithmId: req.AlgorithmId,
			Remark:      req.Remark,
			UpdateBy:    service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改分析任务摄像头算法关联失败")
	})
	return
}

func (s *sSysAnalysisTaskCameraAlgorithm) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).Where(dao.SysAnalysisTaskCameraAlgorithm.Columns().Id+" in(?)", ids).Delete() // 假设存在此表
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
