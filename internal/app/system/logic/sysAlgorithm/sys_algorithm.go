/*
* @desc:摄像头管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysCamera

import (
	"context"
	"encoding/json"
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
		//count, err := dao.SysAlgorithm.Ctx(ctx).Where(do.SysAlgorithm{AlgorithmId: req.AlgorithmId}).Count()
		//if err != nil {
		//	liberr.ErrIsNil(ctx, err, "检查AlgorithmId是否重复时出错")
		//	return
		//}
		//if count > 0 {
		//	liberr.ErrIsNil(ctx, fmt.Errorf("AlgorithmId已存在"), "AlgorithmId已存在")
		//	return
		//}

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
		_, err = dao.SysAlgorithm.Ctx(ctx).WherePri(req.Id).Update(do.SysAlgorithm{
			//Intro:            req.Intro,
			//AlgorithmId:      req.AlgorithmId,
			//AlgorithmTaskId:  req.AlgorithmTaskId,
			//AlgorithmVersion: req.AlgorithmVersion,
			//CnName:           req.CnName,
			//EnName:           req.EnName,
			//CoverImageUrl:    req.CoverImageUrl,
			State: req.State,
			//Remark:           req.Remark,
			//ModelFile:        req.ModelFile,
			//Secret:           req.Secret,
			UpdateBy: service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改算法失败")
	})
	return
}

func (s *sSysAlgorithm) EditBatch(ctx context.Context, req *system.AlgorithmEditPatchReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 先查询需要更新的算法记录，以便获取模型文件信息
		var algorithms []*entity.SysAlgorithm
		err = dao.SysAlgorithm.Ctx(ctx).Where(dao.SysAlgorithm.Columns().Id+" in(?)", req.Ids).Scan(&algorithms)
		liberr.ErrIsNil(ctx, err, "查询算法数据失败")

		_, err = dao.SysAlgorithm.Ctx(ctx).Where(dao.SysAlgorithm.Columns().Id+" in(?)", req.Ids).Update(do.SysAlgorithm{
			State:    req.State,
			UpdateBy: service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "批量修改算法失败")

		// 如果状态变为启用且有模型文件，则调用rknn.install接口
		if req.State == "Installed" {
			for _, algorithm := range algorithms {
				if algorithm.ModelFile != "" {
					// 调用rknn.install接口
					err = s.callRknnInstall(ctx, algorithm.ModelFile, true)
					if err != nil {
						g.Log().Errorf(ctx, "Failed to install %s: %v", algorithm.ModelFile, err)
					}
				}
			}
		} else {
			for _, algorithm := range algorithms {
				if algorithm.ModelFile != "" {
					// 调用rknn.install接口
					err = s.callRknnInstall(ctx, algorithm.ModelFile, false)
					if err != nil {
						g.Log().Errorf(ctx, "Failed to uninstall %s: %v", algorithm.ModelFile, err)
					}
				}
			}
		}
	})
	return
}

// callRknnInstall 调用rknn安装接口
func (s *sSysAlgorithm) callRknnInstall(ctx context.Context, modelFile string, installFlag bool) error {
	// 获取配置中的rknn安装地址
	var rknnInstallUrl string
	if installFlag {
		rknnInstallUrl = g.Cfg().MustGet(ctx, "rknn.install").String()
	} else {
		rknnInstallUrl = g.Cfg().MustGet(ctx, "rknn.uninstall").String()
	}
	if rknnInstallUrl == "" {
		return fmt.Errorf("rknn install url is not configured")
	}

	// 发送POST请求到rknn.install接口
	response, err := g.Client().Post(context.Background(), rknnInstallUrl,
		g.Map{
			"modelFile": modelFile,
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
