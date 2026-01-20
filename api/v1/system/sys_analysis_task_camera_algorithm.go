/*
* @desc:分析任务摄像头算法关联相关参数
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/7 23:09
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type AnalysisTaskCameraAlgorithmSearchReq struct {
	g.Meta      `path:"/analysisTaskCameraAlgorithm/list" tags:"摄像头-算法关联" method:"get" summary:"分析任务摄像头算法关联列表"`
	TaskId      uint `p:"taskId"`      // 任务ID
	CameraId    uint `p:"cameraId"`    // 摄像头ID
	AlgorithmId uint `p:"algorithmId"` // 算法ID
	Page        int  `p:"page"`        // 页码
	Size        int  `p:"size"`        // 每页大小
}

type AnalysisTaskCameraAlgorithmSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []interface{} `json:"list"` // 实际应该使用具体实体
	Page   int           `json:"page"`
	Size   int           `json:"size"`
	Total  int64         `json:"total"`
}

type AnalysisTaskCameraAlgorithmAddReq struct {
	g.Meta      `path:"/analysisTaskCameraAlgorithm/add" tags:"摄像头-算法关联" method:"post" summary:"添加分析任务摄像头算法关联"`
	TaskId      uint   `p:"taskId" v:"required#任务ID不能为空"` // 任务ID
	CameraId    uint   `p:"cameraId" v:"required#摄像头ID不能为空"`
	AlgorithmId uint   `p:"algorithmId" v:"required#算法ID不能为空"`
	Remark      string `p:"remark"` // 备注
}

type AnalysisTaskCameraAlgorithmAddRes struct {
}

type AnalysisTaskCameraAlgorithmAddBatchReq struct {
	g.Meta `path:"/analysisTaskCameraAlgorithm/addPatch" tags:"摄像头-算法关联" method:"post" summary:"批量关联"`
	Data   []*entity.SysAnalysisTaskCameraAlgorithm `p:"data" v:"required#数据不能为空"`
}

type AnalysisTaskCameraAlgorithmAddBatchRes struct {
}

type AnalysisTaskCameraAlgorithmEditReq struct {
	g.Meta      `path:"/analysisTaskCameraAlgorithm/edit" tags:"摄像头-算法关联" method:"put" summary:"修改分析任务摄像头算法关联"`
	Id          uint   `p:"id" v:"required|min:1#id必须"`
	TaskId      uint   `p:"taskId" v:"required#任务ID不能为空"` // 任务ID
	CameraId    uint   `p:"cameraId" v:"required#摄像头ID不能为空"`
	AlgorithmId uint   `p:"algorithmId" v:"required#算法ID不能为空"`
	Remark      string `p:"remark"` // 备注
}

type AnalysisTaskCameraAlgorithmEditRes struct {
}

type AnalysisTaskCameraAlgorithmDeleteReq struct {
	g.Meta `path:"/analysisTaskCameraAlgorithm/delete" tags:"摄像头-算法关联" method:"delete" summary:"删除分析任务摄像头算法关联"`
	Ids    []uint `p:"ids"`
}

type AnalysisTaskCameraAlgorithmDeleteRes struct {
}
