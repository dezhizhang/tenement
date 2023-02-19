package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/utils"
)

func GetSession(c *gin.Context) {
	rsp := make(map[string]string)

	rsp["errno"] = utils.RECODE_SESSIONERR
	rsp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	c.JSON(http.StatusOK, rsp)

	//ctx.JSON(http.StatusOK, gin.H{"errno": "0", "errormsg": "ok", "data": nil})
}
