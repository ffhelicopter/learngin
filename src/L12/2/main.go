package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./templates/index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{
			"now": time.Date(2019, 07, 11, 16, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}
