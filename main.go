package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "goframe-xinfeng/internal/packed"

	_ "goframe-xinfeng/internal/logic"

	"goframe-xinfeng/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
