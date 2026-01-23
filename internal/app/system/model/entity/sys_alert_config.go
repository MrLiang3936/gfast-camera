// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAlertConfig is the golang structure for table sys_alert_config.
type SysAlertConfig struct {
	Id        uint        `json:"id"        orm:"id"         description:"主键ID"`
	VoiceName string      `json:"voiceName" orm:"voice_name" description:"语音文件名（含后缀）"`
	VoiceUrl  string      `json:"voiceUrl"  orm:"voice_url"  description:"语音文件存储URL/路径"`
	IsDefault int         `json:"isDefault" orm:"is_default" description:"是否默认配置：0-否，1-是"`
	Enable    int         `json:"enable"    orm:"enable"     description:"是否启用：0-禁用，1-启用"`
	Degree    string      `json:"degree"    orm:"degree"     description:"适用告警等级，如1,2,3"`
	Pop       int         `json:"pop"       orm:"pop"        description:"是否弹窗：0-否，1-是"`
	PopType   string      `json:"popType"   orm:"pop_type"   description:"弹窗类型：right_small-右侧小窗，full_screen-全屏等"`
	CreatedBy uint64      `json:"createdBy" orm:"created_by" description:"创建人"`
	UpdatedBy uint64      `json:"updatedBy" orm:"updated_by" description:"修改人"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`
}
