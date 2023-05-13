package jwt

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(1, "wyf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(token)
}

func TestParseToken(t *testing.T) {
	c, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ3eWYiLCJleHAiOjE2ODM4NzU4NjcsImlhdCI6MTY4Mzc4OTQ2NywiaXNzIjoicG9nZiJ9.rXzf1fzidsfwe4HRBt7JN_NxAxUceD0HxdpCQsbPuFc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
