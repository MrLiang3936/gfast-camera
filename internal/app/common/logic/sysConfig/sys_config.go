package sysConfig

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/common/consts"
	"github.com/tiger1103/gfast/v3/internal/app/common/dao"
	"github.com/tiger1103/gfast/v3/internal/app/common/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysConfig(New())
}

func New() *sSysConfig {
	return &sSysConfig{}
}

type sSysConfig struct {
}

// List 系统参数列表
func (s *sSysConfig) List(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error) {
	res = new(system.ConfigSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysConfig.Ctx(ctx)
		if req != nil {
			if req.ConfigName != "" {
				m = m.Where("config_name like ?", "%"+req.ConfigName+"%")
			}
			if req.ConfigType != "" {
				m = m.Where("config_type = ", gconv.Int(req.ConfigType))
			}
			if req.ConfigKey != "" {
				m = m.Where("config_key like ?", "%"+req.ConfigKey+"%")
			}
			if len(req.DateRange) > 0 {
				m = m.Where("created_at >= ? AND created_at<=?", req.DateRange[0], req.DateRange[1])
			}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		err = m.Page(req.PageNum, req.PageSize).Order("config_id asc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sSysConfig) Add(ctx context.Context, req *system.ConfigAddReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.CheckConfigKeyUnique(ctx, req.ConfigKey)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			CreateBy:    userId,
			Remark:      req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加系统参数失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// CheckConfigKeyUnique 验证参数键名是否存在
func (s *sSysConfig) CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := (*entity.SysConfig)(nil)
		m := dao.SysConfig.Ctx(ctx).Fields(dao.SysConfig.Columns().ConfigId).Where(dao.SysConfig.Columns().ConfigKey, configKey)
		if len(configId) > 0 {
			m = m.Where(dao.SysConfig.Columns().ConfigId+" != ?", configId[0])
		}
		err = m.Scan(&data)
		liberr.ErrIsNil(ctx, err, "校验失败")
		if data != nil {
			liberr.ErrIsNil(ctx, errors.New("参数键名重复"))
		}
	})
	return
}

// Get 获取系统参数
func (s *sSysConfig) Get(ctx context.Context, id int) (res *system.ConfigGetRes, err error) {
	res = new(system.ConfigGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysConfig.Ctx(ctx).WherePri(id).Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "获取系统参数失败")
	})
	return
}

// Edit 修改系统参数
func (s *sSysConfig) Edit(ctx context.Context, req *system.ConfigEditReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.CheckConfigKeyUnique(ctx, req.ConfigKey, req.ConfigId)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.SysConfig.Ctx(ctx).WherePri(req.ConfigId).Update(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			UpdateBy:    userId,
			Remark:      req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "修改系统参数失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// Delete 删除系统参数
func (s *sSysConfig) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysConfig.Ctx(ctx).Delete(dao.SysConfig.Columns().ConfigId+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// GetConfigByKey 通过key获取参数（从缓存获取）
func (s *sSysConfig) GetConfigByKey(ctx context.Context, key string) (config *entity.SysConfig, err error) {
	if key == "" {
		err = gerror.New("参数key不能为空")
		return
	}
	cache := service.Cache()
	cf := cache.Get(ctx, consts.CacheSysConfigTag+key)
	if cf != nil && !cf.IsEmpty() {
		err = gconv.Struct(cf, &config)
		return
	}
	config, err = s.GetByKey(ctx, key)
	if err != nil {
		return
	}
	if config != nil {
		cache.Set(ctx, consts.CacheSysConfigTag+key, config, 0, consts.CacheSysConfigTag)
	}
	return
}

// GetByKey 通过key获取参数（从数据库获取）
func (s *sSysConfig) GetByKey(ctx context.Context, key string) (config *entity.SysConfig, err error) {
	err = dao.SysConfig.Ctx(ctx).Where("config_key", key).Scan(&config)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取配置失败")
	}
	return
}

// ///////////////////////////////////////////
// GetNetworkList 获取网卡列表（兼容Linux和Windows）
func (s *sSysConfig) GetNetworkList(ctx context.Context) (res *system.NetworkListRes, err error) {
	res = new(system.NetworkListRes)
	res.List = make([]*system.NicInfo, 0)

	// 获取系统所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return res, gerror.Wrap(err, "获取网络接口列表失败")
	}

	for _, iface := range interfaces {
		nicInfo := &system.NicInfo{
			Name:    iface.Name,
			MacAddr: iface.HardwareAddr.String(),
			Status:  "down",
		}

		// 判断网卡状态（是否启用）
		if iface.Flags&net.FlagUp != 0 {
			nicInfo.Status = "up"
		}

		// 获取网卡绑定的IP地址（只取IPv4地址，排除回环地址）
		addrs, err := iface.Addrs()
		if err == nil {
			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					nicInfo.IpAddr = ipNet.IP.String()
					break // 只取第一个有效IPv4地址
				}
			}
		}

		res.List = append(res.List, nicInfo)
	}

	return res, nil
}

