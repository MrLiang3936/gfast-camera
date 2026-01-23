// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAlertConfigDao is the data access object for the table sys_alert_config.
type SysAlertConfigDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SysAlertConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SysAlertConfigColumns defines and stores column names for the table sys_alert_config.
type SysAlertConfigColumns struct {
	Id        string // 主键ID
	VoiceName string // 语音文件名（含后缀）
	VoiceUrl  string // 语音文件存储URL/路径
	IsDefault string // 是否默认配置：0-否，1-是
	Enable    string // 是否启用：0-禁用，1-启用
	Degree    string // 适用告警等级，如1,2,3
	Pop       string // 是否弹窗：0-否，1-是
	PopType   string // 弹窗类型：right_small-右侧小窗，full_screen-全屏等
	CreatedBy string // 创建人
	UpdatedBy string // 修改人
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
}

// sysAlertConfigColumns holds the columns for the table sys_alert_config.
var sysAlertConfigColumns = SysAlertConfigColumns{
	Id:        "id",
	VoiceName: "voice_name",
	VoiceUrl:  "voice_url",
	IsDefault: "is_default",
	Enable:    "enable",
	Degree:    "degree",
	Pop:       "pop",
	PopType:   "pop_type",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

// NewSysAlertConfigDao creates and returns a new DAO object for table data access.
func NewSysAlertConfigDao(handlers ...gdb.ModelHandler) *SysAlertConfigDao {
	return &SysAlertConfigDao{
		group:    "default",
		table:    "sys_alert_config",
		columns:  sysAlertConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAlertConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAlertConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAlertConfigDao) Columns() SysAlertConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAlertConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAlertConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAlertConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
