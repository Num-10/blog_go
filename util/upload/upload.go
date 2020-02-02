package upload

import (
	"blog_go/conf"
	"blog_go/util/e"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func UploadSetUp()  {
	_, err := os.Stat(conf.UploadIni.SaveImagePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(conf.UploadIni.SaveImagePath, 0666)
			if err != nil {
				fmt.Println("create image path fail: " + err.Error())
				os.Exit(e.SERVICE_ERROR)
			}
		}
	}
}

func CheckImage(file *multipart.FileHeader) (bool, int) {
	if file.Size > int64(conf.UploadIni.ImageMaxSize << 20) {
		return false, e. IMAGE_OVER_SIZE
	}
	isAllow := false
	ext := strings.ToLower(filepath.Ext(file.Filename))
	for _, allow := range conf.UploadIni.ImageAllowExtSlice {
		if ext == strings.ToLower(allow) {
			isAllow = true
		}
	}
	if !isAllow {
		return false, e.IMAGE_NOT_ALLOW_EXT
	}

	return true, e.SERVICE_SUCCESS
}

func SaveImage(file *multipart.FileHeader, c *gin.Context) (bool, string) {
	str := []byte(file.Filename + strconv.Itoa(int(time.Now().Unix())))

	date := time.Now().Format("20060102")
	imageSaveUrl := path.Join(date, fmt.Sprintf("%x", md5.Sum(str)) + strings.ToLower(filepath.Ext(file.Filename)))
	imageSavePath := path.Join(conf.AppIni.RootPath, conf.UploadIni.SaveImagePath, imageSaveUrl)

	dir := path.Join(conf.AppIni.RootPath, conf.UploadIni.SaveImagePath, date)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Println("create upload/image/ fail：" + err.Error())
			}
		}
	}

	err = c.SaveUploadedFile(file, imageSavePath)
	if err != nil {
		return false, ""
	}
	return true, imageSaveUrl
}