// GetNetworkConfig 获取网络配置
func (s *sSysConfig) GetNetworkConfig(ctx context.Context, nicName string) (res *system.NetworkConfigGetRes, err error) {
	res = new(system.NetworkConfigGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		// 从数据库查询配置
		configKey := fmt.Sprintf("network.%s.config", nicName)
		config, err := s.GetByKey(ctx, configKey)
		liberr.ErrIsNil(ctx, err, "获取网络配置失败")

		// 若存在配置则解析，否则返回默认值
		if config != nil {
			err = gjson.DecodeTo(config.ConfigValue, &res.Data)
			liberr.ErrIsNil(ctx, err, "解析网络配置失败")
		} else {
			// 默认值：自动获取IP
			res.Data = &system.NetworkConfigReq{
				NicName:       nicName,
				ConnName:      fmt.Sprintf("%s 连接", nicName),
				IpType:        0,  // 0:DHCP
				RoutingWeight: 10, // 默认权重
			}
		}
	})
	return
}

// SaveNetworkConfig 保存并应用网络配置
func (s *sSysConfig) SaveNetworkConfig(ctx context.Context, req *system.NetworkConfigSaveReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 0. 参数非空校验
		if req == nil {
			liberr.ErrIsNil(ctx, errors.New("请求参数不能为空"), "请求参数不能为空")
		}

		// 在访问 req 的任何属性之前，确保执行了非空校验
		// 1. 参数合法性校验（补充业务校验）
		if req.IpType == 1 {
			// 静态IP模式：校验IP+子网掩码是否匹配
			// 检查子网掩码格式，如果是CIDR格式（如24）则转换为标准格式（如255.255.255.0）
			//cidrMask := req.SubnetMask
			//if !strings.Contains(req.SubnetMask, ".") {
			//	// CIDR格式，转换为标准格式
			//	cidrMask = s.cidrToSubnetMask(req.SubnetMask)
			//}
			_, _, err := net.ParseCIDR(fmt.Sprintf("%s/%s", req.IpAddress, req.SubnetMask))
			if err != nil {
				liberr.ErrIsNil(ctx, errors.New(req.IpAddress), "IP地址与子网掩码不匹配")
			}
			// 校验网关（若存在）
			if req.Gateway != "" && net.ParseIP(req.Gateway) == nil {
				liberr.ErrIsNil(ctx, errors.New("网关地址格式不正确"), "网关地址格式不正确")
			}
		}

		// 2. 保存配置到数据库
		configKey := fmt.Sprintf("network.%s.config", req.NicName)

		// 序列化配置
		configValue, err := gjson.EncodeString(req)
		liberr.ErrIsNil(ctx, err, "序列化网络配置失败")

		// 检查配置是否存在
		existConfig, err := s.GetByKey(ctx, configKey)
		liberr.ErrIsNil(ctx, err)

		if existConfig != nil {
			// 更新配置
			_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, configKey).Update(do.SysConfig{
				ConfigName:  fmt.Sprintf("%s 网络配置", req.NicName),
				ConfigValue: configValue,
				UpdateBy:    userId,
			})
		} else {
			// 新增配置
			_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{
				ConfigName:  fmt.Sprintf("%s 网络配置", req.NicName),
				ConfigKey:   configKey,
				ConfigValue: configValue,
				ConfigType:  0, // 非系统内置
				CreateBy:    userId,
			})
		}
		liberr.ErrIsNil(ctx, err, "保存网络配置失败")

		// 3. 应用配置到系统
		err = s.applyNetworkConfig(ctx, req)
		liberr.ErrIsNil(ctx, err, "应用网络配置失败") //"应用网络配置失败")

		// 4. 清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// applyNetworkConfig 应用网络配置到系统（兼容Linux/Windows）
