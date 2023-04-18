package model

type Comment struct {
	CommentId int64  `gorm:"primary_key;type:bigint(20) auto_increment;not null;column:comment_id"`
	UserId    int64  `gorm:"not null;column:user_id"`
	VideoId   int64  `gorm:"not null;column:video_id"`
	Comment   string `gorm:"not null;column:comment"`
	Time      string `gorm:"not null;column:time"`
}
