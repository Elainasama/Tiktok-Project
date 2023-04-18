package service

import (
	"TikTokLite/dao"
	. "TikTokLite/message"
	"TikTokLite/model"
)

func GetVideoListResponse(currentTime int64) (*FeedResponse, error) {
	videolist, err := dao.GetVideoList(currentTime)
	if err != nil {
		return nil, err
	}
	var nextTime int64
	if len(videolist) == 0 {
		nextTime = 0
	} else {
		nextTime = videolist[len(videolist)-1].PublishTime
	}
	feed := &FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: messageVideoList(videolist),
		NextTime:  nextTime,
	}
	return feed, nil
}

func messageVideoList(v []model.Video) []Video {
	var videoList []Video
	for _, video := range v {
		videoList = append(videoList, Video{
			Id:            video.VideoId,
			Author:        *messageUserInfo(video.Author),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			PublishTime:   video.PublishTime,
		})
	}
	return videoList
}
