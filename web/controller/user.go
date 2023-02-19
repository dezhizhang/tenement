package controller

import (
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
	"net/http"
	"web/utils"
)

func GetSession(c *gin.Context) {
	rsp := make(map[string]string)

	rsp["code"] = utils.RECODE_SESSIONERR
	rsp["msg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	c.JSON(http.StatusOK, rsp)
}

//获取图片信息

func GetImageCode(c *gin.Context) {
	uuid := c.Param("uuid")

	cap := captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加
	cap.SetFont("web/conf/comic.ttf")
	// 设置验证码大小
	cap.SetSize(128, 64)
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	img, str := cap.Create(6, captcha.NUM)
	png.Encode(c.Writer, img)

	fmt.Println(uuid, str)
	//fmt.Println("uuid" + uuid)
}
