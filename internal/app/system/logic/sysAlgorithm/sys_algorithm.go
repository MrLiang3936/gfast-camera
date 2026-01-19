/*
* @desc:摄像头管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysCamera

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
	service.RegisterSysAlgorithm(New())
}

func New() *sSysAlgorithm {
	return &sSysAlgorithm{}
}

type sSysAlgorithm struct {
}

// List 算法列表
func (s *sSysAlgorithm) List(ctx context.Context, req *system.AlgorithmSearchReq) (res *system.AlgorithmSearchRes, err error) {
	res = new(system.AlgorithmSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAlgorithm.Ctx(ctx)
		if req != nil {
			if req.CnName != "" {
				m = m.Where("cn_name like ?", "%"+req.CnName+"%")
			}
			if req.EnName != "" {
				m = m.Where("en_name like ?", "%"+req.EnName+"%")
			}
			if req.AlgorithmId != "" {
				m = m.Where("algorithm_id like ?", "%"+req.AlgorithmId+"%")
			}
			if req.AlgorithmTaskId != "" {
				m = m.Where("algorithm_task_id like ?", "%"+req.AlgorithmTaskId+"%")
			}
			if req.State != "" {
				m = m.Where("state", req.State)
			}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取算法失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		res.CurrentPage = req.PageNum
		err = m.Page(req.PageNum, req.PageSize).Order("created_at desc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取算法失败")
	})
	return
}

func (s *sSysAlgorithm) Add(ctx context.Context, req *system.AlgorithmAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查 AlgorithmId 是否已经存在
		count, err := dao.SysAlgorithm.Ctx(ctx).Where(do.SysAlgorithm{AlgorithmId: req.AlgorithmId}).Count()
		if err != nil {
			liberr.ErrIsNil(ctx, err, "检查AlgorithmId是否重复时出错")
			return
		}
		if count > 0 {
			liberr.ErrIsNil(ctx, fmt.Errorf("AlgorithmId已存在"), "AlgorithmId已存在")
			return
		}

		_, err = dao.SysAlgorithm.Ctx(ctx).Insert(do.SysAlgorithm{
			Intro:            req.Intro,
			AlgorithmId:      req.AlgorithmId,
			AlgorithmTaskId:  req.AlgorithmTaskId,
			AlgorithmVersion: req.AlgorithmVersion,
			CnName:           req.CnName,
			EnName:           req.EnName,
			CoverImageUrl:    req.CoverImageUrl,
			State:            req.State,
			Remark:           req.Remark,
			ModelFile:        req.ModelFile,
			Secret:           req.Secret,
			CreateBy:         service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加算法失败")
	})
	return
}

func (s *sSysAlgorithm) Edit(ctx context.Context, req *system.AlgorithmEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlgorithm.Ctx(ctx).WherePri(req.AlgorithmId).Update(do.SysAlgorithm{
			Intro:            req.Intro,
			AlgorithmId:      req.AlgorithmId,
			AlgorithmTaskId:  req.AlgorithmTaskId,
			AlgorithmVersion: req.AlgorithmVersion,
			CnName:           req.CnName,
			EnName:           req.EnName,
			CoverImageUrl:    req.CoverImageUrl,
			State:            req.State,
			Remark:           req.Remark,
			ModelFile:        req.ModelFile,
			Secret:           req.Secret,
			UpdateBy:         service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改算法失败")
	})
	return
}

func (s *sSysAlgorithm) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlgorithm.Ctx(ctx).Where(dao.SysAlgorithm.Columns().Id+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 获取算法详情方法
func (s *sSysAlgorithm) Get(ctx context.Context, id int) (res *system.AlgorithmGetRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var algorithmInfo *entity.SysAlgorithm
		err = dao.SysAlgorithm.Ctx(ctx).Where(dao.SysAlgorithm.Columns().Id, id).Scan(&algorithmInfo)
		liberr.ErrIsNil(ctx, err, "获取算法数据失败")

		res = &system.AlgorithmGetRes{
			Data: algorithmInfo,
		}
	})
	return
}
