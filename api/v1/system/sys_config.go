/*
* @desc:系统参数配置
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/18 21:11
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	commonEntity "github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
)

type ConfigSearchReq struct {
	g.Meta     `path:"/config/list" tags:"系统参数管理" method:"get" summary:"系统参数列表"`
	ConfigName string `p:"configName"` //参数名称
	ConfigKey  string `p:"configKey"`  //参数键名
	ConfigType string `p:"configType"` //状态
	commonApi.PageReq
}

type ConfigSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*commonEntity.SysConfig `json:"list"`
	commonApi.ListRes
}

type ConfigReq struct {
	ConfigName  string `p:"configName"  v:"required#参数名称不能为空"`
	ConfigKey   string `p:"configKey"  v:"required#参数键名不能为空"`
	ConfigValue string `p:"configValue"  v:"required#参数键值不能为空"`
	ConfigType  int    `p:"configType"    v:"required|in:0,1#系统内置不能为空|系统内置类型只能为0或1"`
	Remark      string `p:"remark"`
}

type ConfigAddReq struct {
	g.Meta `path:"/config/add" tags:"系统参数管理" method:"post" summary:"添加系统参数"`
	*ConfigReq
}

type ConfigAddRes struct {
}

type ConfigGetReq struct {
	g.Meta `path:"/config/get" tags:"系统参数管理" method:"get" summary:"获取系统参数"`
	Id     int `p:"id"`
}

type ConfigGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *commonEntity.SysConfig `json:"data"`
}

type ConfigEditReq struct {
	g.Meta   `path:"/config/edit" tags:"系统参数管理" method:"put" summary:"修改系统参数"`
	ConfigId int64 `p:"configId" v:"required|min:1#主键ID不能为空|主键ID参数错误"`
	*ConfigReq
}

type ConfigEditRes struct {
}

type ConfigDeleteReq struct {
	g.Meta `path:"/config/delete" tags:"系统参数管理" method:"delete" summary:"删除系统参数"`
	Ids    []int `p:"ids"`
}

type ConfigDeleteRes struct {
}

// /////////////////////////////////////////////////////////////////////////
// 网络配置相关请求响应结构体
type NetworkConfigReq struct {
	NicName       string `p:"nicName" v:"required#网卡名称不能为空"`                              // 网卡名称
	ConnName      string `p:"connName" v:"required#连接名称不能为空"`                             // 连接名称
	IpType        int    `p:"ipType" v:"required|in:0,1#IP获取方式不能为空|IP获取方式只能为0(自动)或1(静态)"` // 0:DHCP自动获取 1:静态配置
	IpAddress     string `p:"ipAddress" v:"required-if:ipType,1|ip#静态IP地址不能为空|IP地址格式不正确"` // IP地址（静态模式必填）
	SubnetMask    string `p:"subnetMask" v:"required-if:ipType,1#子网掩码不能为空|IP地址格式不正确"`     // 子网掩码（静态模式必填）
	Gateway       string `p:"gateway" v:"ip#网关格式不正确"`                                     // 默认网关（可选）
	DnsServer     string `p:"dnsServer" v:"ip#DNS服务器格式不正确(多个用逗号分隔)"`                      // DNS服务器（可选，多个用逗号分隔）
	RoutingWeight int    `p:"routingWeight" v:"min:0#路由权重不能小于0"`                          // 路由权重（metric值，0-255，数值越小优先级越高）
}

type NetworkConfigGetReq struct {
	g.Meta  `path:"/config/network/get" tags:"系统参数管理" method:"get" summary:"获取网络配置"`
	NicName string `p:"nicName" v:"required#网卡名称不能为空"`
}

type NetworkConfigGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *NetworkConfigReq `json:"data"`
}

type NetworkConfigSaveReq struct {
	g.Meta `path:"/config/network/save" tags:"系统参数管理" method:"post" summary:"保存网络配置"`
	NetworkConfigReq
}

type NetworkConfigSaveRes struct {
	g.Meta  `mime:"application/json"`
	Success bool `json:"success"`
}

type NetworkListReq struct {
	g.Meta `path:"/config/network/list" tags:"系统参数管理" method:"get" summary:"获取网卡列表"`
}

type NetworkListRes struct {
	g.Meta `mime:"application/json"`
	List   []*NicInfo `json:"list"`
}

type NicInfo struct {
	Name    string `json:"name"`    // 网卡名称
	MacAddr string `json:"macAddr"` // MAC地址
	Status  string `json:"status"`  // 状态
	IpAddr  string `json:"ipAddr"`  // 当前IP地址
}
