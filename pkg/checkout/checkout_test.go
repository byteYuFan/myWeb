package checkout

import "testing"

func TestValidateUsername(t *testing.T) {
	res := ValidateUsername("wy151618")
	t.Log(res)
}

func TestValidatePassword(t *testing.T) {
	res := ValidatePassword("4Kn2cWq2fwD4fb")
	t.Log(res)
}

func TestValidateEmail(t *testing.T) {
	res := ValidateEmail("854978151")
	t.Log(res)
}
