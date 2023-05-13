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

// UserRegisterResponseParam 用户信息 输出参数
type UserRegisterResponseParam struct {
	StatusCode  int32  `json:"status_code,omitempty"` // 状态码
	Description string `json:"description,omitempty"` // 描述信息
}
