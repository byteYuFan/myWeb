package ttviper

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func ConfigInit(cfgFileName string) {
	viper.SetConfigName("userConfig")    // 指定配置文件名，不需要指定扩展名
	viper.AddConfigPath("../../config/") // 指定配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	// 读取 Server 配置项
	serverName := viper.GetString("Server.Name")
	serverAddr := viper.GetString("Server.Address")
	serverPort := viper.GetInt("Server.Port")
	serverMemory := viper.GetInt("Server.Argon2ID.Memory")
	serverIterations := viper.GetInt("Server.Argon2ID.Iterations")
	serverParallelism := viper.GetInt("Server.Argon2ID.Parallelism")
	serverSaltLength := viper.GetInt("Server.Argon2ID.SaltLength")
	serverKeyLength := viper.GetInt("Server.Argon2ID.keyLength")

	fmt.Printf("Server.Name: %s\n", serverName)
	fmt.Printf("Server.Address: %s\n", serverAddr)
	fmt.Printf("Server.Port: %d\n", serverPort)
	fmt.Printf("Server.Argon2ID.Memory: %d\n", serverMemory)
	fmt.Printf("Server.Argon2ID.Iterations: %d\n", serverIterations)
	fmt.Printf("Server.Argon2ID.Parallelism: %d\n", serverParallelism)
	fmt.Printf("Server.Argon2ID.SaltLength: %d\n", serverSaltLength)
	fmt.Printf("Server.Argon2ID.keyLength: %d\n", serverKeyLength)
}
