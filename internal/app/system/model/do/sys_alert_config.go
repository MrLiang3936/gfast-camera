// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAlertConfig is the golang structure of table sys_alert_config for DAO operations like Where/Data.
type SysAlertConfig struct {
	g.Meta    `orm:"table:sys_alert_config, do:true"`
	Id        any         // 主键ID
	VoiceName any         // 语音文件名（含后缀）
	VoiceUrl  any         // 语音文件存储URL/路径
	IsDefault any         // 是否默认配置：0-否，1-是
	Enable    any         // 是否启用：0-禁用，1-启用
	Degree    any         // 适用告警等级，如1,2,3
	Pop       any         // 是否弹窗：0-否，1-是
	PopType   any         // 弹窗类型：right_small-右侧小窗，full_screen-全屏等
	CreatedBy any         // 创建人
	UpdatedBy any         // 修改人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
}
