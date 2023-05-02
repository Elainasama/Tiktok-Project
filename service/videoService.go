package service

import (
	"TikTokLite/dao"
	. "TikTokLite/message"
	"TikTokLite/model"
)

func GetVideoListResponse(userId, currentTime int64) (*FeedResponse, error) {
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
	mes := messageVideoList(videolist)
	if userId != -1 {
		for i := range mes {
			mes[i].IsFavorite = dao.IsFavorite(userId, mes[i].Id)
		}
	}
	feed := &FeedResponse{
		Response:  SuccessResponse,
		VideoList: mes,
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
