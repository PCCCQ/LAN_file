package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type FileInfo struct {
	Name string
	Size string
	Date string
}

func GetFileList(engine *gin.Engine) {
	engine.GET("/getFileList", func(c *gin.Context) {
		// 指定文件夹路径
		folderPath := "./uploads"

		// 获取文件夹中的文件列表
		fileList, err := FileList(folderPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 打印文件列表
		fmt.Println("Files in", folderPath, ":")
		for _, file := range fileList {
			fmt.Println(file)
		}
		c.JSON(http.StatusOK, gin.H{
			"FileList": fileList,
		})
	})

}

// FileList 获取指定文件夹中的文件列表
// 参数 folderPath: 文件夹路径
// 返回文件列表的字符串切片和可能的错误
func FileList(folderPath string) ([]FileInfo, error) {
	var fileList []FileInfo // 存储文件列表的字符串切片

	// 打开文件夹
	folder, err := os.Open(folderPath)
	if err != nil {
		return fileList, err // 如果打开文件夹失败，返回可能的错误
	}
	defer folder.Close() // 延迟关闭文件夹，确保函数返回前关闭

	// 读取文件列表
	files, err := folder.Readdir(-1)
	if err != nil {
		return fileList, err // 如果读取文件列表失败，返回可能的错误
	}

	// 遍历文件列表
	for _, file := range files {
		// 排除子文件夹，只获取普通文件
		if !file.IsDir() {

			fileList = append(fileList, FileInfo{
				Name: file.Name(),
				Size: fmt.Sprintf("%.3f", float32(file.Size())/1000/1000),
				//Size: float32(file.Size()),
				Date: file.ModTime().Format("2006-01-02 15:04:05"),
			}) // 将文件名添加到文件列表中
		}
	}
	//按照倒序排序
	n := len(fileList)
	tmp := make([]FileInfo, n)
	for i, v := range fileList {
		tmp[n-i-1] = v
	}
	return tmp, nil // 返回文件列表和nil，表示没有错误发生
}
