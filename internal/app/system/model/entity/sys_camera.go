// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCamera is the golang structure for table sys_camera.
type SysCamera struct {
	CameraId   uint        `json:"cameraId"   orm:"camera_id"   description:"摄像头主键"`
	CameraName string      `json:"cameraName" orm:"camera_name" description:"摄像头名称（必填）"`
	GroupId    uint        `json:"groupId"    orm:"group_id"    description:"所属分组主键（关联sys_camera_group表）"`
	GpsInfo    string      `json:"gpsInfo"    orm:"gps_info"    description:"GPS信息"`
	DeviceType int         `json:"deviceType" orm:"device_type" description:"设备类型（1=实时视频 0=其他）"`
	StreamUrl  string      `json:"streamUrl"  orm:"stream_url"  description:"原始流地址（必填）"`
	PreviewImg string      `json:"previewImg" orm:"preview_img" description:"预览图路径/URL"`
	CreateBy   uint        `json:"createBy"   orm:"create_by"   description:"创建者"`
	UpdateBy   uint        `json:"updateBy"   orm:"update_by"   description:"更新者"`
	Remark     string      `json:"remark"     orm:"remark"      description:"摄像头备注"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"修改时间"`
}
