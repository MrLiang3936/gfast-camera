// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAlgorithm is the golang structure of table sys_algorithm for DAO operations like Where/Data.
type SysAlgorithm struct {
	g.Meta           `orm:"table:sys_algorithm, do:true"`
	Id               any         // 算法配置唯一ID
	Intro            any         // 算法执行逻辑描述
	AlgorithmId      any         // 算法模型全局唯一标识
	AlgorithmTaskId  any         // 算法所属任务ID
	AlgorithmVersion any         // 算法模型版本号
	CnName           any         // 算法中文名称
	EnName           any         // 算法英文名称（label）
	CoverImageUrl    any         // 算法封面图URL
	State            any         // 算法部署状态（Installed/Uninstalled等）
	CreateBy         any         // 创建者
	UpdateBy         any         // 更新者
	Remark           any         // 备注
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 修改时间
	ModelFile        any         // 加密生成的模型路径
	Secret           any
}
