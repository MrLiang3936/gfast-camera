// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAlertLog is the golang structure for table sys_alert_log.
type SysAlertLog struct {
	AlertId     uint        `json:"alertId"     orm:"alert_id"     description:"主键ID"`
	TaskId      uint        `json:"taskId"      orm:"task_id"      description:"任务ID"`
	CameraId    uint        `json:"cameraId"    orm:"camera_id"    description:"摄像头ID"`
	AlgorithmId uint        `json:"algorithmId" orm:"algorithm_id" description:"算法ID"`
	EnName      string      `json:"enName"      orm:"en_name"      description:"报警类别"`
	Degree      int         `json:"degree"      orm:"degree"       description:"报警级别"`
	ImageUrl    string      `json:"imageUrl"    orm:"image_url"    description:"图片路径"`
	ImageName   string      `json:"imageName"   orm:"image_name"   description:"图片名称"`
	Status      int         `json:"status"      orm:"status"       description:"归档状态：0-暂不归档，1-正确处理，2-错误处理"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	CreatedBy   uint64      `json:"createdBy"   orm:"created_by"   description:"创建人"`
	UpdatedBy   uint64      `json:"updatedBy"   orm:"updated_by"   description:"修改人"`
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"    description:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"    description:"更新时间"`
}
