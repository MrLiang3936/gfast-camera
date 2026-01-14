/*
* @desc:摄像头管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:28
 */

package sysCamera

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
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

// 请求 http://192.168.2.125:8080/index/api/getMediaList?secret=qBiDssKtJIk5ngc0M6sa8ZvsCwQAiWcH&vhost=__defaultVhost__&app=proxy&stream=0
// 添加 http://192.168.2.125:8080/index/api/addStreamProxy?secret=qBiDssKtJIk5ngc0M6sa8ZvsCwQAiWcH&vhost=__defaultVhost__&app=proxy&stream=0&url=rtsp://admin:hbAC2023@192.168.2.2:554/h264/ch1/main/av_stream&retry_count=-1&rtp_type=0&timeout_sec=10.0&enable_hls=false&enable_hls_fmp4=false&enable_mp4=false&enable_rtsp=false&enable_rtmp=false&enable_ts=false&enable_fmp4=true&hls_demand=false&rtsp_demand=false&rtmp_demand=false&ts_demand=false&fmp4_demand=true&enable_audio=false&add_mute_audio=true&mp4_save_path=&mp4_max_second=3600&mp4_as_player=false&hls_save_path=&modify_stamp=0&auto_close=false
func (s *sSysCamera) Add(ctx context.Context, req *system.CameraAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查 StreamUrl 是否已经存在
		//count, err := dao.SysCamera.Ctx(ctx).Where(do.SysCamera{StreamUrl: req.StreamUrl}).Count()
		//if err != nil {
		//	liberr.ErrIsNil(ctx, err, "检查StreamUrl是否重复时出错")
		//	return
		//}
		//if count > 0 {
		//	liberr.ErrIsNil(ctx, fmt.Errorf("StreamUrl已存在"), "StreamUrl已存在")
		//	return
		//}

		_, err := dao.SysCamera.Ctx(ctx).Where(do.SysCamera{StreamUrl: req.StreamUrl}).Count()
		if err != nil {
			liberr.ErrIsNil(ctx, err, "检查StreamUrl是否重复时出错")
			return
		}

		// 先插入数据库获取摄像头ID
		var result sql.Result // 声明result变量
		result, err = dao.SysCamera.Ctx(ctx).Insert(do.SysCamera{
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

		// 获取插入的摄像头ID
		cameraId, err := result.LastInsertId()
		if err != nil {
			liberr.ErrIsNil(ctx, err, "获取摄像头ID失败")
			return
		}

		if err := createStreamProxy(ctx, int(cameraId), req.StreamUrl); err != nil {
			g.Log().Warningf(ctx, "创建流媒体代理失败: %v", err)
			liberr.ErrIsNil(ctx, fmt.Errorf("创建流媒体代理失败: %w", err), "创建流媒体代理失败")
			return
		}
	})
	return
}

