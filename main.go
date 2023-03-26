package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"netcut/config"
	"netcut/controller"
	"netcut/model"
	"os"

	"github.com/gin-gonic/gin"
)

// go:embed view
var viewFs embed.FS

func main() {
	config.Init()
	model.Init()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	var cut controller.CutCtr
	r := gin.Default()
	t := r.Group("/api")
	t.Use(cors())
	{
		t.GET("", cut.Get)
		t.POST("", cut.Creat)
		t.GET("/text", cut.CreatText)
		// t.POST("/text", cut.CreatText)
	}

	fss, err := fs.Sub(viewFs, "view")
	if err != nil {
		log.Println(err)
	}
	r.StaticFS("/view", http.FS(fss))

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/view")
	})

	r.Run(config.Port) // 监听并在 0.0.0.0:8080 上启动服务
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// if origin != "" {
		// 可将将* 替换为指定的域名
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// }
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
