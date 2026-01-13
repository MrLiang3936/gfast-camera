// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCamera is the golang structure of table sys_camera for DAO operations like Where/Data.
type SysCamera struct {
	g.Meta     `orm:"table:sys_camera, do:true"`
	CameraId   any         // 摄像头主键
	CameraName any         // 摄像头名称（必填）
	GroupId    any         // 所属分组主键（关联sys_camera_group表）
	GpsInfo    any         // GPS信息
	DeviceType any         // 设备类型（1=实时视频 0=其他）
	StreamUrl  any         // 原始流地址（必填）
	PreviewImg any         // 预览图路径/URL
	CreateBy   any         // 创建者
	UpdateBy   any         // 更新者
	Remark     any         // 摄像头备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
