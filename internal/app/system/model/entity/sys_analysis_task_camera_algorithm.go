// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnalysisTaskCameraAlgorithm is the golang structure for table sys_analysis_task_camera_algorithm.
type SysAnalysisTaskCameraAlgorithm struct {
	Id          uint        `json:"id"          orm:"id"           description:"自增ID"`
	TaskId      uint        `json:"taskId"      orm:"task_id"      description:"任务ID"`
	CameraId    uint        `json:"cameraId"    orm:"camera_id"    description:"摄像头ID"`
	AlgorithmId uint        `json:"algorithmId" orm:"algorithm_id" description:"算法ID"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	CreateBy    uint        `json:"createBy"    orm:"create_by"    description:"创建者"`
	UpdateBy    uint        `json:"updateBy"    orm:"update_by"    description:"更新者"`
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" description:"任务创建时间"`
	UpdatedTime *gtime.Time `json:"updatedTime" orm:"updated_time" description:"任务更新时间"`
}
