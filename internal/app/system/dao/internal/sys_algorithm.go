// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAlgorithmDao is the data access object for the table sys_algorithm.
type SysAlgorithmDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysAlgorithmColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysAlgorithmColumns defines and stores column names for the table sys_algorithm.
type SysAlgorithmColumns struct {
	Id               string // 算法配置唯一ID
	Intro            string // 算法执行逻辑描述
	AlgorithmId      string // 算法模型全局唯一标识
	AlgorithmTaskId  string // 算法所属任务ID
	AlgorithmVersion string // 算法模型版本号
	CnName           string // 算法中文名称
	EnName           string // 算法英文名称（label）
	CoverImageUrl    string // 算法封面图URL
	State            string // 算法部署状态（Installed/Uninstalled等）
	CreateBy         string // 创建者
	UpdateBy         string // 更新者
	Remark           string // 备注
	CreatedAt        string // 创建时间
	UpdatedAt        string // 修改时间
	ModelFile        string // 加密生成的模型路径
	Secret           string
}

// sysAlgorithmColumns holds the columns for the table sys_algorithm.
var sysAlgorithmColumns = SysAlgorithmColumns{
	Id:               "id",
	Intro:            "intro",
	AlgorithmId:      "algorithm_id",
	AlgorithmTaskId:  "algorithm_task_id",
	AlgorithmVersion: "algorithm_version",
	CnName:           "cn_name",
	EnName:           "en_name",
	CoverImageUrl:    "cover_image_url",
	State:            "state",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
	Remark:           "remark",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	ModelFile:        "model_file",
	Secret:           "secret",
}

// NewSysAlgorithmDao creates and returns a new DAO object for table data access.
func NewSysAlgorithmDao(handlers ...gdb.ModelHandler) *SysAlgorithmDao {
	return &SysAlgorithmDao{
		group:    "default",
		table:    "sys_algorithm",
		columns:  sysAlgorithmColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAlgorithmDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAlgorithmDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAlgorithmDao) Columns() SysAlgorithmColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAlgorithmDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAlgorithmDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysAlgorithmDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
