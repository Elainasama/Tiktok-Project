package message

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	LikeCount  int64  `json:"like_count,omitempty"`
	TeaseCount int64  `json:"tease_count,omitempty"`
}

type DouyinCommentActionResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type DouyinCommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"`
}
