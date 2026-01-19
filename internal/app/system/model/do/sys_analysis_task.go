// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnalysisTask is the golang structure of table sys_analysis_task for DAO operations like Where/Data.
type SysAnalysisTask struct {
	g.Meta       `orm:"table:sys_analysis_task, do:true"`
	Id           any         // 任务自增ID
	Name         any         // 任务名称（如：办公室行为观察）
	State        any         // 任务运行状态（run/stop等）
	WorkTimeType any         // 任务执行时间类型（everyday/week指定等）
	WorkTime     any         // 任务执行时间配置，包含weekdays（周几）和periods（时间段）
	Type         any         // 任务类型（video_analysis：视频分析）
	Remark       any         // 备注
	CreateBy     any         // 创建者
	UpdateBy     any         // 更新者
	CreatedTime  *gtime.Time // 任务创建时间
	UpdatedTime  *gtime.Time // 任务更新时间
}
