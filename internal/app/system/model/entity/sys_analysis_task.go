// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnalysisTask is the golang structure for table sys_analysis_task.
type SysAnalysisTask struct {
	Id           uint        `json:"id"           orm:"id"             description:"任务自增ID"`
	Name         string      `json:"name"         orm:"name"           description:"任务名称（如：办公室行为观察）"`
	State        string      `json:"state"        orm:"state"          description:"任务运行状态（run/stop等）"`
	WorkTimeType string      `json:"workTimeType" orm:"work_time_type" description:"任务执行时间类型（everyday/week指定等）"`
	WorkTime     string      `json:"workTime"     orm:"work_time"      description:"任务执行时间配置，包含weekdays（周几）和periods（时间段）"`
	Type         string      `json:"type"         orm:"type"           description:"任务类型（video_analysis：视频分析）"`
	Remark       string      `json:"remark"       orm:"remark"         description:"备注"`
	CreateBy     uint        `json:"createBy"     orm:"create_by"      description:"创建者"`
	UpdateBy     uint        `json:"updateBy"     orm:"update_by"      description:"更新者"`
	CreatedTime  *gtime.Time `json:"createdTime"  orm:"created_time"   description:"任务创建时间"`
	UpdatedTime  *gtime.Time `json:"updatedTime"  orm:"updated_time"   description:"任务更新时间"`
	AlertFreq    uint        `json:"alertFreq"    orm:"alert_freq"     description:"报警频率 秒 (多少秒报警一次)"`
	AlertKeep    uint        `json:"alertKeep"    orm:"alert_keep"     description:"报警持续 秒 (出现多少秒就算)"`
}
