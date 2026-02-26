/*
* @desc:请求zlMediaKit参数结构体
* @company:
* @Author: liangyu
* @Date:   2026/1/13 11:47
 */

package model

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// SysAlertLog 报警日志拓展信息
type SysAlertLog struct {
	entity.SysAlertLog
	Id         string `json:"id"       orm:"id"     description:"算法中文名称"`
	TaskName   string `json:"taskName"       orm:"task_name"     description:"任务名称"`
	CnName     string `json:"cnName"     orm:"cn_name"     description:"算法中文名称"`
	EnName     string `json:"enName"     orm:"en_name"     description:"算法英文名称（label）"`
	CameraName string `json:"cameraName" orm:"camera_name" description:"摄像头名称（必填）"`
}
