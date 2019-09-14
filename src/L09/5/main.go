package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// multipart 表单内存限制 (默认 32 MB)
	// router.MaxMultipartMemory = 8 << 20  // 设置为8 MB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件上传，file为控件名
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		// 上传后的文件名
		dst := "test.png"
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' 上传成功！", file.Filename))
	})
	router.Run(":8080")
}
