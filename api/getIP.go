package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func GetIP(engine *gin.Engine) {
	engine.GET("/getIP", func(c *gin.Context) {
		// 返回文件的URL地址
		c.JSON(http.StatusOK, gin.H{
			"IPList": GetIPFunc(),
		})
	})

}

func GetIPFunc() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	IPList := make([]string, len(addrs))
	for _, address := range addrs {
		// 检查IP地址，其他类型的地址(如link-local或者loopback)忽略
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil && ipnet.IP.String() != "" {
			fmt.Println(ipnet.IP.String())
			IPList = append(IPList, ipnet.IP.String())
		}
	}
	return IPList
}
