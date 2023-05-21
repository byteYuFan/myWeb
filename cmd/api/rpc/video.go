package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"myWeb/kitex_gen/video"
	"myWeb/kitex_gen/video/videoservice"
	"myWeb/pkg/errno"
	"myWeb/pkg/ttviper"
	"net"
	"time"
)

var videoClient videoservice.Client

func UploadVideo(ctx context.Context, req *video.UploadVideoRequest) (resp *video.UploadVideoResponse, err error) {
	resp, err = videoClient.UploadVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}
func GetVideoInfo(ctx context.Context, req *video.GetVideoInfoRequest) (resp *video.GetVideoInfoResponse, err error) {
	resp, err = videoClient.GetVideoInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}

func initVideo() {
	var (
		cfg        = ttviper.ConfigInit("video.yml")
		serverHost = cfg.Viper.GetString("VideoService.Addr")
		serverPort = cfg.Viper.GetString("VideoService.Port")
		serverName = cfg.Viper.GetString("VideoService.ServerName")
	)

	c, err := videoservice.NewClient(serverName,
		client.WithHostPorts(net.JoinHostPort(serverHost, serverPort)),
		client.WithRPCTimeout(30*time.Second),
		client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
			MaxIdlePerAddress: 1000,
			MaxIdleTimeout:    time.Minute}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c

}
func init() {
	initVideo()
}
