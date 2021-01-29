package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("imgfile")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filepath := "static/"+time.Now().Format("20060102")+"/" // viper.GetString(`app.upload_file_path`)
	fmt.Println(filepath)
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}

	filename := filepath + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传成功!",
		"result": gin.H{
			"path": filename,
		},
	})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("realfile")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filepath := "static/"+time.Now().Format("20060102")+"/" // viper.GetString(`app.upload_file_path`)
	fmt.Println(filepath)
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}

	filename := filepath + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传成功!",
		"result": gin.H{
			"path": filename,
		},
	})
}