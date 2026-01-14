/*
* @desc:请求zlMediaKit参数结构体
* @company:
* @Author: liangyu
* @Date:   2026/1/13 11:47
 */

package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// -------------------------- 请求zlMediaKit参数结构体 --------------------------
// CameraAddMediaReq 添加相机+创建zlmediaKit拉流代理的请求参数
type CameraAddMediaReq struct {
	Secret string `p:"secret" v:"required#ZLMediaKit API密钥不能为空"`
	Vhost  string `p:"vhost" v:"required#虚拟主机名不能为空,例如__defaultVhost__"`
	App    string `p:"app" v:"required#应用名不能为空,例如live"`
	Stream string `p:"stream" v:"required#流ID不能为空,例如camera_001"`
	Url    string `p:"url" v:"required#RTSP拉流地址不能为空" dc:"相机RTSP地址，如rtsp://admin:123456@192.168.1.100:554/stream1"`
	// ZLMediaKit 可选参数（可根据业务扩展）
	RetryCount    int     `p:"retryCount" v:"min:-1#拉流重试次数不能小于-1" dc:"拉流重试次数，默认为-1（无限重试）" default:"-1"`
	RtpType       int     `p:"rtpType" v:"in:0,1,2#RTSP拉流方式只能是0(tcp)/1(udp)/2(组播)" dc:"RTSP拉流方式" default:"0"`
	TimeoutSec    float64 `p:"timeoutSec" dc:"拉流超时时间(秒)" default:"10.0"`
	EnableHls     bool    `p:"enableHls" dc:"是否转换成hls-mpegts协议" default:"false"`
	EnableHlsFmp4 bool    `p:"enableHlsFmp4" dc:"是否转换成hls-fmp4协议" default:"false"`
	EnableMp4     bool    `p:"enableMp4" dc:"是否允许mp4录制" default:"false"`
	EnableRtsp    bool    `p:"enableRtsp" dc:"是否转rtsp协议" default:"false"`
	EnableRtmp    bool    `p:"enableRtmp" dc:"是否转rtmp/flv协议" default:"false"`
	EnableTs      bool    `p:"enableTs" dc:"是否转http-ts/ws-ts协议" default:"false"`
	EnableFmp4    bool    `p:"enableFmp4" dc:"是否转http-fmp4/ws-fmp4协议" default:"true"`
	HlsDemand     bool    `p:"hlsDemand" dc:"该协议是否有人观看才生成" default:"false"`
	RtspDemand    bool    `p:"rtspDemand" dc:"该协议是否有人观看才生成" default:"false"`
	RtmpDemand    bool    `p:"rtmpDemand" dc:"该协议是否有人观看才生成" default:"false"`
	TsDemand      bool    `p:"tsDemand" dc:"该协议是否有人观看才生成" default:"false"`
	Fmp4Demand    bool    `p:"fmp4Demand" dc:"该协议是否有人观看才生成" default:"false"`
	EnableAudio   bool    `p:"enableAudio" dc:"转协议时是否开启音频" default:"false"`
	AddMuteAudio  bool    `p:"addMuteAudio" dc:"转协议时，无音频是否添加静音aac音频" default:"false"`
	Mp4SavePath   string  `p:"mp4SavePath" dc:"mp4录制文件保存根目录，置空使用默认"`
	Mp4MaxSecond  int     `p:"mp4MaxSecond" dc:"mp4录制切片大小，单位秒" default:"3600"`
	Mp4AsPlayer   bool    `p:"mp4AsPlayer" dc:"MP4录制是否当作观看者参与播放人数计数" default:"false"`
	HlsSavePath   string  `p:"hlsSavePath" dc:"hls文件保存保存根目录，置空使用默认"`
	ModifyStamp   int     `p:"modifyStamp" v:"in:0,1,2#时间戳覆盖类型只能是0(绝对时间戳)/1(系统时间戳)/2(相对时间戳)" dc:"该流是否开启时间戳覆盖(0:绝对时间戳/1:系统时间戳/2:相对时间戳)" default:"0"`
	AutoClose     bool    `p:"autoClose" dc:"无人观看是否自动关闭流(不触发无人观看hook)" default:"false"`
}

// CameraAddMediaRes 接口返回结构体
type CameraAddMediaRes struct {
	g.Meta `mime:"application/json"`
	Code   int         `json:"code" dc:"状态码 0成功/非0失败"`
	Msg    string      `json:"msg" dc:"提示信息"`
	Data   interface{} `json:"data" dc:"返回数据"`
}

type CameraGetMediaRes struct {
	*entity.SysCamera
	WsFmp4Url string `json:"wsFmp4Url"`
	RtspUrl   string `json:"rtspUrl"`
}

type ZLMediaKitResponse struct {
	Code        int         `json:"code"`
	Data        interface{} `json:"data,omitempty"`
	Msg         string      `json:"msg,omitempty"`          // 可能存在错误信息
	CountClosed int         `json:"count_closed,omitempty"` // 关闭流时候返回的关闭数量
	CountHit    int         `json:"count_hit,omitempty"`    // 关闭流时候返回的关闭数量
}

type AddStreamProxyResponse struct {
	Code int `json:"code"`
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}
