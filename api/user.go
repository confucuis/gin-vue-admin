// 用户接口

package api

import (
	"gin-vue-admin/api/form"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	LoginForm form.LoginForm
}

func (ua *UserApi) Login(ctx *gin.Context) {
	ctx.Bind(&ua.LoginForm)
	// 数据库查询

	// 返回token
	ctx.JSON(200, gin.H{"message": "success"})
}

func (ua *UserApi) GetInfo(ctx *gin.Context) {

}
