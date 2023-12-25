package controllers

import (
	"gin-template/middlewares"
	tr "gin-template/utils/trace"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type RequestResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Name        string
	Version     string
	BasePath    string
	TokenUserId int
	TokenAppId  int
)

func InitRouters(addr string) {
	r := setupRouter()
	err := r.Run(addr)
	if err != nil {
		return
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	v1Path := "apis/"

	//配置中间件
	r.Use(gin.Recovery())
	r.Use(func(gctx *gin.Context) {
		//识别 http 头部中的 trace gctx 信息
		ctx := otel.GetTextMapPropagator().Extract(gctx.Request.Context(), propagation.HeaderCarrier(gctx.Request.Header))

		ctx, span := tr.Tracer.Start(ctx, gctx.Request.RequestURI, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		gctx.Request = gctx.Request.WithContext(ctx)

		gctx.Next()
	})
	r.Use(middlewares.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	t := r.Group(v1Path + "market")
	{
		t.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "test",
			})
		})
	}

	return r
}
