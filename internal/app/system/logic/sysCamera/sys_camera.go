/*
* @desc:摄像头管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysCamera

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysCamera(New())
}

func New() *sSysCamera {
	return &sSysCamera{}
}

type sSysCamera struct {
}

// List 摄像头列表
func (s *sSysCamera) List(ctx context.Context, req *system.CameraSearchReq) (res *system.CameraSearchRes, err error) {
	res = new(system.CameraSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysCamera.Ctx(ctx)
		if req != nil {
			if req.CameraName != "" {
				m = m.Where("camera_name like ?", "%"+req.CameraName+"%")
			}
			//if req.Status != "" {
			//	m = m.Where("status", gconv.Uint(req.Status))
			//}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取摄像头失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		res.CurrentPage = req.PageNum
		err = m.Page(req.PageNum, req.PageSize).Order("created_at desc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取摄像头失败")
	})
	return
}

// 添加 http://192.168.2.125:8080/index/api/addStreamProxy?secret=qBiDssKtJIk5ngc0M6sa8ZvsCwQAiWcH&vhost=__defaultVhost__&app=proxy&stream=0&url=rtsp://admin:hbAC2023@192.168.2.2:554/h264/ch1/main/av_stream&retry_count=-1&rtp_type=0&timeout_sec=10.0&enable_hls=false&enable_hls_fmp4=false&enable_mp4=false&enable_rtsp=false&enable_rtmp=false&enable_ts=false&enable_fmp4=true&hls_demand=false&rtsp_demand=false&rtmp_demand=false&ts_demand=false&fmp4_demand=true&enable_audio=false&add_mute_audio=true&mp4_save_path=&mp4_max_second=3600&mp4_as_player=false&hls_save_path=&modify_stamp=0&auto_close=false
func (s *sSysCamera) Add(ctx context.Context, req *system.CameraAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysCamera.Ctx(ctx).Insert(do.SysCamera{
			GroupId:    req.GroupId,
			CameraName: req.CameraName,
			GpsInfo:    req.GpsInfo,
			DeviceType: req.DeviceType,
			StreamUrl:  req.StreamUrl,
			PreviewImg: req.PreviewImg,
			Remark:     req.Remark,
			CreateBy:   service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加摄像头失败")
	})
	return
}

func (s *sSysCamera) Edit(ctx context.Context, req *system.CameraEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysCamera.Ctx(ctx).WherePri(req.CameraId).Update(do.SysCamera{
			CameraId:   req.CameraId,
			GroupId:    req.GroupId,
			CameraName: req.CameraName,
			GpsInfo:    req.GpsInfo,
			DeviceType: req.DeviceType,
			StreamUrl:  req.StreamUrl,
			PreviewImg: req.PreviewImg,
			Remark:     req.Remark,
			UpdateBy:   service.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改摄像头失败")
	})
	return
}

func (s *sSysCamera) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysCamera.Ctx(ctx).Where(dao.SysCamera.Columns().CameraId+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 获取 ws://192.168.2.125:8080/${app}/${stream}.live.mp4
func (s *sSysCamera) Get(ctx context.Context, id int) (res *model.CameraGetMediaRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysCamera.Ctx(ctx).Where(dao.SysCamera.Columns().CameraId, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取摄像头数据失败")

		// 为 wsFmp4Url 额外赋值
		if res != nil && res.StreamUrl != "" {
			/*// 从环境变量获取服务器IP，如果没有则使用默认值
			serverIP := os.Getenv("SERVER_IP")
			if serverIP == "" {
				// 获取当前服务器IP地址
				currentIP, ipErr := GetInternalIP()
				if ipErr != nil {
					currentIP = "127.0.0.1"
				}
				serverIP = currentIP
			}
			res.WsFmp4Url = fmt.Sprintf("ws://%s:8080/proxy/%d.live.mp4", serverIP, res.CameraId)*/
			// 从配置文件获取 ZLMediaKit 服务地址
			config := g.Cfg()
			zlHost := config.MustGet(ctx, "zlmediaKit.host").String()
			zlPort := config.MustGet(ctx, "zlmediaKit.port").Int()

			// 如果配置不存在，使用默认值
			if zlHost == "" {
				zlHost = "127.0.0.1"
			}
			if zlPort == 0 {
				zlPort = 8080
			}
			res.WsFmp4Url = fmt.Sprintf("ws://%s:%d/proxy/%d.live.mp4", zlHost, zlPort, res.CameraId)
		}
	})
	return
}

// GetInternalIP 获取本机有效内网IP（优先IPv4）
func GetInternalIP() (string, error) {
	// 遍历所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("遍历网卡失败：%v", err)
	}

	for _, iface := range interfaces {
		// 跳过禁用、回环、虚拟网卡
		if iface.Flags&net.FlagUp == 0 || // 网卡未启用
			iface.Flags&net.FlagLoopback != 0 || // 回环地址（127.0.0.1）
			strings.Contains(iface.Name, "docker") || // 过滤docker虚拟网卡
			strings.Contains(iface.Name, "vmnet") { // 过滤虚拟机网卡
			continue
		}

		// 获取该网卡的所有IP地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue // 跳过获取失败的网卡
		}

		// 遍历IP，优先返回IPv4地址
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}

			// 过滤IPv6，只返回IPv4
			if ip4 := ipNet.IP.To4(); ip4 != nil {
				return ip4.String(), nil
			}
		}
	}

	return "", fmt.Errorf("未找到有效内网IP")
}
