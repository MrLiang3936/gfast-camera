package sysAnalysisTaskCameraAlgorithm

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
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
		if len(req.Data) == 0 {
			return // 没有数据则直接返回
		}

		// 获取任务ID
		taskId := req.Data[0].TaskId

		// 先删除该任务下的所有关联关系
		_, err = dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).
			Where(dao.SysAnalysisTaskCameraAlgorithm.Columns().TaskId, taskId).
			Delete()
		liberr.ErrIsNil(ctx, err, "删除任务关联关系失败")

		// 批量插入新的数据
		var records []do.SysAnalysisTaskCameraAlgorithm
		userId := service.Context().GetUserId(ctx)
		for _, item := range req.Data {
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
		// 使用 JOIN 查询关联数据
		m := dao.SysAnalysisTaskCameraAlgorithm.Ctx(ctx).
			LeftJoin(dao.SysCamera.Table(),
				fmt.Sprintf("%s.camera_id = %s.camera_id",
					dao.SysAnalysisTaskCameraAlgorithm.Table(),
					dao.SysCamera.Table())).
			LeftJoin(dao.SysAlgorithm.Table(),
				fmt.Sprintf("%s.algorithm_id = %s.id",
					dao.SysAnalysisTaskCameraAlgorithm.Table(),
					dao.SysAlgorithm.Table()))

		if req != nil {
			if req.TaskId > 0 {
				m = m.Where(fmt.Sprintf("%s.task_id", dao.SysAnalysisTaskCameraAlgorithm.Table()), req.TaskId)
			}
			if req.CameraId > 0 {
				m = m.Where(fmt.Sprintf("%s.camera_id", dao.SysAnalysisTaskCameraAlgorithm.Table()), req.CameraId)
			}
			if req.AlgorithmId > 0 {
				m = m.Where(fmt.Sprintf("%s.algorithm_id", dao.SysAnalysisTaskCameraAlgorithm.Table()), req.AlgorithmId)
			}
		}

		// 查询总数（如果需要）
		total, err := m.Count()
		liberr.ErrIsNil(ctx, err, "获取分析任务摄像头算法关联失败")
		res.Total = int64(total)

		// 取消分页，查询所有数据
		var list []*struct {
			*entity.SysAnalysisTaskCameraAlgorithm
			CameraName     *string `json:"camera_name"`
			GroupId        *uint   `json:"group_id"`
			DeviceType     *int    `json:"device_type"`
			StreamUrl      *string `json:"stream_url"`
			PreviewImg     *string `json:"preview_img"`
			CameraRemark   *string `json:"camera_remark"`
			CnName         *string `json:"cn_name"`
			EnName         *string `json:"en_name"`
			Intro          *string `json:"intro"`
			State          *string `json:"state"`
			AlgorithmState *string `json:"algorithm_state"`
		}

		err = m.Fields(
			fmt.Sprintf("%s.*, %s.camera_name, %s.group_id, %s.device_type, %s.stream_url, %s.preview_img, %s.remark as camera_remark, %s.cn_name, %s.en_name, %s.intro, %s.state as algorithm_state",
				dao.SysAnalysisTaskCameraAlgorithm.Table(),
				dao.SysCamera.Table(),
				dao.SysCamera.Table(),
				dao.SysCamera.Table(),
				dao.SysCamera.Table(),
				dao.SysCamera.Table(),
				dao.SysAlgorithm.Table(),
				dao.SysAlgorithm.Table(),
				dao.SysAlgorithm.Table(),
				dao.SysAlgorithm.Table(),
				dao.SysAlgorithm.Table())).
			Order(fmt.Sprintf("%s.id desc", dao.SysAnalysisTaskCameraAlgorithm.Table())).
			Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取分析任务摄像头算法关联失败")

		// 将结果转换为返回格式
		for _, item := range list {
			var cameraInfo *entity.SysCamera
			if item.CameraName != nil {
				cameraInfo = &entity.SysCamera{
					CameraId:   item.SysAnalysisTaskCameraAlgorithm.CameraId,
					CameraName: *item.CameraName,
					GroupId:    *item.GroupId,
					DeviceType: *item.DeviceType,
					StreamUrl:  *item.StreamUrl,
					PreviewImg: *item.PreviewImg,
				}
			}

			// 创建算法信息对象
			var algorithmInfo *entity.SysAlgorithm
			if item.CnName != nil {
				algorithmInfo = &entity.SysAlgorithm{
					Id:     item.SysAnalysisTaskCameraAlgorithm.AlgorithmId,
					CnName: *item.CnName,
					EnName: *item.EnName,
					Intro:  *item.Intro,
					State:  *item.AlgorithmState,
				}
			}

			res.List = append(res.List, &system.AnalysisTaskCameraAlgorithmInfo{
				Id:          item.SysAnalysisTaskCameraAlgorithm.Id,
				TaskId:      item.SysAnalysisTaskCameraAlgorithm.TaskId,
				CameraId:    item.SysAnalysisTaskCameraAlgorithm.CameraId,
				AlgorithmId: item.SysAnalysisTaskCameraAlgorithm.AlgorithmId,
				Remark:      item.SysAnalysisTaskCameraAlgorithm.Remark,
				Camera:      cameraInfo,
				Algorithm:   algorithmInfo,
			})
		}
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
