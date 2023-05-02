package message

type DouyinRelationActionResponse struct {
	Response
}

type DouyinRelationFollowListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type DouyinRelationFollowerListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type DouyinFriendListResponse struct {
	Response
	UserList []User `json:"user_list"`
}
