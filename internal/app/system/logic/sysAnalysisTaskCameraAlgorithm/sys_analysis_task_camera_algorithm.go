package sysAnalysisTaskCameraAlgorithm

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
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
	service.RegisterSysAnalysisTaskCameraAlgorithm(New())
}

func New() *sSysAnalysisTaskCameraAlgorithm {
	return &sSysAnalysisTaskCameraAlgorithm{}
}

type sSysAnalysisTaskCameraAlgorithm struct {
}

func (s *sSysAnalysisTaskCameraAlgorithm) AddBatch(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmAddBatchReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 根据req.TaskId和Data里面对象的CameraId和AlgorithmId查询有没有重复的,有则去掉
		existingRecords, err := s.getExistingRecords(ctx, req.Data)
		liberr.ErrIsNil(ctx, err, "查询现有记录失败")
		g.Log().Warningf(ctx, "【查询existingRecords现有记录】 %s", gjson.MustEncodeString(existingRecords))
		// 过滤掉已存在的记录
		filteredData := s.filterDuplicateRecords(req.Data, existingRecords)
		g.Log().Warningf(ctx, "【查询filteredData现有记录】 %s", gjson.MustEncodeString(filteredData))

		if len(filteredData) == 0 {
			return // 没有新记录需要插入
		}

		// 批量插入数据
		var records []do.SysAnalysisTaskCameraAlgorithm
		userId := service.Context().GetUserId(ctx)
		for _, item := range filteredData {
			record := do.SysAnalysisTaskCameraAlgorithm{
				TaskId:      item.TaskId,
				CameraId:    item.CameraId,
				AlgorithmId: item.AlgorithmId,
				Remark:      item.Remark,
				CreateBy:    userId,
			}
			records = append(records, record)
		}

		if len(records) > 0 {
			_, err = dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).Insert(&records)
			liberr.ErrIsNil(ctx, err, "批量添加分析任务摄像头算法关联失败")
		}
	})
	return
}

// getExistingRecords 查询已存在的记录
func (s *sSysAnalysisTaskCameraAlgorithm) getExistingRecords(
	ctx context.Context, data []*entity.SysAnalysisTaskCameraAlgorithm) ([]*entity.SysAnalysisTaskCameraAlgorithm, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var taskId uint
	var cameraIds, algorithmIds []uint
	for _, item := range data {
		cameraIds = append(cameraIds, item.CameraId)
		algorithmIds = append(algorithmIds, item.AlgorithmId)
		taskId = item.TaskId
	}

	var existingRecords []*entity.SysAnalysisTaskCameraAlgorithm
	err := dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).
		Where(dao.SysAnalysisTaskCameraAlgorithm.Columns().TaskId, taskId).
		Where(dao.SysAnalysisTaskCameraAlgorithm.Columns().CameraId+" IN (?)", cameraIds).
		Where(dao.SysAnalysisTaskCameraAlgorithm.Columns().AlgorithmId+" IN (?)", algorithmIds).
		Scan(&existingRecords)

	return existingRecords, err
}

// filterDuplicateRecords 过滤重复记录
func (s *sSysAnalysisTaskCameraAlgorithm) filterDuplicateRecords(
	newData []*entity.SysAnalysisTaskCameraAlgorithm, existingRecords []*entity.SysAnalysisTaskCameraAlgorithm) []*entity.SysAnalysisTaskCameraAlgorithm {
	// 构建已存在记录的映射，用于快速查找
	existingMap := make(map[string]bool)
	for _, record := range existingRecords {
		key := fmt.Sprintf("%d_%d_%d", record.TaskId, record.CameraId, record.AlgorithmId)
		existingMap[key] = true
	}

	// 过滤新数据
	var filteredData []*entity.SysAnalysisTaskCameraAlgorithm
	for _, item := range newData {
		key := fmt.Sprintf("%d_%d_%d", item.TaskId, item.CameraId, item.AlgorithmId)
		if !existingMap[key] {
			filteredData = append(filteredData, item)
		}
	}

	return filteredData
}

// List 分析任务摄像头算法关联列表
func (s *sSysAnalysisTaskCameraAlgorithm) List(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmSearchReq) (res *system.AnalysisTaskCameraAlgorithmSearchRes, err error) {
	res = new(system.AnalysisTaskCameraAlgorithmSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx) //
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
		_, err = dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).Insert(do.SysAnalysisTaskCameraAlgorithm{
			TaskId:      req.TaskId,
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
		_, err = dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).WherePri(req.Id).Update(do.SysAnalysisTaskCameraAlgorithm{
			TaskId:      req.TaskId,
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

func (s *sSysAnalysisTaskCameraAlgorithm) DeleteBatch(ctx context.Context, req *system.AnalysisTaskCameraAlgorithmDeleteBatchReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 根据Data数组里面对象的CameraId，AlgorithmId，TaskId删除
		if len(req.Data) == 0 {
			return // 如果没有数据要删除，则直接返回
		}

		// 构建 OR 条件进行批量删除
		queryBuilder := dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx)

		for i, item := range req.Data {
			if i == 0 {
				queryBuilder = queryBuilder.Where(
					"(task_id = ? AND camera_id = ? AND algorithm_id = ?)",
					item.TaskId, item.CameraId, item.AlgorithmId,
				)
			} else {
				queryBuilder = queryBuilder.WhereOr(
					"(task_id = ? AND camera_id = ? AND algorithm_id = ?)",
					item.TaskId, item.CameraId, item.AlgorithmId,
				)
			}
		}

		_, err = queryBuilder.Delete()
		liberr.ErrIsNil(ctx, err, "批量删除分析任务摄像头算法关联失败")
	})
	return
}
