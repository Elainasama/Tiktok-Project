package message

type DouyinPublishActionResponse struct {
	Response
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
