package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/tiger1103/gfast/v3/internal/app/boot"
	_ "github.com/tiger1103/gfast/v3/internal/app/system/packed" //表示GoFrame框架的资源管理，这是一个高级特性。该特性可以将任何资源打包进二进制文件，这样我们在发布的时候，仅需要发布一个二进制文件即可
	"github.com/tiger1103/gfast/v3/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