func (s *sSysCamera) Edit(ctx context.Context, req *system.CameraEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查 StreamUrl 是否已经存在（排除当前摄像头）
		count, err := dao.SysCamera.Ctx(ctx).
			Where(do.SysCamera{StreamUrl: req.StreamUrl}).
			WhereNot("camera_id", req.CameraId).
			Count()
		if err != nil {
			liberr.ErrIsNil(ctx, err, "检查StreamUrl是否重复时出错")
			return
		}
		if count > 0 {
			liberr.ErrIsNil(ctx, fmt.Errorf("StreamUrl已存在"), "StreamUrl已存在")
			return
		}

		// 检查ZLMediaKit中是否已存在流代理
		status, err := checkStreamStatus(ctx, req.CameraId, req.StreamUrl)
		if err != nil {
			// 如果检查失败，先不加
			g.Log().Warningf(ctx, "检查流媒体代理状态失败: %v", err)
			liberr.ErrIsNil(ctx, fmt.Errorf("检查流媒体代理状态失败: %w", err), "检查流媒体代理状态失败")
			return
		}

		switch status {
		case 0: // 流不存在，调用API创建流代理
			if err := createStreamProxy(ctx, req.CameraId, req.StreamUrl); err != nil {
				g.Log().Warningf(ctx, "创建流媒体代理失败: %v", err)
				liberr.ErrIsNil(ctx, fmt.Errorf("创建流媒体代理失败: %w", err), "创建流媒体代理失败")
				return
			}
		case 1: // 流存在且源地址相同，无需操作
			g.Log().Debugf(ctx, "流代理已存在且URL相同，无需操作")
		case 2: // 流存在，但是源地址不一样，先关闭再调用API创建流代理
			if err := closeStreamProxy(ctx, req.CameraId); err != nil {
				g.Log().Warningf(ctx, "关闭流媒体代理失败: %v", err)
				liberr.ErrIsNil(ctx, fmt.Errorf("关闭流媒体代理失败: %w", err), "关闭流媒体代理失败")
				return
			}
			// 成功关闭后，再创建新的流代理
			if err := createStreamProxy(ctx, req.CameraId, req.StreamUrl); err != nil {
				g.Log().Warningf(ctx, "创建流媒体代理失败: %v", err)
				liberr.ErrIsNil(ctx, fmt.Errorf("创建流媒体代理失败: %w", err), "创建流媒体代理失败")
				return
			}
		}

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

// 检查流媒体代理是否存在
func checkStreamStatus(ctx context.Context, cameraId int, streamUrl string) (int, error) {
	config := g.Cfg()
	zlHost := config.MustGet(ctx, "zlmediaKit.host").String()
	zlPort := config.MustGet(ctx, "zlmediaKit.port").Int()
	zlVhost := config.MustGet(ctx, "zlmediaKit.vhost").String()
	zlSecret := config.MustGet(ctx, "zlmediaKit.secret").String()
	zlProxyApp := config.MustGet(ctx, "zlmediaKit.proxyApp").String()

	// "http://192.168.2.125:8080/index/api/getMediaList?secret=qBiDssKtJIk5ngc0M6sa8ZvsCwQAiWcH&vhost=__defaultVhost__&app=proxy&stream=0"
	queryUrl := fmt.Sprintf("http://%s:%d/index/api/getMediaList?secret=%s&vhost=%s&app=%s&stream=%d", zlHost, zlPort, zlSecret, zlVhost, zlProxyApp, cameraId)

	resp, err := http.Get(queryUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	// 检查返回的数据
	if data, ok := result["data"].([]interface{}); ok {
		if len(data) > 0 {
			for _, item := range data {
				if mediaItem, ok := item.(map[string]interface{}); ok {
					if originUrl, exists := mediaItem["originUrl"].(string); exists {
						if originUrl == streamUrl {
							return 1, nil // 找到相同URL的流
						} else {
							return 2, nil // 找到不同URL的流
						}
					}
				}
			}
		}
	}
	return 0, nil // 未找到流
}

// 创建流代理并验证返回结果
func createStreamProxy(ctx context.Context, cameraId int, streamUrl string) error {
	config := g.Cfg()
	zlHost := config.MustGet(ctx, "zlmediaKit.host").String()
	zlPort := config.MustGet(ctx, "zlmediaKit.port").Int()
	zlVhost := config.MustGet(ctx, "zlmediaKit.vhost").String()
	zlSecret := config.MustGet(ctx, "zlmediaKit.secret").String()
	zlProxyApp := config.MustGet(ctx, "zlmediaKit.proxyApp").String()

	// 构建创建流代理的 URL
	createUrl := fmt.Sprintf(
		"http://%s:%d/index/api/addStreamProxy?secret=%s&vhost=%s&app=%s&stream=%d&url=%s&retry_count=-1&rtp_type=0&timeout_sec=10.0&enable_hls=false&enable_hls_fmp4=false&enable_mp4=false&enable_rtsp=true&enable_rtmp=false&enable_ts=false&enable_fmp4=true&hls_demand=false&rtsp_demand=false&rtmp_demand=false&ts_demand=false&fmp4_demand=true&enable_audio=false&add_mute_audio=true&mp4_save_path=&mp4_max_second=3600&mp4_as_player=false&hls_save_path=&modify_stamp=0&auto_close=false",
		zlHost, zlPort, zlSecret, zlVhost, zlProxyApp, cameraId, url.QueryEscape(streamUrl),
	)

	resp, err := http.Get(createUrl)
	if err != nil {
		return fmt.Errorf("创建流代理请求失败: %w", err)
	}
	defer resp.Body.Close()

	var result model.ZLMediaKitResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("解析启动API响应失败: %w", err)
	}

	// 检查返回的 code 是否为 0（表示成功）
	if result.Code != 0 {
		return fmt.Errorf("创建流代理失败，错误代码: %d, 错误信息: %s", result.Code, result.Msg)
	}

	g.Log().Debugf(ctx, "成功创建流代理，返回键值: %v", result.Data)
	return nil
}

func closeStreamProxy(ctx context.Context, cameraId int) error {
	config := g.Cfg()
	zlHost := config.MustGet(ctx, "zlmediaKit.host").String()
	zlPort := config.MustGet(ctx, "zlmediaKit.port").Int()
	zlVhost := config.MustGet(ctx, "zlmediaKit.vhost").String()
	zlSecret := config.MustGet(ctx, "zlmediaKit.secret").String()
	zlProxyApp := config.MustGet(ctx, "zlmediaKit.proxyApp").String()

	// 先关闭已有的 http://192.168.2.125:8080/index/api/close_streams?secret=qBiDssKtJIk5ngc0M6sa8ZvsCwQAiWcH&vhost=__defaultVhost__&app=proxy&stream=10&force=1
	closeUrl := fmt.Sprintf(
		"http://%s:%d/index/api/close_streams?secret=%s&vhost=%s&app=%s&stream=%d&force=1",
		zlHost, zlPort, zlSecret, zlVhost, zlProxyApp, cameraId,
	)

	closeResp, closeErr := http.Get(closeUrl)
	if closeErr != nil {
		return fmt.Errorf("关闭旧的流代理请求失败: %w", closeErr)
	}
	defer closeResp.Body.Close()
	var closeResult model.ZLMediaKitResponse
	if err := json.NewDecoder(closeResp.Body).Decode(&closeResult); err != nil {
		return fmt.Errorf("解析关闭API响应失败: %w", err)
	}
	// 检查返回的 code 是否为 0（表示成功）
	if closeResult.Code != 0 {
		return fmt.Errorf("关闭流代理失败，错误代码: %d, 错误信息: %s", closeResult.Code, closeResult.Msg)
	}

	// 检查是否成功关闭了流
	if closeResult.CountClosed == 0 {
		g.Log().Debugf(ctx, "摄像头ID %d 的流代理可能不存在或已关闭 (count_closed: %d, count_hit: %d)",
			cameraId, closeResult.CountClosed, closeResult.CountHit)
	} else {
		g.Log().Debugf(ctx, "成功关闭摄像头ID %d 的流代理 (count_closed: %d, count_hit: %d)",
			cameraId, closeResult.CountClosed, closeResult.CountHit)
	}

	g.Log().Debugf(ctx, "关闭创建流代理，返回键值: %v", closeResult.Data)
	return nil
}

func (s *sSysCamera) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 先关闭对应的流代理
		for _, id := range ids {
			if closeErr := closeStreamProxy(ctx, id); closeErr != nil {
				g.Log().Warningf(ctx, "关闭流媒体代理失败，摄像头ID: %d, 错误: %v", id, closeErr)
				liberr.ErrIsNil(ctx, closeErr, "关闭流媒体代理失败")
				return
			}
		}

		// 删除数据库记录
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
			zlProxyApp := config.MustGet(ctx, "zlmediaKit.proxyApp").String()

			// 如果配置不存在，使用默认值
			if zlHost == "" {
				zlHost = "127.0.0.1"
			}
			if zlPort == 0 {
				zlPort = 8080
			}
			res.WsFmp4Url = fmt.Sprintf("ws://%s:%d/%s/%d.live.mp4", zlHost, zlPort, zlProxyApp, res.CameraId)
			res.RtspUrl = fmt.Sprintf("rtsp://%s:8554/%s/%d", zlHost, zlProxyApp, res.CameraId)
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