func (s *sSysConfig) applyNetworkConfig(ctx context.Context, req *system.NetworkConfigSaveReq) error {
	nicName := req.NicName
	switch runtime.GOOS {
	case "linux":
		return s.applyLinuxConfig(ctx, nicName, req)
	case "windows":
		return s.applyWindowsConfig(ctx, nicName, req)
	default:
		return gerror.Newf("不支持的操作系统：%s", runtime.GOOS)
	}
}

// applyLinuxConfig Linux系统应用网络配置（基于netplan/ip命令）
func (s *sSysConfig) applyLinuxConfig(ctx context.Context, nic string, req *system.NetworkConfigSaveReq) error {
	// 1. 禁用网卡
	//if err := exec.Command("ip", "link", "set", nic, "down").Run(); err != nil {
	//	return gerror.Newf("禁用网卡 %s 失败: %v", nic, err)
	//}

	// 2. 根据IP类型配置
	if req.IpType == 1 { // 静态IP
		// 清空现有IP
		//if err := exec.Command("ip", "addr", "flush", "dev", nic).Run(); err != nil {
		//	return gerror.Newf("清空网卡IP失败: %v", err)
		//}

		// 设置IP和子网掩码
		subnetMask := req.SubnetMask
		// 如果子网掩码是CIDR格式（如24），转换为标准格式（如255.255.255.0）
		if !strings.Contains(req.SubnetMask, ".") {
			subnetMask = s.cidrToSubnetMask(req.SubnetMask)
		}
		ipWithMask := fmt.Sprintf("%s/%s", req.IpAddress, subnetMask)
		cmd := exec.Command("ip", "addr", "add", ipWithMask, "dev", nic)
		g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
		if err := cmd.Run(); err != nil {
			return gerror.Newf("设置静态IP失败: %v", err)
		}

		// 设置网关（若存在）
		if req.Gateway != "" {
			// 删除现有默认网关
			cmd := exec.Command("ip", "route", "del", "default", "dev", nic)
			g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
			cmd.Run()
			// 添加新网关（带权重）
			cmd = exec.Command(
				"ip", "route", "add", "default",
				"via", req.Gateway,
				"dev", nic,
				"metric", fmt.Sprintf("%d", req.RoutingWeight),
			)
			g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
			if err := cmd.Run(); err != nil {
				return gerror.Newf("设置网关失败: %v", err)
			}
		}

		// 配置DNS（临时生效，持久化需修改netplan）
		if req.DnsServer != "" {
			dnsContent := ""
			for _, dns := range strings.Split(req.DnsServer, ",") {
				dnsContent += fmt.Sprintf("nameserver %s\n", strings.TrimSpace(dns))
			}
			// 注意：生产环境建议通过netplan配置DNS（/etc/netplan/*.yaml）
			if err := gfile.PutContents("/etc/resolv.conf", dnsContent); err != nil {
				return gerror.Newf("配置DNS失败: %v", err)
			}
		}
	} else { // DHCP自动获取
		// 释放旧IP
		cmd := exec.Command("dhclient", "-r", nic)
		g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
		cmd.Run()
		// 获取新IP
		cmd = exec.Command("dhclient", nic)
		g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
		if err := cmd.Run(); err != nil {
			return gerror.Newf("DHCP获取IP失败: %v", err)
		}
	}

	// 3. 启用网卡
	cmd := exec.Command("ip", "link", "set", nic, "up")
	g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
	if err := cmd.Run(); err != nil {
		return gerror.Newf("启用网卡 %s 失败: %v", nic, err)
	}

	return nil
}

