/*
* @desc:分析任务管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysAnalysisTask

import (
	"context"
	"encoding/json"
	"fmt"

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

func (s *sSysAnalysisTask) UpdateState(ctx context.Context, id int, state string) (res *system.AnalysisTaskUpdateStateRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		res = new(system.AnalysisTaskUpdateStateRes)

		// 验证状态值是否合法
		if state != "run" && state != "stop" {
			err = fmt.Errorf("invalid state value: %s, only 'run' or 'stop' allowed", state)
			return
		}

		// 更新数据库中的任务状态
		result, err := dao.SysAnalysisTask.Ctx(ctx).
			Where(dao.SysAnalysisTask.Columns().Id, id).
			Update(do.SysAnalysisTask{
				State:    state,
				UpdateBy: service.Context().GetUserId(ctx),
			})
		if err != nil {
			liberr.ErrIsNil(ctx, err, "更新分析任务状态失败")
			return
		}

		// 检查是否有记录被更新
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			err = fmt.Errorf("未找到指定的分析任务或无需更新")
			return
		}
	})
	return
}

func (s *sSysAnalysisTask) callRknnStart(ctx context.Context, modelFile string, installFlag bool) error {
	// 获取配置中的rknn安装地址
	var rknnStartUrl string
	if installFlag {
		rknnStartUrl = g.Cfg().MustGet(ctx, "rknn.start").String()
	} else {
		rknnStartUrl = g.Cfg().MustGet(ctx, "rknn.stop").String()
	}
	if rknnStartUrl == "" {
		return fmt.Errorf("rknn install url is not configured")
	}

	// 发送POST请求到rknn.install接口
	response, err := g.Client().Post(context.Background(), rknnStartUrl,
		g.Map{
			//"data": [],
		})
	if err != nil {
		g.Log().Errorf(context.Background(), "Failed to call rknn install API: %v", err)
		return err
	}
	defer response.Close()

	if response.StatusCode != 200 {
		errMsg := fmt.Sprintf("rknn install API returned status code: %d", response.StatusCode)
		g.Log().Errorf(context.Background(), errMsg)
		return fmt.Errorf(errMsg)
	}

	respStr := response.ReadAllString()
	if err != nil {
		errMsg := fmt.Sprintf("failed to read rknn install API response body: %v", err)
		g.Log().Errorf(ctx, errMsg)
		return fmt.Errorf(errMsg)
	}
	g.Log().Infof(ctx, "Rknn install response: %s", respStr)

	// 定义result变量，解析JSON响应体到map（核心完善点）
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(respStr), &result); err != nil {
		errMsg := fmt.Sprintf("failed to parse rknn install API response JSON: %v, raw response: %s", err, respStr)
		g.Log().Errorf(ctx, errMsg)
		return fmt.Errorf(errMsg)
	}

	// 检查返回的code是否为0（成功）或status是否为"repeat"（也视为成功）
	code, codeOk := result["code"].(float64)
	data, dataOk := result["data"].(map[string]interface{})

	// 判断是否成功（code为0）或重复安装（status为"repeat"）
	isSuccess := codeOk && code == 0
	isRepeat := dataOk && data["status"] == "repeat"

	if isSuccess || isRepeat {
		if isRepeat {
			g.Log().Infof(context.Background(), "Rknn install repeat for model: %s, message: %s", modelFile, data["message"])
		} else {
			g.Log().Infof(context.Background(), "Rknn install success for model: %s", modelFile)
		}
	} else {
		g.Log().Errorf(context.Background(), "Rknn install failed for model: %s, response: %+v", modelFile, result)
		return fmt.Errorf("rknn install failed for model: %s, response: %+v", modelFile, result)
	}

	return nil
}
