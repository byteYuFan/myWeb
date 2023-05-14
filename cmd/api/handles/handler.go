package handles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendResponse pack response
func SendResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

// UserRegisterRequestParam  handler 输入参数
type UserRegisterRequestParam struct {
	UserName string `json:"username"` // 用户名
	Email    string `json:"email"`    //用户邮箱
	Password string `json:"password"` // 用户密码
	Confirm  string `json:"confirm"`  //确认密码
}

type UserLoginRequestParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
