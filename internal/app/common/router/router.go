/*
* @desc:后台路由
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/2/18 17:34
 */

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/common/controller"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/pub", func(group *ghttp.RouterGroup) {
		group.Group("/captcha", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.Captcha,
			)
		})
		group.Group("/file", func(group *ghttp.RouterGroup) {
			//登录验证拦截
			service.GfToken().Middleware(group)
			//context拦截器
			group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
			//后台操作日志记录
			group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
			group.Bind(
				controller.File,
			)
			//自动绑定定义的控制器
			if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
				panic(err)
			}
		})
	})
}
