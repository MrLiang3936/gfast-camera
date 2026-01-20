// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAnalysisTaskCameraAlgorithmDao is the data access object for the table sys_analysis_task_camera_algorithm.
type SysAnalysisTaskCameraAlgorithmDao struct {
	table    string                                // table is the underlying table name of the DAO.
	group    string                                // group is the database configuration group name of the current DAO.
	columns  SysAnalysisTaskCameraAlgorithmColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler                    // handlers for customized model modification.
}

// SysAnalysisTaskCameraAlgorithmColumns defines and stores column names for the table sys_analysis_task_camera_algorithm.
type SysAnalysisTaskCameraAlgorithmColumns struct {
	Id          string // 自增ID
	TaskId      string // 摄像头ID
	CameraId    string // 摄像头ID
	AlgorithmId string // 算法ID
	Remark      string // 备注
	CreateBy    string // 创建者
	UpdateBy    string // 更新者
	CreatedTime string // 任务创建时间
	UpdatedTime string // 任务更新时间
}

// sysAnalysisTaskCameraAlgorithmColumns holds the columns for the table sys_analysis_task_camera_algorithm.
var sysAnalysisTaskCameraAlgorithmColumns = SysAnalysisTaskCameraAlgorithmColumns{
	Id:          "id",
	TaskId:      "task_id",
	CameraId:    "camera_id",
	AlgorithmId: "algorithm_id",
	Remark:      "remark",
	CreateBy:    "create_by",
	UpdateBy:    "update_by",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// NewSysAnalysisTaskCameraAlgorithmDao creates and returns a new DAO object for table data access.
func NewSysAnalysisTaskCameraAlgorithmDao(handlers ...gdb.ModelHandler) *SysAnalysisTaskCameraAlgorithmDao {
	return &SysAnalysisTaskCameraAlgorithmDao{
		group:    "default",
		table:    "sys_analysis_task_camera_algorithm",
		columns:  sysAnalysisTaskCameraAlgorithmColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAnalysisTaskCameraAlgorithmDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAnalysisTaskCameraAlgorithmDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAnalysisTaskCameraAlgorithmDao) Columns() SysAnalysisTaskCameraAlgorithmColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAnalysisTaskCameraAlgorithmDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAnalysisTaskCameraAlgorithmDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAnalysisTaskCameraAlgorithmDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
