// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAlgorithm is the golang structure for table sys_algorithm.
type SysAlgorithm struct {
	Id               uint        `json:"id"               orm:"id"                description:"算法配置唯一ID"`
	Intro            string      `json:"intro"            orm:"intro"             description:"算法执行逻辑描述"`
	AlgorithmId      string      `json:"algorithmId"      orm:"algorithm_id"      description:"算法模型全局唯一标识"`
	AlgorithmTaskId  string      `json:"algorithmTaskId"  orm:"algorithm_task_id" description:"算法所属任务ID"`
	AlgorithmVersion string      `json:"algorithmVersion" orm:"algorithm_version" description:"算法模型版本号"`
	CnName           string      `json:"cnName"           orm:"cn_name"           description:"算法中文名称"`
	EnName           string      `json:"enName"           orm:"en_name"           description:"算法英文名称（label）"`
	CoverImageUrl    string      `json:"coverImageUrl"    orm:"cover_image_url"   description:"算法封面图URL"`
	State            string      `json:"state"            orm:"state"             description:"算法部署状态（Installed/Uninstalled等）"`
	CreateBy         uint        `json:"createBy"         orm:"create_by"         description:"创建者"`
	UpdateBy         uint        `json:"updateBy"         orm:"update_by"         description:"更新者"`
	Remark           string      `json:"remark"           orm:"remark"            description:"备注"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:"修改时间"`
	ModelFile        string      `json:"modelFile"        orm:"model_file"        description:"加密生成的模型路径"`
	Secret           string      `json:"secret"           orm:"secret"            description:"加密"`
}
