package message

type DouyinFavoriteActionResponse struct {
	Response
}

type DouyinFavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
