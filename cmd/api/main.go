package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"myWeb/cmd/api/auth"
	"myWeb/cmd/api/handles"
	"myWeb/pkg/ttviper"
	"net/http"
	"strings"
)

var (
	cfg           = ttviper.ConfigInit("config.yml")
	ServerAddress = cfg.Viper.GetString("Server.Address")
	ServerPort    = cfg.Viper.GetString("Server.Port")
	certFile      = cfg.Viper.GetString("Server.CertFile")
	keyFile       = cfg.Viper.GetString("Server.KeyFile")
)

func main() {
	router := gin.Default()
	router.Use(Cors())
	user := router.Group("/xaut")
	user.POST("/user/register/", handles.Register)
	user.POST("/user/login/", handles.Login)
	user.POST("/user/login-email/", handles.LoginByEmail)
	user.POST("/user/send-email/", handles.SendEmail)
	user.POST("/user/login/forget-password/", handles.RestPassword)
	user.POST("/user/modify-info/", auth.TokenAuthMiddleware(), handles.UpdateUserInfo)
	user.POST("/user/modify-password/", auth.TokenAuthMiddleware(), handles.ChangeUserPassword)
	video := router.Group("/video")
	video.POST("/upload-video/", auth.TokenAuthMiddleware(), handles.UploadVideo)
	video.GET("/get-video-info/", auth.TokenAuthMiddleware(), handles.GetVideoInfo)
	// 设置证书和私钥文件的路径

	// 加载证书和私钥文件
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    ServerAddress + ":" + ServerPort,
		Handler: router,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
	log.Println("service start successfully.")
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range context.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		context.Next()
	}
}
func init() {
}
