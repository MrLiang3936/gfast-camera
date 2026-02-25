package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// AlertLogSearchReq 告警日志列表查询请求
type AlertLogSearchReq struct {
	g.Meta            `path:"/alertLog/list" tags:"告警日志管理" method:"get" summary:"告警日志列表"`
	TaskId            *int   `p:"taskId"`      // 任务ID
	CameraId          *int   `p:"cameraId"`    // 摄像头ID
	AlgorithmId       *int   `p:"algorithmId"` // 算法ID
	EnName            string `p:"enName"`      // 报警类别
	Degree            *int   `p:"degree"`      // 报警级别
	Status            *int   `p:"status"`      // 归档状态
	commonApi.PageReq        // 分页参数
}

// AlertLogSearchRes 告警日志列表查询响应
type AlertLogSearchRes struct {
	g.Meta            `mime:"application/json"`
	List              []*entity.SysAlertLog `json:"list"` // 告警日志列表数据
	commonApi.ListRes                       // 分页结果
}

// AlertLogBaseReq 告警日志添加/编辑的基础参数
type AlertLogBaseReq struct {
	TaskId      *int   `p:"taskId"`                                    // 任务ID
	CameraId    *int   `p:"cameraId"`                                  // 摄像头ID
	AlgorithmId *int   `p:"algorithmId"`                               // 算法ID
	EnName      string `p:"enName"`                                    // 报警类别
	Degree      *int   `p:"degree"`                                    // 报警级别
	ImageUrl    string `p:"imageUrl" v:"required#图片路径不能为空"`            // 图片路径
	ImageName   string `p:"imageName"`                                 // 图片名称
	Status      int    `p:"status" v:"required|in:0,1,2#归档状态只能为0,1,2"` // 归档状态
	Remark      string `p:"remark"`                                    // 备注
}

// AlertLogAddReq 告警日志添加请求
type AlertLogAddReq struct {
	g.Meta `path:"/alertLog/add" tags:"告警日志管理" method:"post" summary:"添加告警日志"`
	*AlertLogBaseReq
}

// AlertLogAddRes 告警日志添加响应
type AlertLogAddRes struct {
}

// AlertLogGetReq 告警日志详情获取请求
type AlertLogGetReq struct {
	g.Meta  `path:"/alertLog/get" tags:"告警日志管理" method:"get" summary:"获取告警日志详情"`
	AlertId int `p:"alertId" v:"required|min:1#告警日志ID不能为空|告警日志ID必须大于0"` // 告警日志主键ID
}

// AlertLogGetRes 告警日志详情获取响应
type AlertLogGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.SysAlertLog `json:"data"` // 告警日志详情数据
}

// AlertLogEditReq 告警日志修改请求
type AlertLogEditReq struct {
	g.Meta  `path:"/alertLog/edit" tags:"告警日志管理" method:"put" summary:"修改告警日志（无用）"`
	AlertId int `p:"alertId" v:"required|min:1#告警日志ID不能为空|告警日志ID必须大于0"` // 告警日志主键ID
	*AlertLogBaseReq
}

// AlertLogEditRes 告警日志修改响应
type AlertLogEditRes struct {
}

// AlertLogDeleteReq 告警日志删除请求
type AlertLogDeleteReq struct {
	g.Meta   `path:"/alertLog/delete" tags:"告警日志管理" method:"delete" summary:"批量删除告警日志"`
	AlertIds []int `p:"alertIds" v:"required#请选择需要删除的告警日志"` // 告警日志ID数组
}

// AlertLogDeleteRes 告警日志删除响应
type AlertLogDeleteRes struct {
}

// AlertLogBatchUpdateStatusReq 告警日志批量更新状态请求
type AlertLogBatchUpdateStatusReq struct {
	g.Meta   `path:"/alertLog/batchUpdateStatus" tags:"告警日志管理" method:"put" summary:"批量更新告警日志状态"`
	AlertIds []int `p:"alertIds" v:"required#请选择需要更新的告警日志"`        // 告警日志ID数组
	Status   int   `p:"status" v:"required|in:0,1,2#归档状态只能为0,1,2"` // 目标状态
}

// AlertLogBatchUpdateStatusRes 告警日志批量更新状态响应
type AlertLogBatchUpdateStatusRes struct {
}
