// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnalysisTaskCameraAlgorithm is the golang structure of table sys_analysis_task_camera_algorithm for DAO operations like Where/Data.
type SysAnalysisTaskCameraAlgorithm struct {
	g.Meta      `orm:"table:sys_analysis_task_camera_algorithm, do:true"`
	Id          any         // 自增ID
	TaskId      any         // 任务ID
	CameraId    any         // 摄像头ID
	AlgorithmId any         // 算法ID
	Remark      any         // 备注
	CreateBy    any         // 创建者
	UpdateBy    any         // 更新者
	CreatedTime *gtime.Time // 任务创建时间
	UpdatedTime *gtime.Time // 任务更新时间
}
