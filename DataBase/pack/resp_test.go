package pack

import (
	"myWeb/pkg/errno"
	"testing"
)

func TestBuildUserRegisterResponse(t *testing.T) {
	resp := BuildUserRegisterResponse(errno.ErrBind)
	t.Log(resp)
}
