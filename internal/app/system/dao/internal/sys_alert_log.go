// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAlertLogDao is the data access object for the table sys_alert_log.
type SysAlertLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysAlertLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysAlertLogColumns defines and stores column names for the table sys_alert_log.
type SysAlertLogColumns struct {
	AlertId     string // 主键ID
	TaskId      string // 任务ID
	CameraId    string // 摄像头ID
	AlgorithmId string // 算法ID
	EnName      string // 报警类别
	Degree      string // 报警级别
	ImageUrl    string // 图片路径
	ImageName   string // 图片名称
	Status      string // 归档状态：0-暂不归档，1-正确处理，2-错误处理
	Remark      string // 备注
	CreatedBy   string // 创建人
	UpdatedBy   string // 修改人
	CreateAt    string // 创建时间
	UpdateAt    string // 更新时间
	DeleteFlag  string
}

// sysAlertLogColumns holds the columns for the table sys_alert_log.
var sysAlertLogColumns = SysAlertLogColumns{
	AlertId:     "alert_id",
	TaskId:      "task_id",
	CameraId:    "camera_id",
	AlgorithmId: "algorithm_id",
	EnName:      "en_name",
	Degree:      "degree",
	ImageUrl:    "image_url",
	ImageName:   "image_name",
	Status:      "status",
	Remark:      "remark",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
	DeleteFlag:  "delete_flag",
}

// NewSysAlertLogDao creates and returns a new DAO object for table data access.
func NewSysAlertLogDao(handlers ...gdb.ModelHandler) *SysAlertLogDao {
	return &SysAlertLogDao{
		group:    "default",
		table:    "sys_alert_log",
		columns:  sysAlertLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAlertLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAlertLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAlertLogDao) Columns() SysAlertLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAlertLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAlertLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAlertLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
