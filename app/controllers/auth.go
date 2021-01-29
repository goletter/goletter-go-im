package controllers

import (
	"github.com/gin-gonic/gin"
)

type VisitorOnline struct {
	Uid         string `json:"uid"`
	Username    string `json:"username"`
	Avator      string `json:"avator"`
	LastMessage string `json:"last_message"`
}

func Login(c *gin.Context) {
	//
}