package handles

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRegister(t *testing.T) {
	r := gin.Default()
	r.POST("/register", Register)
	r.Run(":8080")
}
