package service

import (
	"TikTokLite/dao"
	. "TikTokLite/message"
	"TikTokLite/model"
	"errors"
)

func CommentActionService(userId int64, videoId int64, actionType int32, text string, commentId int64) (*DouyinCommentActionResponse, error) {
	if actionType == 1 {
		comment, err := dao.InsertComment(userId, videoId, text)
		if err != nil {
			return nil, err
		}
		return &DouyinCommentActionResponse{
			Response: Response{
				StatusCode: 0,
			},
			Comment: messageComment(comment),
		}, nil
	} else if actionType == 2 {
		comment, err := dao.DeleteComment(videoId, commentId)
		if err != nil {
			return nil, err
		}
		return &DouyinCommentActionResponse{
			Response: Response{
				StatusCode: 0,
			},
			Comment: messageComment(comment),
		}, nil
	} else {
		return nil, errors.New("invalid actionType")
	}
}

func GetCommentListService(videoId int64) (*DouyinCommentListResponse, error) {
	commentList, err := dao.GetCommentList(videoId)
	if err != nil {
		return nil, err
	}
	var messageCommentList []Comment
	for _, comment := range commentList {
		messageCommentList = append(messageCommentList, messageComment(comment))
	}
	return &DouyinCommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: messageCommentList,
	}, nil
}

func messageComment(comment model.Comment) Comment {
	user, _ := GetUserInfo(comment.UserId, comment.UserId)
	return Comment{
		Id:         comment.CommentId,
		User:       *user,
		Content:    comment.Comment,
		CreateDate: comment.Time,
		LikeCount:  0,
		TeaseCount: 0,
	}
}
