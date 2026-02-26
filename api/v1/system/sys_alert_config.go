package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type AlertConfigSearchReq struct {
	g.Meta    `path:"/alertConfig/list" tags:"告警配置管理" method:"get" summary:"告警配置列表"`
	VoiceName string `p:"voiceName"` // 语音文件名
	VoiceUrl  string `p:"voiceUrl"`  // 语音文件URL
	IsDefault *int   `p:"isDefault"` // 是否默认
	Enable    *int   `p:"enable"`    // 是否启用
	Degree    string `p:"degree"`    // 适用告警等级
	Pop       *int   `p:"pop"`       // 是否弹窗
	PopType   string `p:"popType"`   // 弹窗类型
	commonApi.PageReq
}

type AlertConfigSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysAlertConfig `json:"list"`
}

type AlertConfigGetReq struct {
	g.Meta `path:"/alertConfig/get" tags:"告警配置管理" method:"get" summary:"告警配置列表"`
	Id     int `p:"id" v:"required#ID必须"`
}

type AlertConfigAddReq struct {
	g.Meta    `path:"/alertConfig/add" tags:"告警配置管理" method:"post" summary:"添加告警配置"`
	VoiceName string `p:"voiceName" v:"required#语音文件名不能为空"`
	VoiceUrl  string `p:"voiceUrl" v:"required#语音文件URL不能为空"`
	IsDefault int    `p:"isDefault" v:"required|in:0,1#是否默认值只能是0或1"`
	Enable    int    `p:"enable" v:"required|in:0,1#启用状态只能是0或1"`
	Degree    string `p:"degree" v:"required#适用告警等级不能为空"`
	Pop       int    `p:"pop" v:"required|in:0,1#弹窗选项只能是0或1"`
	PopType   string `p:"popType" v:"required#弹窗类型不能为空"`
	Remark    string `p:"remark"`
}

type AlertConfigAddRes struct {
}

type AlertConfigEditReq struct {
	g.Meta    `path:"/alertConfig/edit" tags:"告警配置管理" method:"put" summary:"修改告警配置"`
	Id        uint   `p:"id" v:"required#ID必须"`
	VoiceName string `p:"voiceName" v:"required#语音文件名不能为空"`
	VoiceUrl  string `p:"voiceUrl" v:"required#语音文件URL不能为空"`
	IsDefault int    `p:"isDefault" v:"required|in:0,1#是否默认值只能是0或1"`
	Enable    int    `p:"enable" v:"required|in:0,1#启用状态只能是0或1"`
	Degree    string `p:"degree" v:"required#适用告警等级不能为空"`
	Pop       int    `p:"pop" v:"required|in:0,1#弹窗选项只能是0或1"`
	PopType   string `p:"popType" v:"required#弹窗类型不能为空"`
	Remark    string `p:"remark"`
}

type AlertConfigEditRes struct {
}

type AlertConfigDeleteReq struct {
	g.Meta `path:"/alertConfig/delete" tags:"告警配置管理" method:"delete" summary:"删除告警配置"`
	Ids    []uint `p:"ids"`
}

type AlertConfigDeleteRes struct {
}
