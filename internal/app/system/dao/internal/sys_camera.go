// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCameraDao is the data access object for the table sys_camera.
type SysCameraDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysCameraColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysCameraColumns defines and stores column names for the table sys_camera.
type SysCameraColumns struct {
	CameraId   string // 摄像头主键
	CameraName string // 摄像头名称（必填）
	GroupId    string // 所属分组主键（关联sys_camera_group表）
	GpsInfo    string // GPS信息
	DeviceType string // 设备类型（1=实时视频 0=其他）
	StreamUrl  string // 原始流地址（必填）
	PreviewImg string // 预览图路径/URL
	CreateBy   string // 创建者
	UpdateBy   string // 更新者
	Remark     string // 摄像头备注
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// sysCameraColumns holds the columns for the table sys_camera.
var sysCameraColumns = SysCameraColumns{
	CameraId:   "camera_id",
	CameraName: "camera_name",
	GroupId:    "group_id",
	GpsInfo:    "gps_info",
	DeviceType: "device_type",
	StreamUrl:  "stream_url",
	PreviewImg: "preview_img",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSysCameraDao creates and returns a new DAO object for table data access.
func NewSysCameraDao(handlers ...gdb.ModelHandler) *SysCameraDao {
	return &SysCameraDao{
		group:    "default",
		table:    "sys_camera",
		columns:  sysCameraColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCameraDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCameraDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCameraDao) Columns() SysCameraColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCameraDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCameraDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysCameraDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
