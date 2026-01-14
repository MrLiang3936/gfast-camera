/*
* @desc:验证码获取
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/3/2 17:45
 */

package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/common"
)

var File = fileController{}

type fileController struct {
}

// 文件上传
func (a *fileController) Get(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	//req.File.Filename = "MyCustomFileName.txt"
	upUrl := g.Cfg().MustGet(ctx, "upload.url").String()
	names, err := req.File.Save(upUrl)

	if err != nil {
		return nil, err
	}

	res = &common.FileUploadRes{
		Name: names,
		Url:  upUrl,
	}
	return
}
