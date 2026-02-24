// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAlertLog is the golang structure of table sys_alert_log for DAO operations like Where/Data.
type SysAlertLog struct {
	g.Meta      `orm:"table:sys_alert_log, do:true"`
	AlertId     any         // 主键ID
	TaskId      any         // 任务ID
	CameraId    any         // 摄像头ID
	AlgorithmId any         // 算法ID
	EnName      any         // 报警类别
	Degree      any         // 报警级别
	ImageUrl    any         // 图片路径
	ImageName   any         // 图片名称
	Status      any         // 归档状态：0-暂不归档，1-正确处理，2-错误处理
	Remark      any         // 备注
	CreatedBy   any         // 创建人
	UpdatedBy   any         // 修改人
	CreateAt    *gtime.Time // 创建时间
	UpdateAt    *gtime.Time // 更新时间
}