// applyWindowsConfig Windows系统应用网络配置（基于netsh命令）
func (s *sSysConfig) applyWindowsConfig(ctx context.Context, nic string, req *system.NetworkConfigSaveReq) error {
	// 获取网卡连接名称（Windows需要用连接名而非接口名）
	connName, err := s.getWindowsConnName(nic)
	if err != nil {
		return err
	}

	// 重置IP配置
	cmd := exec.Command("netsh", "interface", "ipv4", "reset")
	g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
	cmd.Run()

	if req.IpType == 1 { // 静态IP
		// 设置IP和子网掩码
		subnetMask := req.SubnetMask
		// 如果子网掩码是CIDR格式（如24），转换为标准格式（如255.255.255.0）
		if !strings.Contains(req.SubnetMask, ".") {
			subnetMask = s.cidrToSubnetMask(req.SubnetMask)
		}
		cmd := exec.Command(
			"netsh", "interface", "ipv4", "set", "address",
			"name="+connName,
			"static", req.IpAddress, subnetMask, req.Gateway,
			fmt.Sprintf("%d", req.RoutingWeight),
		)
		g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
		if err := cmd.Run(); err != nil {
			return gerror.Newf("设置静态IP失败: %v", err)
		}

		// 配置DNS
		if req.DnsServer != "" {
			dnsServers := strings.ReplaceAll(req.DnsServer, ",", " ")
			cmd = exec.Command(
				"netsh", "interface", "ipv4", "set", "dns",
				"name="+connName,
				"static", dnsServers,
			)
			g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
			if err := cmd.Run(); err != nil {
				return gerror.Newf("配置DNS失败: %v", err)
			}
		}
	} else { // DHCP自动获取
		// 设置IP自动获取
		cmd := exec.Command(
			"netsh", "interface", "ipv4", "set", "address",
			"name="+connName,
			"dhcp",
		)
		g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
		if err := cmd.Run(); err != nil {
			return gerror.Newf("启用DHCP失败: %v", err)
		}

		// 设置DNS自动获取
		cmd = exec.Command(
			"netsh", "interface", "ip", "set", "dns",
			"name="+connName,
			"dhcp",
		)
		g.Log().Infof(ctx, "执行命令: %s", strings.Join(cmd.Args, " "))
		if err := cmd.Run(); err != nil {
			return gerror.Newf("DNS自动获取配置失败: %v", err)
		}
	}

	return nil
}

// getWindowsConnName 通过网卡接口名获取Windows连接名称（辅助函数）
func (s *sSysConfig) getWindowsConnName(nicName string) (string, error) {
	output, err := exec.Command("netsh", "interface", "show", "interface").Output()
	if err != nil {
		return "", gerror.Newf("获取Windows连接列表失败: %v", err)
	}

	// 解析输出，匹配接口名对应的连接名（简化逻辑，实际需更健壮的解析）
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, nicName) {
			parts := strings.Fields(line)
			if len(parts) > 3 {
				return strings.Join(parts[3:], " "), nil
			}
		}
	}

	return "", gerror.Newf("未找到网卡 %s 对应的连接名称", nicName)
}

// cidrToSubnetMask 将CIDR格式转换为标准子网掩码格式
func (s *sSysConfig) cidrToSubnetMask(cidr string) string {
	// 将字符串转换为整数
	maskLen := gconv.Int(cidr)
	if maskLen <= 0 || maskLen > 32 {
		return "255.255.255.255" // 无效的CIDR值
	}

	// 计算子网掩码
	mask := ^((1 << (32 - maskLen)) - 1)
	return fmt.Sprintf("%d.%d.%d.%d",
		(mask>>24)&0xFF,
		(mask>>16)&0xFF,
		(mask>>8)&0xFF,
		mask&0xFF,
	)
}
