package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// CameraSearchReq 摄像头列表查询请求
type CameraSearchReq struct {
	g.Meta            `path:"/camera/list" tags:"摄像头管理" method:"get" summary:"摄像头列表"`
	CameraName        string `p:"cameraName"` // 摄像头名称
	GroupId           int    `p:"groupId"`    // 所属分组ID
	DeviceType        int    `p:"deviceType"` // 设备类型（1=实时视频 0=其他）
	commonApi.PageReq        // 分页参数
}

// CameraSearchRes 摄像头列表查询响应
type CameraSearchRes struct {
	g.Meta            `mime:"application/json"`
	List              []*entity.SysCamera `json:"list"` // 摄像头列表数据
	commonApi.ListRes                     // 分页结果
}

// CameraBaseReq 摄像头添加/编辑的基础参数
type CameraBaseReq struct {
	CameraName string `p:"cameraName"  v:"required#摄像头名称不能为空"`
	GroupId    int    `p:"groupId"` // 所属分组主键
	GpsInfo    string `p:"gpsInfo"` // GPS信息
	DeviceType int    `p:"deviceType"  v:"required|in:0,1#设备类型不能为空|设备类型只能为0或1"`
	StreamUrl  string `p:"streamUrl"   v:"required#原始流地址不能为空"`
	PreviewImg string `p:"previewImg"` // 预览图路径/URL
	Remark     string `p:"remark"`     // 摄像头备注
}

// CameraAddReq 摄像头添加请求
type CameraAddReq struct {
	g.Meta `path:"/camera/add" tags:"摄像头管理" method:"post" summary:"添加摄像头"`
	*CameraBaseReq
}

// CameraAddRes 摄像头添加响应
type CameraAddRes struct {
}

// CameraGetReq 摄像头详情获取请求
type CameraGetReq struct {
	g.Meta `path:"/camera/get" tags:"摄像头管理" method:"get" summary:"获取摄像头详情"`
	Id     int `p:"id" v:"required|min:1#摄像头ID不能为空|摄像头ID必须大于0"` // 摄像头主键ID
}

// CameraGetRes 摄像头详情获取响应
type CameraGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.SysCamera `json:"data"` // 摄像头详情数据
}

// CameraEditReq 摄像头修改请求
type CameraEditReq struct {
	g.Meta   `path:"/camera/edit" tags:"摄像头管理" method:"put" summary:"修改摄像头"`
	CameraId int `p:"cameraId" v:"required|min:1#摄像头ID不能为空|摄像头ID必须大于0"` // 摄像头主键ID
	*CameraBaseReq
}

// CameraEditRes 摄像头修改响应
type CameraEditRes struct {
}

// CameraDeleteReq 摄像头删除请求
type CameraDeleteReq struct {
	g.Meta `path:"/camera/delete" tags:"摄像头管理" method:"delete" summary:"删除摄像头"`
	Ids    []int `p:"ids" v:"required#请选择需要删除的摄像头"` // 摄像头ID数组
}

// CameraDeleteRes 摄像头删除响应
type CameraDeleteRes struct {
}
