package handles

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"myWeb/DataBase/pack"
	"myWeb/cmd/api/common"
	"myWeb/cmd/api/rpc"
	"myWeb/kitex_gen/video"
	"myWeb/pkg/errno"
	"path/filepath"
	"strconv"
	"strings"
)

func UploadVideo(ctx *gin.Context) {
	r := &VideoUploadRequestParam{}
	id, _ := ctx.Get("id")
	tagsStr := ctx.PostForm("tags")
	tagsArr := strings.Split(tagsStr, ",")
	tags := make([]int32, 0, len(tagsArr))
	for _, tagStr := range tagsArr {
		tag, _ := strconv.ParseInt(tagStr, 10, 32)
		tags = append(tags, int32(tag))
	}
	err := ctx.ShouldBind(r)
	if err != nil {
		pack.BuildUploadVideoResp(errno.ErrBind)
		return
	}
	file, err := ctx.FormFile("video")
	if err != nil {
		SendResponse(ctx, pack.BuildUploadVideoResp(err))
		return
	}
	err = ctx.SaveUploadedFile(file, "./video/"+file.Filename)
	abs, err := filepath.Abs("./video/" + file.Filename)
	if err != nil {
		SendResponse(ctx, pack.BuildUpdateUserInfoResp(err))
		return
	}
	oss, err := common.UploadVideoToOSS(file.Filename, abs)
	if err != nil {
		SendResponse(ctx, pack.BuildUploadVideoResp(err))
		return
	}
	v := &video.Video{
		Name:     r.Name,
		Url:      oss,
		Duration: r.Duration,
	}

	resp, err := rpc.UploadVideo(context.Background(), &video.UploadVideoRequest{
		UserId: id.(int64),
		Video:  v,
		Tags:   tags,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildUploadVideoResp(err))
		return
	}
	SendResponse(ctx, resp)
}

func GetVideoInfo(ctx *gin.Context) {
	nums, _ := strconv.ParseInt(ctx.PostForm("nums"), 10, 64)
	id, _ := ctx.Get("id")
	resp, err := rpc.GetVideoInfo(context.Background(), &video.GetVideoInfoRequest{
		UserId: id.(int64),
		Nums:   int32(nums),
	})
	log.Println(resp)
	if err != nil {
		SendResponse(ctx, &video.UploadVideoResponse{
			StatusCode:  10601,
			Description: err.Error(),
		})
		return
	}
	SendResponse(ctx, resp)

}
