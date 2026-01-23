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
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysAlertConfig(New())
}

func New() *sSysAlertConfig {
	return &sSysAlertConfig{}
}

type sSysAlertConfig struct {
}

func (s sSysAlertConfig) List(ctx context.Context, req *system.AlertConfigSearchReq) (res *system.AlertConfigSearchRes, err error) {
	res = new(system.AlertConfigSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAlertConfig.Ctx(ctx)
		if req != nil {
			if req.VoiceName != "" {
				m = m.Where("voice_name like ?", "%"+req.VoiceName+"%")
			}
			if req.VoiceUrl != "" {
				m = m.Where("voice_url like ?", "%"+req.VoiceUrl+"%")
			}
			if req.IsDefault != nil {
				m = m.Where("is_default", *req.IsDefault)
			}
			if req.Enable != nil {
				m = m.Where("enable", *req.Enable)
			}
			if req.Degree != "" {
				m = m.Where("degree like ?", "%"+req.Degree+"%")
			}
			if req.Pop != nil {
				m = m.Where("pop", *req.Pop)
			}
			if req.PopType != "" {
				m = m.Where("pop_type like ?", "%"+req.PopType+"%")
			}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取告警配置失败")

		pageNum := req.PageNum
		if pageNum == 0 {
			pageNum = 1
		}
		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = consts.PageSize
		}
		res.CurrentPage = pageNum

		err = m.Page(pageNum, pageSize).Order("id asc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取告警配置失败")
	})
	return
}

func (s sSysAlertConfig) Add(ctx context.Context, req *system.AlertConfigAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlertConfig.Ctx(ctx).Insert(do.SysAlertConfig{
			VoiceName: req.VoiceName,
			VoiceUrl:  req.VoiceUrl,
			IsDefault: req.IsDefault,
			Enable:    req.Enable,
			Degree:    req.Degree,
			Pop:       req.Pop,
			PopType:   req.PopType,
			CreatedBy: service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加告警配置失败")
	})
	return
}

func (s sSysAlertConfig) Edit(ctx context.Context, req *system.AlertConfigEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlertConfig.Ctx(ctx).WherePri(req.Id).Update(do.SysAlertConfig{
			VoiceName: req.VoiceName,
			VoiceUrl:  req.VoiceUrl,
			IsDefault: req.IsDefault,
			Enable:    req.Enable,
			Degree:    req.Degree,
			Pop:       req.Pop,
			PopType:   req.PopType,
			UpdatedBy: service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改告警配置失败")
	})
	return
}

func (s sSysAlertConfig) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAlertConfig.Ctx(ctx).Where(dao.SysAlertConfig.Columns().Id+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除告警配置失败")
	})
	return
}

func (s sSysAlertConfig) Get(ctx context.Context, id int) (res *entity.SysAlertConfig, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysAlertConfig.Ctx(ctx).WherePri(id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取告警配置失败")
	})
	return
}
