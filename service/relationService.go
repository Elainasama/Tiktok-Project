package service

import (
	"TikTokLite/dao"
	. "TikTokLite/message"
)

func RelationActionService(userId int64, toUserId int64, actionType int32) (*DouyinRelationActionResponse, error) {
	var err error
	Followed := dao.IsFollow(userId, toUserId)
	if actionType == 1 {
		if Followed {
			return &DouyinRelationActionResponse{Response: Response{
				StatusCode: 1,
				StatusMsg:  "already followed",
			}}, nil
		}
		err = dao.InsertRelation(userId, toUserId)
	} else if actionType == 2 {
		if !Followed {
			return &DouyinRelationActionResponse{Response: Response{
				StatusCode: 1,
				StatusMsg:  "not followed yet",
			}}, nil
		}
		err = dao.DeleteRelation(userId, toUserId)
	}
	if err != nil {
		return nil, err
	}
	return &DouyinRelationActionResponse{Response: SuccessResponse}, nil
}

func GetRelationFollowListService(userId int64) (*DouyinRelationFollowListResponse, error) {
	userList, err := dao.GetFollowUserList(userId)
	if err != nil {
		return nil, err
	}
	return &DouyinRelationFollowListResponse{
		Response: SuccessResponse,
		UserList: messageUserList(userList),
	}, nil
}

func GetRelationFollowerListService(userId int64) (*DouyinRelationFollowerListResponse, error) {
	userList, err := dao.GetFollowerUserList(userId)
	if err != nil {
		return nil, err
	}
	return &DouyinRelationFollowerListResponse{
		Response: SuccessResponse,
		UserList: messageUserList(userList),
	}, nil
}

func GetRelationFriendService(userId int64) (*DouyinFriendListResponse, error) {
	userList, err := dao.GetFriendList(userId)
	if err != nil {
		return nil, err
	}
	return &DouyinFriendListResponse{
		Response: SuccessResponse,
		UserList: messageUserList(userList),
	}, nil
}
