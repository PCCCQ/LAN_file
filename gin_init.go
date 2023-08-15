package main

import (
	"changeme/api"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // 允许所有源
		//c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")     // 允许的请求方法
		//c.Header("Access-Control-Allow-Headers", "Origin, Content-Type")    // 允许的请求头
		//c.Header("Access-Control-Expose-Headers", "Content-Length")         // 允许客户端访问的响应头
		//c.Header("Access-Control-Allow-Credentials", "true")               // 允许发送cookie
		//c.Header("Access-Control-Max-Age", "86400")                        // 预检请求的有效期，单位为秒，默认值为24小时

		// 预检请求处理
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 返回空响应
			return
		}

		c.Next()
	}
}
func GinInit() {
	r := gin.Default()
	r.Use(corsMiddleware())
	sub, _ := fs.Sub(dist, "dist")
	r.StaticFS("/index", http.FS(sub))
	r.Static("/uploads", "./uploads")
	//r.GET("/:filename", func(c *gin.Context) {
	//	filename := c.Param("filename")
	//	//c.File("./uploads/" + filename)
	//	c.Writer.Write([]byte("./uploads/" + filename))
	//})
	api.Upload(r)
	api.GetIP(r)
	api.GetFileList(r)
	//r.NoRoute(func(context *gin.Context) {
	//	context.Redirect(http.StatusFound, "/index")
	//})
	r.Run("0.0.0.0:9999")

}
