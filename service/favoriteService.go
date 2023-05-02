package service

import (
	"TikTokLite/dao"
	. "TikTokLite/message"
	"errors"
)

func FavoriteActionService(userId int64, videoId int64, actionType int32) (*DouyinFavoriteActionResponse, error) {
	var err error
	if actionType == 1 {
		err = dao.InsertFavoriteAction(userId, videoId)
	} else if actionType == 2 {
		err = dao.DeleteFavoriteAction(userId, videoId)
	} else {
		return nil, errors.New("invalid ActionType")
	}
	if err != nil {
		return nil, err
	}
	return &DouyinFavoriteActionResponse{
		Response: SuccessResponse,
	}, nil
}

func FavoriteListService(userId int64) (*DouyinFavoriteListResponse, error) {
	videoList, err := dao.GetFavoriteList(userId)
	if err != nil {
		return nil, err
	}
	return &DouyinFavoriteListResponse{
		Response:  SuccessResponse,
		VideoList: messageVideoList(videoList),
	}, nil
}
