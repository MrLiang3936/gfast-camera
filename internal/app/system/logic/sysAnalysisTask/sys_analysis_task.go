/*
* @desc:分析任务管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysAnalysisTask

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
	service.RegisterSysAnalysisTask(New())
}

func New() *sSysAnalysisTask {
	return &sSysAnalysisTask{}
}

type sSysAnalysisTask struct {
}

// List 分析任务列表
func (s *sSysAnalysisTask) List(ctx context.Context, req *system.AnalysisTaskSearchReq) (res *system.AnalysisTaskSearchRes, err error) {
	res = new(system.AnalysisTaskSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAnalysisTask.Ctx(ctx)
		// 修复点1：先初始化分页参数，避免req为nil时panic
		pageNum := 1
		pageSize := consts.PageSize
		if req != nil {
			// 条件过滤
			if req.Name != "" {
				m = m.Where("name like ?", "%"+req.Name+"%")
			}
			if req.State != "" {
				m = m.Where("state", req.State)
			}
			if req.Type != "" {
				m = m.Where("type", req.Type)
			}
			if req.WorkTimeType != "" {
				m = m.Where("work_time_type", req.WorkTimeType)
			}
			// 覆盖分页参数（仅当req非nil时）
			if req.PageNum > 0 {
				pageNum = req.PageNum
			}
			if req.PageSize > 0 {
				pageSize = req.PageSize
			}
		}
		// 获取总数
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取分析任务失败")
		// 赋值分页参数
		res.CurrentPage = pageNum
		// 分页查询
		err = m.Page(pageNum, pageSize).Order("created_time desc").Scan(&res.TaskList)
		liberr.ErrIsNil(ctx, err, "获取分析任务失败")
	})
	return
}

func (s *sSysAnalysisTask) Add(ctx context.Context, req *system.AnalysisTaskAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAnalysisTask.Ctx(ctx).Insert(do.SysAnalysisTask{
			Name:         req.Name,
			State:        req.State,
			WorkTimeType: req.WorkTimeType,
			WorkTime:     req.WorkTime,
			Type:         req.Type,
			Remark:       req.Remark,
			CreateBy:     service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加分析任务失败")
	})
	return
}

func (s *sSysAnalysisTask) Edit(ctx context.Context, req *system.AnalysisTaskEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAnalysisTask.Ctx(ctx).WherePri(req.Id).Update(do.SysAnalysisTask{
			Name:         req.Name,
			State:        req.State,
			WorkTimeType: req.WorkTimeType,
			WorkTime:     req.WorkTime,
			Type:         req.Type,
			Remark:       req.Remark,
			UpdateBy:     service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改分析任务失败")
	})
	return
}

// Delete 删除分析任务
func (s *sSysAnalysisTask) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 修复点2：使用WhereIn替代拼接字符串，更规范且安全
		_, err = dao.SysAnalysisTask.Ctx(ctx).WhereIn(dao.SysAnalysisTask.Columns().Id, ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// Get 根据ID获取分析任务
func (s *sSysAnalysisTask) Get(ctx context.Context, id int) (res *system.AnalysisTaskGetRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		res = new(system.AnalysisTaskGetRes)
		err = dao.SysAnalysisTask.Ctx(ctx).
			Where(dao.SysAnalysisTask.Columns().Id, id).
			Scan(res)
		liberr.ErrIsNil(ctx, err, "获取分析任务数据失败")
	})
	return
}
