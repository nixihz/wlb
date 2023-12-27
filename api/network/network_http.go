package common

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)

// RegisterUploadHTTPServer 使用gin的文件上传服务
func RegisterUploadHTTPServer(s *http.Server) {
	router := s.Route("/")

	router.POST("/common/file/upload", func(ctx http.Context) error {
		return ctx.Result(200, out)
	})
}
