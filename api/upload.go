package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func Upload(engine *gin.Engine) {
	engine.POST("/uploadFile", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传文件失败：%s", err.Error()))
			return
		}

		// 生成时间作为文件名
		timeName := strconv.FormatInt(time.Now().Unix(), 10)

		//文件名后缀
		fileExt := filepath.Ext(file.Filename)
		//文件名不含后缀
		filename := file.Filename[:len(file.Filename)-len(fileExt)]
		newFileName := timeName + "_" + filename + fileExt

		// 将文件保存在本地文件夹中
		filePath := "uploads/" + newFileName
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("保存文件失败：%s", err.Error()))
			return
		}

		// 返回文件的URL地址
		c.JSON(http.StatusOK, gin.H{
			"url": "/api/uploads/" + newFileName,
		})
	})

}
