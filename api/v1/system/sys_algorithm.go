package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// AlgorithmSearchReq 算法列表查询请求
type AlgorithmSearchReq struct {
	g.Meta            `path:"/algorithm/list" tags:"算法管理" method:"get" summary:"算法列表"`
	CnName            string `p:"cnName"`          // 算法中文名称
	EnName            string `p:"enName"`          // 算法英文名称
	AlgorithmId       string `p:"algorithmId"`     // 算法模型全局唯一标识
	AlgorithmTaskId   string `p:"algorithmTaskId"` // 算法所属任务ID
	State             string `p:"state"`           // 算法部署状态
	commonApi.PageReq        // 分页参数
}

// AlgorithmSearchRes 算法列表查询响应
type AlgorithmSearchRes struct {
	g.Meta            `mime:"application/json"`
	List              []*entity.SysAlgorithm `json:"list"` // 算法列表数据
	commonApi.ListRes                        // 分页结果
}

// AlgorithmBaseReq 算法添加/编辑的基础参数
type AlgorithmBaseReq struct {
	Intro            string `p:"intro" v:"required#算法执行逻辑描述不能为空"`           // 算法执行逻辑描述
	AlgorithmId      string `p:"algorithmId" v:"required#算法模型全局唯一标识不能为空"`   // 算法模型全局唯一标识
	AlgorithmTaskId  string `p:"algorithmTaskId" v:"required#算法所属任务ID不能为空"` // 算法所属任务ID
	AlgorithmVersion string `p:"algorithmVersion"`                          // 算法模型版本号
	CnName           string `p:"cnName" v:"required#算法中文名称不能为空"`            // 算法中文名称
	EnName           string `p:"enName" v:"required#算法英文名称不能为空"`            // 算法英文名称
	CoverImageUrl    string `p:"coverImageUrl"`                             // 算法封面图URL
	State            string `p:"state" v:"required#算法部署状态不能为空"`             // 算法部署状态
	Remark           string `p:"remark"`                                    // 备注
	ModelFile        string `p:"modelFile"`                                 // 加密生成的模型路径
	Secret           string `p:"secret"`
}

// AlgorithmAddReq 算法添加请求
type AlgorithmAddReq struct {
	g.Meta `path:"/algorithm/add" tags:"算法管理" method:"post" summary:"添加算法"`
	*AlgorithmBaseReq
}

// AlgorithmAddRes 算法添加响应
type AlgorithmAddRes struct {
	Id int `json:"id"` // 新增算法ID
}

// AlgorithmGetReq 算法详情获取请求
type AlgorithmGetReq struct {
	g.Meta `path:"/algorithm/get" tags:"算法管理" method:"get" summary:"获取算法详情"`
	Id     int `p:"id" v:"required|min:1#算法ID不能为空|算法ID必须大于0"` // 算法主键ID
}

// AlgorithmGetRes 算法详情获取响应
type AlgorithmGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.SysAlgorithm `json:"data"` // 算法详情数据
}

// AlgorithmEditReq 算法修改请求
type AlgorithmEditReq struct {
	g.Meta      `path:"/algorithm/edit" tags:"算法管理" method:"put" summary:"修改算法"`
	AlgorithmId int `p:"algorithmId" v:"required|min:1#算法ID不能为空|算法ID必须大于0"` // 算法主键ID
	*AlgorithmBaseReq
}

// AlgorithmEditRes 算法修改响应
type AlgorithmEditRes struct {
}

// AlgorithmDeleteReq 算法删除请求
type AlgorithmDeleteReq struct {
	g.Meta `path:"/algorithm/delete" tags:"算法管理" method:"delete" summary:"删除算法"`
	Ids    []int `p:"ids" v:"required#请选择需要删除的算法"` // 算法ID数组
}

// AlgorithmDeleteRes 算法删除响应
type AlgorithmDeleteRes struct {
}
