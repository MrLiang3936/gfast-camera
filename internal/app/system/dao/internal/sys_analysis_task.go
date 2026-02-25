// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAnalysisTaskDao is the data access object for the table sys_analysis_task.
type SysAnalysisTaskDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  SysAnalysisTaskColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// SysAnalysisTaskColumns defines and stores column names for the table sys_analysis_task.
type SysAnalysisTaskColumns struct {
	Id           string // 任务自增ID
	Name         string // 任务名称（如：办公室行为观察）
	State        string // 任务运行状态（run/stop等）
	WorkTimeType string // 任务执行时间类型（everyday/week指定等）
	WorkTime     string // 任务执行时间配置，包含weekdays（周几）和periods（时间段）
	Type         string // 任务类型（video_analysis：视频分析）
	Remark       string // 备注
	CreateBy     string // 创建者
	UpdateBy     string // 更新者
	CreatedTime  string // 任务创建时间
	UpdatedTime  string // 任务更新时间
	AlertFreq    string // 报警频率 秒 (多少秒报警一次)
	AlertKeep    string // 报警持续 秒 (出现多少秒就算)
}

// sysAnalysisTaskColumns holds the columns for the table sys_analysis_task.
var sysAnalysisTaskColumns = SysAnalysisTaskColumns{
	Id:           "id",
	Name:         "name",
	State:        "state",
	WorkTimeType: "work_time_type",
	WorkTime:     "work_time",
	Type:         "type",
	Remark:       "remark",
	CreateBy:     "create_by",
	UpdateBy:     "update_by",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
	AlertFreq:    "alert_freq",
	AlertKeep:    "alert_keep",
}

// NewSysAnalysisTaskDao creates and returns a new DAO object for table data access.
func NewSysAnalysisTaskDao(handlers ...gdb.ModelHandler) *SysAnalysisTaskDao {
	return &SysAnalysisTaskDao{
		group:    "default",
		table:    "sys_analysis_task",
		columns:  sysAnalysisTaskColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAnalysisTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAnalysisTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAnalysisTaskDao) Columns() SysAnalysisTaskColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAnalysisTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAnalysisTaskDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAnalysisTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
