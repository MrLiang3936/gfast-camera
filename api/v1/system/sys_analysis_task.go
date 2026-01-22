/*
* @desc:分析任务相关参数
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/7 23:09
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type AnalysisTaskSearchReq struct {
	g.Meta       `path:"/analysisTask/list" tags:"分析任务管理" method:"get" summary:"分析任务列表"`
	Name         string `p:"name"`         // 任务名称
	State        string `p:"state"`        // 任务运行状态
	Type         string `p:"type"`         // 任务类型
	WorkTimeType string `p:"workTimeType"` // 任务执行时间类型
	commonApi.PageReq
}

type AnalysisTaskSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	TaskList []*entity.SysAnalysisTask `json:"taskList"`
}

type AnalysisTaskAddReq struct {
	g.Meta       `path:"/analysisTask/add" tags:"分析任务管理" method:"post" summary:"添加分析任务"`
	Name         string `p:"name" v:"required#任务名称不能为空"`
	State        string `p:"state" v:"required#任务运行状态不能为空"`
	WorkTimeType string `p:"workTimeType" `
	WorkTime     string `p:"workTime" default:"{}"` // 任务执行时间配置
	Type         string `p:"type" v:"required#任务类型不能为空"`
	Remark       string `p:"remark"`
}

type AnalysisTaskAddRes struct {
}

type AnalysisTaskEditReq struct {
	g.Meta       `path:"/analysisTask/edit" tags:"分析任务管理" method:"put" summary:"修改分析任务"`
	Id           uint   `p:"id" v:"required|min:1#id必须"`
	Name         string `p:"name" v:"required#任务名称不能为空"`
	State        string `p:"state" v:"required#任务运行状态不能为空"`
	WorkTimeType string `p:"workTimeType" v:"required#任务执行时间类型不能为空"`
	WorkTime     string `p:"workTime"` // 任务执行时间配置
	Type         string `p:"type" v:"required#任务类型不能为空"`
	Remark       string `p:"remark"`
}

type AnalysisTaskEditRes struct {
}

type AnalysisTaskDeleteReq struct {
	g.Meta `path:"/analysisTask/delete" tags:"分析任务管理" method:"delete" summary:"删除分析任务"`
	Ids    []int `p:"ids"`
}

type AnalysisTaskDeleteRes struct {
}

type AnalysisTaskGetReq struct {
	g.Meta `path:"/analysisTask/get" tags:"分析任务管理" method:"get" summary:"获取分析任务详情"`
	Id     int `p:"id" v:"required|min:1#id必须"`
}

type AnalysisTaskGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.SysAnalysisTask `json:"data"`
}

type AnalysisTaskUpdateStateReq struct {
	g.Meta `path:"/analysisTask/enable" tags:"分析任务管理" method:"get" summary:"启用分析任务"`
	Id     int    `p:"id" v:"required|min:1#id必须"`
	State  string `json:"state" v:"required|in:run,stop#state必须且只能是run或stop" dc:"任务运行状态(run/stop)等"`
}

type AnalysisTaskUpdateStateRes struct {
}
