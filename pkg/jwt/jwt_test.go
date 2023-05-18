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
	c, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTgsInVzZXJuYW1lIjoid3d5eWZmMTIzIiwiZXhwIjoxNjg0NDEyMTk5LCJpYXQiOjE2ODQzMjU3OTksImlzcyI6InBvZ2YifQ.fFdeGn5ltyW2R18SAAUSfRRn_aOHkm-nQpVPLLBRBao")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
