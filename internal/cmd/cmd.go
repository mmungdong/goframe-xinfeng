package cmd

import (
	"context"

	"goframe-xinfeng/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmode"

	"goframe-xinfeng/utility/response"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start focus server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			// OpenApi自定义信息
			oai.Info.Title = `API Reference`
			oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware()
				group.Bind(
					controller.Rotation, // 轮播图
				)
				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {

				})
			})
			// 自定义丰富文档
			enhanceOpenAPIDoc(s)
			// 启动Http Server
			s.Run()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `Focus Project`
	openapi.Info.Description = ``

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		// {Name: consts.OpenAPITagNameIndex},
		// {Name: consts.OpenAPITagNameLogin},
		// {Name: consts.OpenAPITagNameRegister},
		// {Name: consts.OpenAPITagNameArticle},
		// {Name: consts.OpenAPITagNameTopic},
		// {Name: consts.OpenAPITagNameAsk},
		// {Name: consts.OpenAPITagNameReply},
		// {Name: consts.OpenAPITagNameContent},
		// {Name: consts.OpenAPITagNameSearch},
		// {Name: consts.OpenAPITagNameInteract},
		// {Name: consts.OpenAPITagNameCategory},
		// {Name: consts.OpenAPITagNameProfile},
		// {Name: consts.OpenAPITagNameUser},
		// {Name: consts.OpenAPITagNameMess},
	}
}
