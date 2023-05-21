package common

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"myWeb/pkg/ttviper"
	"os"
)

func UploadVideoToOSS(fileName, path string) (string, error) {
	cfg := ttviper.ConfigInit("video.yml")
	endPoint := cfg.Viper.GetString("VideoService.AliyunOss.Endpoint")
	accessKeyID := cfg.Viper.GetString("VideoService.AliyunOss.AccessKeyID")
	accessKeySecret := cfg.Viper.GetString("VideoService.AliyunOss.AccessKeySecret")
	client, err := oss.New(endPoint, accessKeyID, accessKeySecret)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket("pogf")
	if err != nil {
		return "", err
	}
	objectKey := "video/" + fileName
	err = bucket.PutObjectFromFile(objectKey, path)
	if err != nil {
		return "", err
	}
	err = os.Remove(path)
	if err != nil {
		return "", err
	}
	url := "https://pogf.oss-cn-beijing.aliyuncs.com/" + objectKey
	return url, nil
}
