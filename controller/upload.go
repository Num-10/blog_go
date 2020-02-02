package controller

import (
	"blog_go/conf"
	"blog_go/util/e"
	"blog_go/util/upload"
	"github.com/gin-gonic/gin"
)

func UploadImageLocal(c *gin.Context)  {
	file, _ := c.FormFile("image")
	var url string

	if file != nil {
		if ok, code := upload.CheckImage(file); !ok {
			e.Json(c, &e.Return{Code:code})
			return
		}

		if ok, image_url := upload.SaveImage(file, c); !ok {
			e.Json(c, &e.Return{Code:e.IMAGE_SAVE_FIAL})
			return
		} else {
			url = conf.AppIni.DomainUrl + conf.AppIni.ImageUrl + image_url
		}
	}

	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data: map[string]interface{}{"url": url}})
}