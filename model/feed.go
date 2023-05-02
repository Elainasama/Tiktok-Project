package model

type Video struct {
	VideoId       int64  `gorm:"primary_key;type:bigint(20) auto_increment;not null;column:video_id"`
	AuthorId      int64  `gorm:"not null;column:author_id"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Title         string `gorm:"column:title;"`
	Author        User   `gorm:"foreignkey:AuthorId"`
}
