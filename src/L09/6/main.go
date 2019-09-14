package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// multipart 表单内存限制 (默认 32 MB)
	// router.MaxMultipartMemory = 8 << 20  // 设置为8 MB

	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		// 多文件上传，files为控件名
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("表单异常： %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			// 上传后的文件名
			if err := c.SaveUploadedFile(file, "up-"+filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("文件上传错误： %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf(" %d 个文件已上传！", len(files)))

	})
	router.Run(":8080")
}
