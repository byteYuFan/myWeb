package ttviper

import (
	"testing"
)

func TestConfigInit(t *testing.T) {
	cfg := ConfigInit("Config.yml")
	cfg.Viper.ReadInConfig()
	t.Log(cfg.Viper.GetString("Server.Name"))
	t.Log(cfg.Viper.GetString("MySQL.Address"))

}
