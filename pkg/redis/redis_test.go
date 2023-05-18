package redisMiddleware

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	Insert()
}
func TestStoreCode(t *testing.T) {
	err := StoreCode("email", "1523768711@qq.com", "128673", 128*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestGetEmailCode(t *testing.T) {
	code, err := GetEmailCode("email", "1523768711@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(code)
}
